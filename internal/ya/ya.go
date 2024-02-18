package ya

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

func GetBinDir() string {
	usr, err := user.Current()
	if err != nil {
		fmt.Printf("Failed to get the user's home directory: %s\n", err)
		return ""
	}
	bin := filepath.Join(usr.HomeDir, ".ya/bin")
	if err := os.MkdirAll(bin, 0755); err != nil {
		fmt.Printf("Failed to create the destination directory: %s\n", err)
		return ""
	}
	return bin
}

func GetSubscribeDir() string {
	usr, err := user.Current()
	if err != nil {
		fmt.Printf("Failed to get the user's home directory: %s\n", err)
		return ""
	}
	sub := filepath.Join(usr.HomeDir, ".ya/subscribe")
	if err := os.MkdirAll(sub, 0755); err != nil {
		fmt.Printf("Failed to create the destination directory: %s\n", err)
		return ""
	}
	return sub
}

func GetDaemonDir() string {
	usr, err := user.Current()
	if err != nil {
		fmt.Printf("Failed to get the user's home directory: %s\n", err)
		return ""
	}
	script := filepath.Join(usr.HomeDir, ".ya/daemon")
	if err := os.MkdirAll(script, 0755); err != nil {
		fmt.Printf("Failed to create the destination directory: %s\n", err)
		return ""
	}
	return script
}

func GetYaDatabasePath() string {
	usr, err := user.Current()
	if err != nil {
		fmt.Printf("Failed to get the user's home directory: %s\n", err)
		return ""
	}
	home := filepath.Join(usr.HomeDir, ".ya")
	path := filepath.Join(home, "ya.db")
	return path
}

func GetUser() string {
	c, err := user.Current()
	if err != nil {
		fmt.Printf("Failed to get the user's home directory: %s\n", err)
		return ""
	}
	r := c.Username
	parts := strings.Split(r, "\\")
	if len(parts) > 0 {
		r = parts[len(parts)-1]
	}
	return r
}
