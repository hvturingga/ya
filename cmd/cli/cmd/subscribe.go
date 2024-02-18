package cmd

import (
	"context"
	"fmt"
	"github.com/hvturingga/ya/cmd/cli/internal/subscribe"
	"github.com/hvturingga/ya/ent"
	"github.com/hvturingga/ya/ent/provider"
	S "github.com/hvturingga/ya/ent/subscribe"
	"github.com/hvturingga/ya/ent/user"
	"github.com/hvturingga/ya/internal/entclient"
	"github.com/hvturingga/ya/internal/ya"
	"net/url"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

// subscribeCmd represents the subscribe command
var subscribeCmd = &cobra.Command{
	Use:     "subscribe",
	Short:   "Subscription management.",
	Long:    `This command is used to manage subscriptions.`,
	Aliases: []string{"s", "sub"},
}

var listSubscribeCmd = &cobra.Command{
	Use:     "list --id <id>",
	Short:   "List subscriptions.",
	Aliases: []string{"ls", "l"},
	Run: func(cmd *cobra.Command, args []string) {

		ctx := context.Background()

		client, err := entclient.New()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
		defer client.Close()

		id, _ := cmd.Flags().GetString("id")
		if strings.TrimSpace(id) != "" {
			i, err := strconv.Atoi(id)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				os.Exit(1)
			}

			sub, err := client.Subscribe.Query().
				Where(
					S.IDEQ(i),
				).
				Only(ctx)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				os.Exit(1)
			}

			w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
			fmt.Fprintf(w, "%s\t%s\n", "Name", sub.Name)
			fmt.Fprintf(w, "%s\t%s\n", "Link", sub.Link)
			fmt.Fprintf(w, "%s\t%s\n", "Path", sub.Conf)
			w.Flush()

		} else {
			U, err := client.User.Query().
				Where(
					user.NameEQ(ya.GetUser()),
				).
				WithProvider().
				WithSubscribe().
				Only(ctx)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				os.Exit(1)
			}
			if U.Edges.Provider == nil {
				fmt.Println("No provider found, please add one.")
				os.Exit(1)
				return
			}

			all, err := client.Subscribe.Query().
				Where(
					S.HasProviderWith(
						provider.IDEQ(
							U.Edges.Provider.ID,
						),
					),
				).
				All(ctx)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				os.Exit(1)
			}
			w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

			fmt.Fprintf(w, "%s\t%s\t%s\n", "ID", "Name", "Now")
			for _, sub := range all {
				if U.Edges.Subscribe != nil && sub.ID == U.Edges.Subscribe.ID {
					fmt.Fprintf(w, "%d\t%s\t%s\n", sub.ID, sub.Name, "✔")
				} else {
					fmt.Fprintf(w, "%d\t%s\t%s\n", sub.ID, sub.Name, " ")
				}
			}
			w.Flush()
		}
	},
}

var addSubscribeCmd = &cobra.Command{
	Use:     "add <name> --link <link> --path <path>",
	Short:   "Add a new subscription.",
	Aliases: []string{"a"},
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		if len(name) > 24 {
			fmt.Println("Error: The name cannot be more than 24 characters long.")
			os.Exit(1)
		}

		link, _ := cmd.Flags().GetString("link")
		path, _ := cmd.Flags().GetString("path")
		if (strings.TrimSpace(link) == "" && strings.TrimSpace(path) == "") || (strings.TrimSpace(link) != "" && strings.TrimSpace(path) != "") {
			fmt.Println("Error: Either link or path must be provided, but not both.")
			os.Exit(1)
		}

		client, err := entclient.New()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
		defer client.Close()

		ctx := context.Background()

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

		sub := client.Subscribe.Query()
		exist, err := sub.
			Where(
				S.NameEQ(name),
				S.HasProviderWith(
					provider.IDEQ(
						query.Edges.Provider.ID,
					),
				),
			).
			Exist(ctx)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
		if exist {
			fmt.Println("The subscription name already exists.")
			os.Exit(1)
		}

		if strings.TrimSpace(path) != "" {
			if _, err := os.Stat(path); os.IsNotExist(err) {
				fmt.Println("Cannot be added because the config file does not exist in the file system.")
				os.Exit(1)
			}
		} else if strings.TrimSpace(link) != "" {
			if _, err := url.ParseRequestURI(link); err != nil {
				fmt.Println("The link must be a valid URL")
				os.Exit(1)
			}
			path = subscribe.Fetch(link)
		}

		create, err := client.Subscribe.Create().
			SetName(name).
			SetLink(link).
			SetProvider(query.Edges.Provider).
			SetConf(path).
			Save(ctx)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}

		if query.Edges.Subscribe == nil {
			client.User.Update().SetSubscribe(create).ExecX(ctx)
		}
		listSubscribeCmd.Run(cmd, args)
	},
}

var removeSubscribeCmd = &cobra.Command{
	Use:     "remove <id>",
	Short:   "Remove a subscription.",
	Args:    cobra.ExactArgs(1),
	Aliases: []string{"rm", "r", "delete", "d"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Removing subscription with id %s\n", args[0])

		client, err := entclient.New()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
		defer client.Close()

		ctx := context.Background()
		rid, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
		client.User.Query().
			Where(
				user.NameEQ(
					ya.GetUser(),
				),
				user.HasSubscribeWith(
					S.IDEQ(rid),
				),
			).
			QuerySubscribe().
			OnlyX(ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				fmt.Println("You cannot delete the subscription you are using, please switch and then delete it.")
				os.Exit(1)
			}
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
		client.Subscribe.DeleteOneID(rid).ExecX(ctx)
		listSubscribeCmd.Run(cmd, args)
	},
}

var editSubscribeCmd = &cobra.Command{
	Use:     "edit <id> --name <name> --link <link>",
	Short:   "Edit a subscription.",
	Args:    cobra.ExactArgs(1),
	Aliases: []string{"e"},
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		link, _ := cmd.Flags().GetString("link")
		path, _ := cmd.Flags().GetString("path")

		if len(name) > 24 {
			fmt.Println("Error: The name cannot be more than 24 characters long.")
			os.Exit(1)
		}

		eid, err := strconv.Atoi(args[0])

		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}

		client, err := entclient.New()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
		defer client.Close()

		ctx := context.Background()

		query, err := client.User.Query().
			Where(
				user.NameEQ(
					ya.GetUser(),
				),
			).
			QueryProvider().
			Only(ctx)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
		edit := client.Subscribe.
			Update().
			Where(
				S.IDEQ(eid),
				S.HasProviderWith(
					provider.IDEQ(query.ID),
				),
			)
		if strings.TrimSpace(name) != "" {
			edit.SetName(name)
		}
		if strings.TrimSpace(path) != "" {
			if _, err := os.Stat(path); os.IsNotExist(err) {
				fmt.Println("Cannot be added because the config file does not exist in the file system.")
				os.Exit(1)
			}
			edit.SetConf(path)
		}
		if strings.TrimSpace(link) != "" {
			if _, err := url.ParseRequestURI(link); err != nil {
				fmt.Println("The link must be a valid URL")
				os.Exit(1)
			}
			edit.SetLink(link)
		}
		edit.ExecX(ctx)
		listSubscribeCmd.Run(cmd, args)
	},
}

var switchSubscribeCmd = &cobra.Command{
	Use:     "switch --id <id>",
	Short:   "Switch subscription.",
	Aliases: []string{"s"},
	Run: func(cmd *cobra.Command, args []string) {

		client, err := entclient.New()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
		defer client.Close()
		ctx := context.Background()

		U, err := client.User.Query().
			Where(
				user.NameEQ(
					ya.GetUser(),
				),
			).
			WithProvider().
			WithSubscribe().
			Only(ctx)

		if U.Edges.Provider == nil {
			fmt.Println("No provider found, please add one.")
			os.Exit(1)
		}

		id, _ := cmd.Flags().GetString("id")
		var sid int
		if strings.TrimSpace(id) != "" {
			i, err := strconv.Atoi(id)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				os.Exit(1)
			}
			exist, err := U.QueryProvider().
				QuerySubscribes().
				Where(
					S.IDEQ(i),
				).
				Exist(ctx)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				os.Exit(1)
			}
			if !exist {
				fmt.Println("No subscription found")
				os.Exit(1)
			}
			sid = i
		} else {
			subscribes, err := U.QueryProvider().QuerySubscribes().All(ctx)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				os.Exit(1)
			}
			if len(subscribes) != 0 {
				w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

				fmt.Fprintf(w, "%s\t%s\t%s\t\n", "ID", "Name", "Now")
				for _, sub := range subscribes {
					if U.Edges.Subscribe != nil && sub.ID == U.Edges.Subscribe.ID {
						fmt.Fprintf(w, "%d\t%s\t%s\n", sub.ID, sub.Name, "✔")
					} else {
						fmt.Fprintf(w, "%d\t%s\t%s\n", sub.ID, sub.Name, " ")
					}
				}
				w.Flush()

				fmt.Print("Enter the ID of the subscription you want to select (or press Enter to exit): ")

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
				for _, sub := range subscribes {
					if id == sub.ID {
						found = true
						break
					}
				}
				if err != nil || !found {
					fmt.Println("Invalid ID. Exiting.")
					os.Exit(1)
				}

				sid = id
				fmt.Printf("You have selected subscription ID: %d\n", sid)
			} else {
				fmt.Println("No subscription found, please add one.")
				os.Exit(1)
			}

		}

		if U.Edges.Subscribe != nil {
			if U.Edges.Subscribe.ID == sid {
				fmt.Println("The subscription is in use.")
				os.Exit(0)
			}
		}

		U.Update().
			Where(
				user.NameEQ(
					ya.GetUser(),
				),
			).
			SetSubscribeID(sid).
			ExecX(ctx)

		fmt.Println("INFO: Subscription switched. Re-run daemon to use the new subscription.")
		fmt.Println("OK!")

		listSubscribeCmd.Run(cmd, args)
	},
}

var syncSubscribeCmd = &cobra.Command{
	Use:   "sync",
	Short: "Synchronize the current subscription.",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := entclient.New()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
		defer client.Close()

		ctx := context.Background()

		query, err := client.User.Query().
			Where(
				user.NameEQ(
					ya.GetUser(),
				),
			).
			WithSubscribe().
			WithProvider().
			Only(ctx)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
		if query.Edges.Subscribe != nil {
			subscribe.Fetch(query.Edges.Subscribe.Link)
		} else {
			fmt.Println("No subscription found")
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(subscribeCmd)
	subscribeCmd.AddCommand(listSubscribeCmd)
	subscribeCmd.AddCommand(addSubscribeCmd)
	subscribeCmd.AddCommand(editSubscribeCmd)
	subscribeCmd.AddCommand(removeSubscribeCmd)
	subscribeCmd.AddCommand(switchSubscribeCmd)
	subscribeCmd.AddCommand(syncSubscribeCmd)

	listSubscribeCmd.Flags().StringP("id", "i", "", "ID of the subscription")

	editSubscribeCmd.Flags().StringP("name", "n", "", "New name of the subscription")
	editSubscribeCmd.Flags().StringP("link", "l", "", "New link of the subscription")
	editSubscribeCmd.Flags().StringP("path", "p", "", "New link of the subscription")

	addSubscribeCmd.Flags().StringP("link", "l", "", "New link of the subscription")
	addSubscribeCmd.Flags().StringP("path", "p", "", "New path of the subscription")

	switchSubscribeCmd.Flags().StringP("id", "i", "", "ID of the subscription")
}
