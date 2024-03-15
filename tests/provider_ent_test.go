package tests

import (
	"context"
	"fmt"
	"github.com/hvturingga/ya/internal/entclient"
	"os/user"
	"testing"
)

func Test_ProviderCreate(t *testing.T) {
	client, err := entclient.New()
	if err != nil {
		t.Errorf("Error: %s", err)
		return
	}
	defer client.Close()

	current, err := user.Current()
	if err != nil {
		t.Errorf("Error: %s", err)
		return
	}

	if err := client.Provider.Create().
		SetName("clash").
		SetVersion("v1.9.0-alpha.9").
		SetPath(
			fmt.Sprintf("%s/.ya/clash", current.HomeDir),
		).Exec(context.Background()); err != nil {
		t.Errorf("Error: %s", err)
		return
	}
}
