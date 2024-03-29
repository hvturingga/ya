package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var version = "dev"

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("YA Version: %s\n", version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
