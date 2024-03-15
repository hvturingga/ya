package cmd

import (
	"context"
	"fmt"
	"github.com/google/go-github/v59/github"
	"github.com/hvturingga/ya/cmd/cli/embed"
	"github.com/hvturingga/ya/cmd/cli/internal"
	"github.com/hvturingga/ya/conf"
	"github.com/hvturingga/ya/ent"
	"github.com/hvturingga/ya/internal/entclient"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:     "install",
	Short:   "A brief description of your command",
	Long:    `Before use, you must initialise the application to make it operational.`,
	Aliases: []string{"i"},
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()
		client, err := entclient.New()
		if err != nil {
			return err
		}
		defer client.Close()

		tx, err := client.Tx(ctx)
		if err != nil {
			return err
		}
		user, err := internal.GetUser(ctx, tx.Client())
		if err != nil {
			return err
		}

		goos, arch := runtime.GOOS, runtime.GOARCH

		var provider *ent.Provider

		providerName, _ := cmd.Flags().GetString("provider")
		providerPathFlag, _ := cmd.Flags().GetString("path")

		if strings.TrimSpace(providerName) != "" {
			create := tx.Provider.Create()

			if providerPathFlag != "" {
				if _, err := os.Stat(providerPathFlag); os.IsNotExist(err) {
					fmt.Println("Cannot be added because the config file does not exist in the file system.")
					os.Exit(1)
				}
				create.SetPath(providerPathFlag)
			}

			provider = create.SetName(providerName).SaveX(ctx)

			fmt.Printf("Provider %s-%s initialized.\n", provider.Path, provider.Version)
		} else {
			providers := conf.GetProviderList()
			fmt.Println("Please select a provider to initialize: ")
			for i, provider := range providers {
				fmt.Printf("%d. %s\n", i+1, provider.Repo)
			}

			// select provider.
			pch, err := internal.GetUserSelection("Please enter the selected number for provider: ", len(providers))
			if err != nil {
				return err
			}

			// select version.
			sel := providers[pch-1]
			fmt.Printf("Please select a version for %s: \n", sel.Repo)
			for i, version := range sel.Version {
				fmt.Printf("%d. %s\n", i+1, version)
			}

			vch, err := internal.GetUserSelection("Please enter the selected number for version: ", len(sel.Version))
			if err != nil {
				return err
			}

			version := sel.Version[vch-1]

			fmt.Printf("You have selected Provider: %s, Version: %s\n", sel.Repo, version)

			gh := github.NewClient(nil)

			list, _, err := gh.Repositories.GetReleaseByTag(ctx, sel.Owner, sel.Repo, version)
			if err != nil {
				return err
			}

			var assets []*github.ReleaseAsset
			for _, asset := range list.Assets {
				if strings.Contains(asset.GetName(), goos) && strings.Contains(asset.GetName(), arch) {
					assets = append(assets, asset)
				}
			}
			if len(assets) == 0 {
				fmt.Println("No assets found for the current OS and architecture.")
				os.Exit(1)
			}

			for i, asset := range assets {
				fmt.Printf("%d. %s \n", i+1, asset.GetName())
			}

			ach, err := internal.GetUserSelection("Please enter the number of the asset to download: ", len(assets))
			if err != nil {
				return err
			}

			providerPath, err := internal.DownloadProviderRelease(sel.Repo, version, assets[ach-1])
			if err != nil {
				return err
			}

			provider = tx.Provider.Create().SetPath(providerPath).SetName(sel.Repo).SetVersion(version).SaveX(ctx)

			fmt.Printf("Provider %s-%s initialized.\n", provider.Path, provider.Version)
		}

		var bindir string
		if goos == "windows" {
			bindir = fmt.Sprintf("ya-daemon-%s-%s.exe", goos, arch)
		} else if goos == "linux" {
			bindir = fmt.Sprintf("ya-daemon-%s-%s", goos, arch)
		}

		daemonPath := filepath.Join(conf.GetDaemonPath(), bindir)

		if err := os.MkdirAll(filepath.Dir(daemonPath), 0755); err != nil {
			fmt.Printf("Error creating directories: %s\n", err)
			os.Exit(1)
		}

		if err := os.WriteFile(daemonPath, embed.Daemon, 0755); err != nil {
			fmt.Printf("Error: %s\n", err)
			os.Exit(1)
		}
		daemon := tx.Daemon.Create().SetPath(daemonPath).SaveX(ctx)

		user.Update().ClearDaemon().ClearProvider().SetProvider(provider).SetDaemon(daemon).ExecX(ctx)

		fmt.Printf("Daemon %s initialized.\n", daemon.Path)

		if err := tx.Commit(); err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(installCmd)

	installCmd.Flags().StringP("provider", "p", "", "The name of the provider path")
	installCmd.Flags().String("path", "", "The path to the provider binary")

}
