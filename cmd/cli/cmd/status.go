package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hvturingga/ya/conf"
	"github.com/hvturingga/ya/ent/user"
	api_server "github.com/hvturingga/ya/internal/api-server"
	"github.com/hvturingga/ya/internal/entclient"
	"github.com/hvturingga/ya/internal/ya"
	"io"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:     "status",
	Short:   "Print the status of the application.",
	Aliases: []string{"state"},
	Run: func(cmd *cobra.Command, args []string) {
		client, err := entclient.New()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
		defer client.Close()

		ctx := context.Background()
		client.Daemon.Query().Where()

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

		{
			fmt.Fprintf(w, "%s\t%s\n", "User", ya.GetUser())
			w.Flush()
		}

		{
			r, err := http.Get(fmt.Sprintf("http://%s/configs", conf.ClashAPI))
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				os.Exit(1)
			}

			defer r.Body.Close()
			body, err := io.ReadAll(r.Body)
			if err != nil {
				fmt.Println("Error reading response body:", err)
				os.Exit(1)
			}

			var result map[string]interface{}
			if err := json.Unmarshal(body, &result); err != nil {
				fmt.Println("Error parsing JSON:", err)
				os.Exit(1)
			}

			fmt.Println("\nConfiguration:")
			fmt.Fprintf(w, "%s\t%s\n", "Mode", result["mode"])
			fmt.Fprintf(w, "%s\t%s\n", "Clash API", conf.ClashAPI)
			w.Flush()
		}

		{ // Interact with Database.

			// Query the Provider in use by the user.
			query, err := client.User.Query().
				Where(
					user.NameEQ(
						ya.GetUser(),
					),
				).
				WithProvider().
				WithSubscribe().
				Only(ctx)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				os.Exit(1)
			}
			fmt.Println("\nActive Provider:")
			if query.Edges.Provider != nil {
				fmt.Fprintf(w, "%s\t%s\n", "Name", query.Edges.Provider.Name)
				fmt.Fprintf(w, "%s\t%s\n", "Version", query.Edges.Provider.Version)
				fmt.Fprintf(w, "%s\t%s\n", "Path", query.Edges.Provider.Path)
				w.Flush()
			}

			fmt.Println("\nActive Subscription:")
			if query.Edges.Subscribe != nil {
				fmt.Fprintf(w, "%s\t%s\n", "Name", query.Edges.Subscribe.Name)
				fmt.Fprintf(w, "%s\t%s\n", "Config", query.Edges.Subscribe.Conf)
				w.Flush()
			}

			fmt.Println("\n Daemon:")
			daemon, err := client.Daemon.Query().Only(ctx)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				os.Exit(1)
			}
			fmt.Fprintf(w, "%s\t%s\n", "Enabled", daemon.Enable)
			fmt.Fprintf(w, "%s\t%s\n", "Path", daemon.Path)
			w.Flush()
		}

		fmt.Println("\n ")
		go api_server.Traffic()

		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

		<-sigChan
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// statusCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// statusCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
