package cmd

import (
	"context"
	"fmt"
	"github.com/hvturingga/ya/cmd/daemon-go/adapter"
	"github.com/hvturingga/ya/internal/entclient"
	"github.com/spf13/cobra"
	"os"
)

// stopCmd represents the stop command
var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop daemon",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		client, err := entclient.New()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
		defer client.Close()

		getUser := client.User.Query().WithProvider().WithSubscribe().WithDaemon().FirstX(ctx)
		adapter.NewAdapter(ctx, getUser).Stop()
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// stopCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// stopCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
