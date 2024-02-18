package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hvturingga/ya/cmd/cli/internal/tmpl"
	"github.com/hvturingga/ya/conf"
	"github.com/hvturingga/ya/ent"
	U "github.com/hvturingga/ya/ent/user"
	"github.com/hvturingga/ya/internal/entclient"
	"github.com/hvturingga/ya/internal/ya"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"runtime"
	"text/template"
)

// daemonCmd represents the daemon command
var daemonCmd = &cobra.Command{
	Use:     "daemon",
	Short:   "Used to manage the daemon processes.",
	Long:    `Commands for daemon processes related.`,
	Aliases: []string{"d"},
}

var daemonRunCmd = &cobra.Command{
	Use:     "run",
	Short:   "Run the daemon.",
	Long:    `Run the daemon service.`,
	Aliases: []string{"start"},
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		client, err := entclient.New()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer client.Close()

		user := client.User.
			Query().
			Where(
				U.NameEQ(
					ya.GetUser(),
				),
			).
			OnlyX(ctx)

		subscribe, err := user.
			QuerySubscribe().
			Only(ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				fmt.Println("No subscription found.")
				os.Exit(0)
			}
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}

		f, err := os.ReadFile(subscribe.Conf)
		if err != nil {
			fmt.Printf("Failed to read JSON file: %v\n", err)
			os.Exit(1)
		}

		var j map[string]interface{}
		if err = json.Unmarshal(f, &j); err != nil {
			fmt.Printf("Failed to parse JSON: %v\n", err)
			os.Exit(1)
		}
		experimental, ok := j["experimental"].(map[string]interface{})
		if !ok {
			experimental = make(map[string]interface{})
			j["experimental"] = experimental
		}

		cla, exists := experimental["clash_api"].(map[string]interface{})
		if !exists || cla["external_controller"] != "127.0.0.1:50210" {
			experimental["clash_api"] = map[string]interface{}{
				"external_controller": conf.ClashAPI,
				"secret":              "",
			}
		}

		modified, err := json.MarshalIndent(j, "", "    ")
		if err != nil {
			fmt.Printf("Failed to marshal JSON: %v\n", err)
			os.Exit(1)
		}

		if err = os.WriteFile(subscribe.Conf, modified, 0644); err != nil {
			fmt.Printf("Failed to write modified JSON back to file: %v\n", err)
			os.Exit(1)
		}

		daemon := user.QueryDaemon().OnlyX(ctx)
		goos := runtime.GOOS
		if goos == "windows" {
			cmd := fmt.Sprintf("Start-Process powershell -WindowStyle Hidden -ArgumentList '%s start'", daemon.Path)
			fmt.Println(cmd)

			command := exec.Command("powershell", "-Command", cmd)
			if err := command.Run(); err != nil {
				fmt.Printf("Error starting command on Windows: %s\n", err)
				os.Exit(1)
			}
		} else if goos == "linux" {
			command := exec.Command("sudo", daemon.Path, "start")
			fmt.Println(command)

			if err := command.Run(); err != nil {
				fmt.Printf("Error starting command on Linux: %s\n", err)
				os.Exit(1)
			}
		}
	},
}

var daemonKillCmd = &cobra.Command{
	Use:     "kill",
	Short:   "Kill the daemon.",
	Long:    `Kill the daemon service.`,
	Aliases: []string{"k", "stop"},
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		client, err := entclient.New()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer client.Close()

		user := client.User.
			Query().
			Where(
				U.NameEQ(
					ya.GetUser(),
				),
			).
			OnlyX(ctx)

		daemon := user.QueryDaemon().OnlyX(ctx)
		goos := runtime.GOOS
		if goos == "windows" {
			cmd := fmt.Sprintf("Start-Process powershell -WindowStyle Hidden -ArgumentList '%s stop'", daemon.Path)
			command := exec.Command("powershell", "-Command", cmd)
			if err := command.Start(); err != nil {
				fmt.Printf("Error starting command: %s\n", err)
				os.Exit(1)
			}
			if err := command.Wait(); err != nil {
				fmt.Printf("Command finished with error: %s\n", err)
				os.Exit(1)
			}
		} else if goos == "linux" {
			command := exec.Command("sudo", daemon.Path, "stop")
			fmt.Println(command)

			if err := command.Run(); err != nil {
				fmt.Printf("Error starting command on Linux: %s\n", err)
				os.Exit(1)
			}
		}
	},
}

var daemonEnableCmd = &cobra.Command{
	Use:     "enable",
	Short:   "Enable autostart on boot.",
	Long:    `Enable the daemon to autostart on boot.`,
	Aliases: []string{"e"},
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()

		client, err := entclient.New()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer client.Close()

		user := client.User.
			Query().
			Where(
				U.NameEQ(
					ya.GetUser(),
				),
			).
			OnlyX(ctx)

		daemon := user.QueryDaemon().OnlyX(ctx)

		goos := runtime.GOOS
		if goos == "windows" {
			conf := tmpl.YaDaemonPs1{
				TaskName: "YaDaemon",
				Command:  fmt.Sprintf("%s start -d %s", daemon.Path, ya.GetYaDatabasePath()),
			}

			a, err := template.New("a").Parse(tmpl.YaDaemonTpl)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Failed to parse auto start template: %v\n", err)
				os.Exit(1)
			}

			ps1, err := os.CreateTemp("", "*.ps1")
			if err != nil {
				fmt.Fprintf(os.Stderr, "Failed to create temp file for auto start: %v\n", err)
				os.Exit(1)
			}

			if err := a.Execute(ps1, conf); err != nil {
				fmt.Fprintf(os.Stderr, "Failed to execute auto start template: %v\n", err)
				os.Exit(1)
			}

			if err := ps1.Close(); err != nil {
				fmt.Fprintf(os.Stderr, "Failed to close auto start file: %v\n", err)
				os.Exit(1)
			}

			run := fmt.Sprintf("Start-Process -Verb RunAs powershell -WindowStyle Hidden -ArgumentList \"-ExecutionPolicy Bypass\", \"-File `\"%s`\"\"", ps1.Name())

			command := exec.Command("powershell", "-Command", run)
			if err := command.Run(); err != nil {
				fmt.Fprintf(os.Stderr, "Failed to run auto start script: %v\n", err)
				os.Exit(1)
			} else {
				user.QueryDaemon().OnlyX(ctx).Update().SetEnable(true).SaveX(ctx)
			}
		} else if goos == "linux" {
			conf := tmpl.YaDaemonSystemd{
				User:      ya.GetUser(),
				ExecStart: fmt.Sprintf("%s start", daemon.Path),
			}

			a, err := template.New("a").Parse(tmpl.YaDaemonSystemdTpl)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Failed to parse systemd service template: %v\n", err)
				os.Exit(1)
			}

			f, err := os.CreateTemp("", "ya-daemon-*.service")
			if err != nil {
				fmt.Fprintf(os.Stderr, "Failed to create temp file for systemd service: %v\n", err)
				os.Exit(1)
			}
			defer f.Close()
			defer os.Remove(f.Name())

			if err := a.Execute(f, conf); err != nil {
				fmt.Fprintf(os.Stderr, "Failed to execute systemd service template: %v\n", err)
				os.Exit(1)
			}

			sys := "/etc/systemd/system/ya-daemon.service"
			cpCmd := exec.Command("sudo", "cp", f.Name(), sys)
			if err := cpCmd.Run(); err != nil {
				fmt.Fprintf(os.Stderr, "Failed to copy systemd service file: %v\n", err)
				os.Exit(1)
			}

			if err := exec.Command("sudo", "systemctl", "enable", "ya-daemon.service").Run(); err != nil {
				fmt.Fprintf(os.Stderr, "Failed to enable ya-daemon service: %v\n", err)
				os.Exit(1)
			} else {
				daemonRunCmd.Run(cmd, args)
				user.QueryDaemon().OnlyX(ctx).Update().SetEnable(true).SaveX(ctx)
			}
		}
	},
}

var daemonDisableCmd = &cobra.Command{
	Use:     "disable",
	Short:   "Disable daemon autostart on boot",
	Long:    `Disable the daemon from autostart on boot.`,
	Aliases: []string{"d"},
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		client, err := entclient.New()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer client.Close()

		user := client.User.
			Query().
			Where(
				U.NameEQ(
					ya.GetUser(),
				),
			).
			OnlyX(ctx)

		daemon := user.QueryDaemon().OnlyX(ctx)

		goos := runtime.GOOS
		if goos == "windows" {
			name := "YaDaemon"
			cmd := fmt.Sprintf("start-Process powershell -WindowStyle Hidden -ArgumentList 'Unregister-ScheduledTask -TaskName %s -Confirm:$false' -Verb RunAs", name)
			fmt.Println(cmd)

			command := exec.Command("powershell", "-Command", cmd)
			if err := command.Start(); err != nil {
				fmt.Printf("Error starting PowerShell command: %s\n", err)
				os.Exit(1)
			}
			if err := command.Wait(); err != nil {
				fmt.Printf("PowerShell command finished with error: %s\n", err)
				os.Exit(1)
			}
			daemon.Update().SetEnable(false).ExecX(ctx)
		} else if goos == "linux" {
			if err := exec.Command("sudo", "systemctl", "disable", "ya-daemon.service").Run(); err != nil {
				fmt.Fprintf(os.Stderr, "Failed to disable ya-daemon service: %v\n", err)
				os.Exit(1)
			} else {
				user.QueryDaemon().OnlyX(ctx).Update().SetEnable(false).SaveX(ctx)
			}
		}
	},
}

var daemonRestartCmd = &cobra.Command{
	Use:   "restart",
	Short: "Restart the daemon.",
	Long:  `Restart the daemon service.`,
	Run: func(cmd *cobra.Command, args []string) {
		daemonRunCmd.Run(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(daemonCmd)
	daemonCmd.AddCommand(daemonRunCmd)
	daemonCmd.AddCommand(daemonKillCmd)
	daemonCmd.AddCommand(daemonEnableCmd)
	daemonCmd.AddCommand(daemonDisableCmd)
	daemonCmd.AddCommand(daemonRestartCmd)
}
