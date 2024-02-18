package cmd

import (
	"bufio"
	"fmt"
	"github.com/hvturingga/ya/conf"
	"github.com/spf13/cobra"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var confCmd = &cobra.Command{
	Use:     "conf",
	Short:   "For manage application settings and display program information.",
	Aliases: []string{"c"},
}

var confModeCmd = &cobra.Command{
	Use:     "mode --mode <mode>",
	Short:   "Switch the proxy mode (Global | Rule | Direct).",
	Aliases: []string{"m", ""},
	Run: func(cmd *cobra.Command, args []string) {
		mode, _ := cmd.Flags().GetString("mode")

		if strings.TrimSpace(mode) == "" {
			modes := []string{"global", "rule", "direct"}
			for i, m := range modes {
				fmt.Printf("%d. %s\n", i+1, m)
			}
			fmt.Print("Please choose a mode by entering the corresponding number (or press Enter to exit): ")

			reader := bufio.NewReader(os.Stdin)
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)

			if input == "" {
				fmt.Println("Operation cancelled.")
				os.Exit(0)
			}

			if num, err := strconv.Atoi(input); err == nil && num > 0 && num <= len(modes) {
				mode = modes[num-1]
			} else {
				fmt.Println("Invalid selection. Please choose a valid number.")
				os.Exit(1)
			}
		}

		valid := map[string]bool{"global": true, "rule": true, "direct": true}
		if !valid[mode] {
			fmt.Println("Invalid mode. Please choose from Global, Rule, or Direct.")
			os.Exit(1)
		}

		fmt.Printf("Proxy mode set to: %s\n", mode)

		url := fmt.Sprintf("http://%s/configs", conf.ClashAPI)
		payload := strings.NewReader(fmt.Sprintf(`{"mode": "%s"}`, mode))

		req, err := http.NewRequest(http.MethodPatch, url, payload)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		req.Header.Add("Content-Type", "application/json")

		client := &http.Client{}
		res, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusNoContent {
			fmt.Printf("Error: %s\n", res.Status)
			body, err := io.ReadAll(res.Body)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			fmt.Println(string(body))
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(confCmd)
	confCmd.AddCommand(confModeCmd)

	confModeCmd.Flags().StringP("mode", "m", "", "Proxy mode (Global | Rule | Direct)")
}
