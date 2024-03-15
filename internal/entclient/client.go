package entclient

import (
	"context"
	"fmt"
	"github.com/hvturingga/ya/ent"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"os/user"
	"path/filepath"
)

func New(address ...string) (*ent.Client, error) {
	var path string
	if len(address) > 0 {
		path = address[0]
	} else {
		usr, err := user.Current()
		if err != nil {
			return nil, fmt.Errorf("failed to get current user: %w", err)
		}
		home := filepath.Join(usr.HomeDir, ".ya")
		path = filepath.Join(home, "ya.db")

		if err := os.MkdirAll(home, 0755); err != nil {
			return nil, fmt.Errorf("failed to create .ya directory: %w", err)
		}
	}

	client, err := ent.Open("sqlite3", fmt.Sprintf("file:%s?_fk=1", path))
	if err != nil {
		return nil, fmt.Errorf("failed to open client: %w", err)
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		client.Close()
		return nil, fmt.Errorf("failed to create schema: %w", err)
	}

	return client, nil
}
