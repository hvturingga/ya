package cmd

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/hvturingga/ya/cmd/cli/internal"
	"github.com/hvturingga/ya/conf"
	"github.com/hvturingga/ya/ent/node"
	"github.com/hvturingga/ya/internal/entclient"
	"github.com/hvturingga/ya/internal/nodeswitch"
	"github.com/spf13/cobra"
	"io"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"
	"text/tabwriter"
)

var nodeCmd = &cobra.Command{
	Use:     "node",
	Short:   "View and switch subscription nodes (Proxies).",
	Aliases: []string{"n"},
}

func proxies() []byte {
	url := fmt.Sprintf("http://%s/proxies", conf.ClashAPI)

	r, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching proxies:", err)
		os.Exit(1)
	}
	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		os.Exit(1)
	}
	return body
}

func indexOf(slice []string, str string) bool {
	for _, v := range slice {
		if v == str {
			return true
		}
	}
	return false
}

var nodesCmd = &cobra.Command{
	Use:     "list --group <name>",
	Short:   "Used to display nodeswitch (Proxies) information in subscription.",
	Long:    `This command prints all nodeswitch (Proxies) and group information in the subscription.`,
	Aliases: []string{"proxies", "l", "ls"},
	Run: func(cmd *cobra.Command, args []string) {
		all := proxies()

		reader := bytes.NewReader(all)

		decoder := json.NewDecoder(reader)

		if _, err := decoder.Token(); err != nil {
			fmt.Println("Error reading start of JSON:", err)
			return
		}

		var result map[string]map[string]map[string]interface{}
		if err := json.Unmarshal(all, &result); err != nil {
			fmt.Println("Error parsing JSON:", err)
			os.Exit(1)
		}

		group, _ := cmd.Flags().GetString("group")
		if strings.TrimSpace(group) == "" {
			var selectorsInOrder []string
			var globalSelectors []string
			var hasGlobal bool

			foundProxies := false
			for decoder.More() {
				token, err := decoder.Token()
				if err != nil {
					fmt.Println("Error reading JSON token:", err)
					return
				}

				if key, ok := token.(string); ok && key == "proxies" {
					foundProxies = true
					if _, err := decoder.Token(); err != nil {
						fmt.Println("Error reading start of proxies object:", err)
						return
					}

					for decoder.More() {
						proxyToken, err := decoder.Token()
						if err != nil {
							fmt.Println("Error reading proxy token:", err)
							return
						}
						proxyKey, ok := proxyToken.(string)
						if !ok {
							fmt.Println("Error asserting proxy key as string")
							return
						}

						var proxyValue map[string]interface{}
						if err := decoder.Decode(&proxyValue); err != nil {
							fmt.Println("Error decoding proxy value:", err)
							return
						}
						if _, ok := proxyValue["all"]; ok && reflect.TypeOf(proxyValue["all"]).Kind() == reflect.Slice {
							if proxyKey == "GLOBAL" {
								hasGlobal = true
								for _, v := range proxyValue["all"].([]interface{}) {
									globalSelectors = append(globalSelectors, v.(string))
								}
								continue
							}
							selectorsInOrder = append(selectorsInOrder, proxyKey)
						}
					}

					if _, err := decoder.Token(); err != nil {
						fmt.Println("Error reading end of proxies object:", err)
						return
					}
					break
				} else {
					if err := decoder.Decode(&json.RawMessage{}); err != nil {
						fmt.Println("Error skipping non-proxy value:", err)
						return
					}
				}
			}

			if !foundProxies {
				fmt.Println("Did not find 'proxies' in JSON")
				return
			}

			if hasGlobal {
				var orderedSelectors []string
				for _, globalKey := range globalSelectors {
					for _, selector := range selectorsInOrder {
						if globalKey == selector {
							orderedSelectors = append(orderedSelectors, selector)
							break
						}
					}
				}
				for _, selector := range selectorsInOrder {
					if !indexOf(globalSelectors, selector) {
						orderedSelectors = append(orderedSelectors, selector)
					}
				}
				orderedSelectors = append(orderedSelectors, "GLOBAL")
				selectorsInOrder = orderedSelectors
			}

			w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
			fmt.Fprintln(w, "ID\tGroup\t")
			for i, selector := range selectorsInOrder {
				fmt.Fprintf(w, "%d\t%s\t\n", i+1, selector)
			}
			w.Flush()

			fmt.Print("Enter the index of the group you want to inspect (or press Enter to exit): ")
			var input string
			if _, err := fmt.Scanln(&input); err != nil {
				if err.Error() == "unexpected newline" {
					fmt.Println("No input group. Exiting.")
					return
				}
				fmt.Println("Error reading input:", err)
				os.Exit(1)
			}

			id, err := strconv.Atoi(input)
			if err != nil || id < 1 || id > len(selectorsInOrder) {
				fmt.Println("Invalid ID. Exiting.")
				os.Exit(1)
			}

			group = selectorsInOrder[id-1]
		}
		if proxies, ok := result["proxies"][group]; ok {
			w := tabwriter.NewWriter(os.Stdout, 0, 0, 5, ' ', 0)
			for _, proxy := range proxies["all"].([]interface{}) {
				name := proxy.(string)
				if name == proxies["now"] {
					fmt.Fprintf(w, "%s\t%s \n", name, "✔")
				} else {
					fmt.Fprintf(w, "%s\t%s \n", name, " ")
				}
			}
			w.Flush()
		} else {
			fmt.Println("Group not found")
		}

	},
}

var nodeSwitchCmd = &cobra.Command{
	Use:     "switch --group <name>",
	Short:   "Switch switch.",
	Aliases: []string{"s"},
	Run: func(cmd *cobra.Command, args []string) {
		all := proxies()

		reader := bytes.NewReader(all)

		decoder := json.NewDecoder(reader)

		if _, err := decoder.Token(); err != nil {
			fmt.Println("Error reading start of JSON:", err)
			return
		}

		var result map[string]map[string]map[string]interface{}
		if err := json.Unmarshal(all, &result); err != nil {
			fmt.Println("Error parsing JSON:", err)
			os.Exit(1)
		}

		group, _ := cmd.Flags().GetString("group")
		if strings.TrimSpace(group) == "" {
			var selectorsInOrder []string
			var globalSelectors []string
			var hasGlobal bool

			foundProxies := false
			for decoder.More() {
				token, err := decoder.Token()
				if err != nil {
					fmt.Println("Error reading JSON token:", err)
					return
				}

				if key, ok := token.(string); ok && key == "proxies" {
					foundProxies = true
					if _, err := decoder.Token(); err != nil {
						fmt.Println("Error reading start of proxies object:", err)
						return
					}

					for decoder.More() {
						proxyToken, err := decoder.Token()
						if err != nil {
							fmt.Println("Error reading proxy token:", err)
							return
						}
						proxyKey, ok := proxyToken.(string)
						if !ok {
							fmt.Println("Error asserting proxy key as string")
							return
						}

						var proxyValue map[string]interface{}
						if err := decoder.Decode(&proxyValue); err != nil {
							fmt.Println("Error decoding proxy value:", err)
							return
						}
						if _, ok := proxyValue["all"]; ok && reflect.TypeOf(proxyValue["all"]).Kind() == reflect.Slice {
							if proxyKey == "GLOBAL" {
								hasGlobal = true
								for _, v := range proxyValue["all"].([]interface{}) {
									globalSelectors = append(globalSelectors, v.(string))
								}
								continue
							}
							selectorsInOrder = append(selectorsInOrder, proxyKey)
						}
					}

					if _, err := decoder.Token(); err != nil {
						fmt.Println("Error reading end of proxies object:", err)
						return
					}
					break
				} else {
					if err := decoder.Decode(&json.RawMessage{}); err != nil {
						fmt.Println("Error skipping non-proxy value:", err)
						return
					}
				}
			}

			if !foundProxies {
				fmt.Println("Did not find 'proxies' in JSON")
				return
			}

			if hasGlobal {
				var orderedSelectors []string
				for _, globalKey := range globalSelectors {
					for _, selector := range selectorsInOrder {
						if globalKey == selector {
							orderedSelectors = append(orderedSelectors, selector)
							break
						}
					}
				}
				for _, selector := range selectorsInOrder {
					if !indexOf(globalSelectors, selector) {
						orderedSelectors = append(orderedSelectors, selector)
					}
				}
				orderedSelectors = append(orderedSelectors, "GLOBAL")
				selectorsInOrder = orderedSelectors
			}

			w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
			fmt.Fprintln(w, "ID\tGroup\t")
			for i, selector := range selectorsInOrder {
				fmt.Fprintf(w, "%d\t%s\t\n", i+1, selector)
			}
			w.Flush()

			fmt.Print("Enter the index of the group you want to inspect (or press Enter to exit): ")
			var input string
			if _, err := fmt.Scanln(&input); err != nil {
				if err.Error() == "unexpected newline" {
					fmt.Println("No input group. Exiting.")
					return
				}
				fmt.Println("Error reading input:", err)
				os.Exit(1)
			}

			id, err := strconv.Atoi(input)
			if err != nil || id < 1 || id > len(selectorsInOrder) {
				fmt.Println("Invalid ID. Exiting.")
				os.Exit(1)
			}

			group = selectorsInOrder[id-1]
		}

		if proxies, ok := result["proxies"][group]; ok {
			w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

			var nodes []string
			allProxies := proxies["all"].([]interface{})
			for index, proxy := range allProxies {
				name := proxy.(string)
				if name == proxies["now"] {
					fmt.Fprintf(w, "%d\t%s\t%s\n", index, name, "✔")
				} else {
					fmt.Fprintf(w, "%d\t%s\t%s\n", index, name, " ")
				}
				nodes = append(nodes, name)
			}
			w.Flush()

			fmt.Print("Please enter the index of the nodeswitch you want to switch to (1-based index or press Enter to exit): ")
			var input string
			if _, err := fmt.Scanln(&input); err != nil {
				if err.Error() == "unexpected newline" {
					fmt.Println("No input nodeswitch. Exiting.")
					return
				}
				fmt.Println("Error reading input:", err)
				os.Exit(1)
			}

			id, err := strconv.Atoi(input)
			if err != nil || id < 0 || id >= len(allProxies) {
				fmt.Println("Invalid ID. Exiting.")
				os.Exit(1)
			}

			sel := nodes[id]

			fmt.Printf("You have selected the group: %s, node: %s\n", group, sel)

			if err := nodeswitch.New(group, sel).Switch(); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			ctx := context.Background()
			client, err := entclient.New()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			getUser, err := internal.GetUser(ctx, client)
			if err != nil {
				return
			}

			subscribe := getUser.QuerySubscribe().OnlyX(ctx)

			exist := subscribe.
				QueryNodes().
				Where(
					node.GroupEQ(group),
				).
				ExistX(ctx)
			if !exist {
				client.Node.Create().
					SetGroup(group).
					SetName(sel).
					SetSubscribe(subscribe).
					SaveX(ctx)
			} else {
				client.Node.Update().
					Where(
						node.GroupEQ(group),
					).
					SetName(sel).
					SaveX(ctx)
			}
		} else {
			fmt.Println("Group not found")
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(nodeCmd)

	nodeCmd.AddCommand(nodesCmd)
	nodeCmd.AddCommand(nodeSwitchCmd)

	nodesCmd.Flags().StringP("group", "g", "", "group name")
	nodeSwitchCmd.Flags().StringP("group", "g", "", "group name")
}
