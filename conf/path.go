package conf

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
)

func GetYaPath() string {
	c, err := user.Current()
	if err != nil {
		fmt.Printf("Failed to get the user's home directory: %s\n", err)
		os.Exit(1)
	}
	return filepath.Join(c.HomeDir, ".ya")
}

func GetProviderPath() string {
	return filepath.Join(GetYaPath(), "provider")
}

func GetDaemonPath() string {
	return filepath.Join(GetYaPath(), "daemon")
}

func GetSubscribePath() string {
	return filepath.Join(GetYaPath(), "subscribe")
}

func GetDatabasePath() string {
	return filepath.Join(GetYaPath(), "ya.db")
}
