package adapter

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func (s *SingBox) Start() {
	provider := s.user.QueryProvider().OnlyX(s.ctx)
	subscribe := s.user.QuerySubscribe().OnlyX(s.ctx)

	// Inline runProvider logic
	goos := runtime.GOOS
	switch goos {
	case "windows":
		run := fmt.Sprintf(
			`Start-Process powershell -WindowStyle Hidden -ArgumentList "%s run -c %s" -Verb RunAs`,
			provider.Path,
			subscribe.Conf,
		)
		fmt.Println(run)
		command := exec.Command("powershell", "-Command", run)
		if err := command.Run(); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to run script: %v\n", err)
			os.Exit(1)
		}

	case "linux":
		run := fmt.Sprintf("%s run -c %s", provider.Path, subscribe.Conf)

		sh := fmt.Sprintf(`pid=$(pgrep -f "%s")
if [ ! -z "$pid" ]; then
    kill $pid
    sleep 3
    if kill -0 $pid > /dev/null 2>&1; then
        kill -9 $pid
        sleep 2
    fi
fi
nohup %s > /dev/null 2>&1 &
`, provider.Name, run)

		f, err := os.CreateTemp("", "run-*.sh")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to create temp file: %v\n", err)
			os.Exit(1)
		}
		defer os.Remove(f.Name())

		if _, err := f.WriteString(sh); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to write to temp file: %v\n", err)
			os.Exit(1)
		}
		if err := f.Close(); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to close temp file: %v\n", err)
			os.Exit(1)
		}

		command := exec.Command("bash", f.Name())
		if err := command.Run(); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to run provider on Linux: %v\n", err)
			os.Exit(1)
		}
	}
}

func (s *SingBox) Stop() {
	provider := s.user.QueryProvider().OnlyX(s.ctx)
	name := provider.Name
	goos := runtime.GOOS
	if goos == "windows" {
		cmd := fmt.Sprintf(`Start-Process powershell -WindowStyle Hidden -ArgumentList "Stop-Process -Name %s -Force" -Verb runAs`, name)
		fmt.Println(cmd)

		command := exec.Command("powershell", "-Command", cmd)
		if err := command.Run(); err != nil {
			fmt.Printf("Failed to run command on Windows: %s\n", err)
			os.Exit(1)
		}
	} else if goos == "linux" {
		cmd := fmt.Sprintf("pkill -f %s", name)
		fmt.Println(cmd)

		command := exec.Command("bash", "-c", cmd)
		if err := command.Start(); err != nil {
			fmt.Printf("Failed to run command on Linux: %s\n", err)
			os.Exit(1)
		}
	}
}
