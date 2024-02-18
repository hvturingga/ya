package tests

import (
	"os/user"
	"runtime"
	"testing"
)

func Test_GetOSArch(t *testing.T) {
	os, arch := runtime.GOOS, runtime.GOARCH
	t.Logf("OS: %s, Arch: %s", os, arch)
}

func Test_GetUserProfileDir(t *testing.T) {
	dir, err := user.Current()
	if err != nil {
		t.Logf("Error: %s", err)
		return
	}

	t.Logf("User Profile Dir: %s", dir.HomeDir)
}

func Test_GetCurrentUser(t *testing.T) {
	currentUser, err := user.Current()
	if err != nil {
		t.Errorf("failed to get current user: %v", err)
		return
	}

	t.Logf("current user name: %s", currentUser.Username)
}
