package cmd

import (
	"context"
	"fmt"
	"github.com/hvturingga/ya/cmd/cli/internal"
	"github.com/hvturingga/ya/internal/entclient"
	"github.com/spf13/cobra"
	"os"
	"strconv"
	"text/tabwriter"
)

// providerCmd represents the provider command
var providerCmd = &cobra.Command{
	Use:   "provider",
	Short: "Used to manage the proxy platform utilised.",
	Long: `It allows for the display or alteration of the
proxy platform that ya operates on, among othe
r functionalities, for provider management.
`,
	Aliases: []string{"p"},
}

var switchProviderCmd = &cobra.Command{
	Use:     "switch",
	Short:   "Switch between different providers.",
	Long:    `This command allows the user to switch between different proxy providers.`,
	Aliases: []string{"s"},
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := entclient.New()
		if err != nil {
			return err
		}
		defer client.Close()

		ctx := context.Background()

		getUser, err := internal.GetUser(ctx, client)
		if err != nil {
			return err
		}
		all, err := client.Provider.Query().All(ctx)
		if err != nil {
			return err
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\n", "ID", "Name", "Version", "Now", "Path")
		for _, provider := range all {
			if getUser.Edges.Provider != nil && provider.ID == getUser.Edges.Provider.ID {
				fmt.Fprintf(w, "%d\t%s\t%s\t%s\t%s\n", provider.ID, provider.Name, provider.Version, "✔", provider.Path)
			} else {
				fmt.Fprintf(w, "%d\t%s\t%s\t%s\t%s\n", provider.ID, provider.Name, provider.Version, " ", provider.Path)
			}
		}
		w.Flush()

		fmt.Print("Enter the ID of the provider you want to select (or press Enter to exit): ")
		var input string
		if _, err := fmt.Scanln(&input); err != nil {
			if err.Error() == "unexpected newline" {
				fmt.Println("No input provided. Exiting.")
				os.Exit(0)
			} else {
				fmt.Println("Error reading input:", err)
				os.Exit(1)
			}
		}

		id, err := strconv.Atoi(input)
		found := false
		for _, sub := range all {
			if id == sub.ID {
				found = true
				break
			}
		}
		if err != nil || !found {
			fmt.Println("Invalid ID. Exiting.")
			os.Exit(1)
		}

		if getUser.Edges.Provider != nil {
			if id == getUser.Edges.Provider.ID {
				fmt.Println("The provider is in use.")
				os.Exit(0)
			}
		}

		getUser.Update().
			ClearSubscribe().
			SetProviderID(id).
			ExecX(ctx)

		fmt.Println("INFO: Provider switched. Re-run daemon to use the new provider.")

		return nil
	},
}

func init() {
	rootCmd.AddCommand(providerCmd)
	providerCmd.AddCommand(switchProviderCmd)
}
