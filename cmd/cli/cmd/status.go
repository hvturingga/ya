package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hvturingga/ya/cmd/cli/internal"
	"github.com/hvturingga/ya/conf"
	api_server "github.com/hvturingga/ya/internal/api-server"
	"github.com/hvturingga/ya/internal/entclient"
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

			getUser, err := internal.GetUser(ctx, client)
			if err != nil {
				return
			}
			fmt.Println("\nActive Provider:")
			if getUser.Edges.Provider != nil {
				fmt.Fprintf(w, "%s\t%s\n", "Name", getUser.Edges.Provider.Name)
				fmt.Fprintf(w, "%s\t%s\n", "Version", getUser.Edges.Provider.Version)
				fmt.Fprintf(w, "%s\t%s\n", "Path", getUser.Edges.Provider.Path)
				w.Flush()
			}

			fmt.Println("\nActive Subscription:")
			if getUser.Edges.Subscribe != nil {
				fmt.Fprintf(w, "%s\t%s\n", "Name", getUser.Edges.Subscribe.Name)
				fmt.Fprintf(w, "%s\t%s\n", "Config", getUser.Edges.Subscribe.Conf)
				w.Flush()
			}

			fmt.Println("\nDaemon:")
			if getUser.Edges.Daemon != nil {
				fmt.Fprintf(w, "%s\t%v\n", "Enabled", getUser.Edges.Daemon.Enable)
				fmt.Fprintf(w, "%s\t%s\n", "Path", getUser.Edges.Daemon.Path)
				w.Flush()
			}
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
}
