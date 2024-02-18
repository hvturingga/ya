package cmd

import (
	"context"
	"fmt"
	"github.com/hvturingga/ya/cmd/daemon-go/adapter"
	"github.com/hvturingga/ya/ent"
	U "github.com/hvturingga/ya/ent/user"
	api_server "github.com/hvturingga/ya/internal/api-server"
	"github.com/hvturingga/ya/internal/entclient"
	"github.com/hvturingga/ya/internal/nodeswitch"
	"github.com/hvturingga/ya/internal/ya"
	"github.com/spf13/cobra"
	"os"
	"strings"
	"sync"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start <database>",
	Short: "Start daemon",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()

		path, _ := cmd.Flags().GetString("database")

		if strings.TrimSpace(path) == "" {
			path = ya.GetYaDatabasePath()
		}

		client, err := entclient.New(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to create client: %v\n", err)
			return
		}
		defer client.Close()

		// Inline fetchUser logic
		user, err := client.User.Query().
			Where(
				U.ActiveEQ(true),
			).
			WithProvider().
			WithSubscribe(func(query *ent.SubscribeQuery) {
				query.WithNodes()
			}).
			Only(ctx)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to fetch user: %v\n", err)
			return
		}
		if user.Edges.Provider == nil {
			fmt.Fprintf(os.Stderr, "User %s has no provider\n", user.Name)
			return
		}
		if user.Edges.Subscribe == nil {
			fmt.Fprintf(os.Stderr, "User %s has no subscribe\n", user.Name)
			return
		}

		adapter.NewAdapter(ctx, user).Start()

		// Inline switchNodes logic
		if !api_server.WaitForAPIServer() {
			fmt.Fprintf(os.Stderr, "Failed to detect API server startup\n")
			return
		}

		var wg sync.WaitGroup
		errors := make(chan error, len(user.Edges.Subscribe.Edges.Nodes))

		for _, l := range user.Edges.Subscribe.Edges.Nodes {
			wg.Add(1)
			go func(l *ent.Node) {
				defer wg.Done()
				if err := nodeswitch.New(l.Group, l.Name).Switch(); err != nil {
					errors <- err
				}
			}(l)
		}

		wg.Wait()
		close(errors)

		for err := range errors {
			if err != nil {
				fmt.Fprintf(os.Stderr, "Failed to switch nodes: %v\n", err)
				return
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	startCmd.Flags().StringP("database", "d", "", "Run as daemon")
}
