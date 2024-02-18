package cmd

import (
	"bufio"
	"context"
	"fmt"
	"github.com/hvturingga/ya/cmd/cli/embed"
	provider2 "github.com/hvturingga/ya/cmd/cli/internal/provider"
	"github.com/hvturingga/ya/ent"
	"github.com/hvturingga/ya/ent/user"
	"github.com/hvturingga/ya/internal/entclient"
	P "github.com/hvturingga/ya/internal/provider"
	"github.com/hvturingga/ya/internal/ya"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:     "init --provider <path>",
	Short:   "Initialise the application.",
	Long:    `Before use, you must initialise the application to make it operational.`,
	Aliases: []string{"i"},
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()

		client, err := entclient.New()
		if err != nil {
			fmt.Printf("Error: %s", err)
			os.Exit(1)
		}
		defer client.Close()

		providers := P.GetProviderList()
		fmt.Println("Please select a provider to initialize: ")
		for i, provider := range providers {
			fmt.Printf("%d. %s %s\n", i+1, provider.Repo, provider.Version)
		}

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Please enter the selected number: ")

		number, _ := reader.ReadString('\n')
		pch, err := strconv.Atoi(strings.TrimSpace(number))
		if err != nil || pch < 1 || pch > len(providers) {
			fmt.Println("invalid selection, please rerun the command")
			os.Exit(1)
		}

		sel := providers[pch-1]

		user, err := func() (*ent.User, error) {
			query, err := client.User.Query().
				Where(
					user.NameEQ(
						ya.GetUser(),
					),
				).
				Only(ctx)
			if err != nil {
				if ent.IsNotFound(err) {
					create, err := client.User.Create().
						SetName(
							ya.GetUser(),
						).
						Save(ctx)
					if err != nil {
						return nil, err
					}

					return create, nil
				}
				return nil, err
			}
			return query, nil
		}()
		if err != nil {
			fmt.Printf("Error: %s", err)
			os.Exit(1)
		}

		provider, _ := cmd.Flags().GetString("provider")
		if strings.TrimSpace(provider) != "" {
			if _, err := os.Stat(provider); os.IsNotExist(err) {
				fmt.Println("Cannot be added because the config file does not exist in the file system.")
				os.Exit(1)
			}
			provider, err := client.Provider.Create().
				SetName(sel.Repo).
				SetVersion(sel.Version).
				SetPath(provider).
				Save(ctx)
			if err != nil {
				fmt.Printf("Error: %s", err)
				os.Exit(1)
			}
			user.Update().ClearSubscribe().SetProvider(provider).ExecX(ctx)

			fmt.Printf("Provider %s-%s initialized.\n", provider.Path, provider.Version)
		} else {
			provider2.Select(client, user, sel)
		}

		goos := runtime.GOOS
		goarch := runtime.GOARCH
		var bindir string
		if goos == "windows" {
			bindir = fmt.Sprintf("ya-daemon-%s-%s.exe", goos, goarch)
		} else if goos == "linux" {
			bindir = fmt.Sprintf("ya-daemon-%s-%s", goos, goarch)
		}

		path := filepath.Join(ya.GetDaemonDir(), bindir)

		if err := os.WriteFile(path, embed.Daemon, 0755); err != nil {
			fmt.Printf("Error: %s\n", err)
			os.Exit(1)
		}
		daemon := client.Daemon.Create().SetPath(path).SaveX(ctx)
		user.Update().ClearDaemon().SetDaemon(daemon).ExecX(ctx)

		fmt.Printf("Daemon %s initialized.\n", daemon.Path)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	initCmd.Flags().StringP("provider", "p", "", "path to the provider.")
}
