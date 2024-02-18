package provider

import (
	"bufio"
	"context"
	"fmt"
	"github.com/google/go-github/v59/github"
	"github.com/hvturingga/ya/conf"
	"github.com/hvturingga/ya/ent"
	"github.com/hvturingga/ya/internal/unzip"
	"github.com/hvturingga/ya/internal/ya"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

func Select(client *ent.Client, U *ent.User, sel conf.Provider) {

	ctx := context.Background()
	gh := github.NewClient(nil)

	list, _, err := gh.Repositories.GetReleaseByTag(ctx, sel.Owner, sel.Repo, sel.Version)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	goos, arch := runtime.GOOS, runtime.GOARCH

	var assets []*github.ReleaseAsset
	for _, asset := range list.Assets {
		if strings.Contains(asset.GetName(), goos) && strings.Contains(asset.GetName(), arch) {
			assets = append(assets, asset)
		}
	}
	if len(assets) == 0 {
		fmt.Println("No assets found for the current OS and architecture.")
		os.Exit(1)
	}

	for i, asset := range assets {
		fmt.Printf("%d. %s \n", i+1, asset.GetName())
	}

	fmt.Print("Please enter the number of the asset to download: ")

	reader := bufio.NewReader(os.Stdin)
	number, _ := reader.ReadString('\n')
	ach, err := strconv.Atoi(strings.TrimSpace(number))
	if err != nil || ach < 1 || ach > len(list.Assets) {
		fmt.Println("Invalid selection, please rerun the command")
		os.Exit(1)
	}

	url := assets[ach-1].GetBrowserDownloadURL()

	download := "./.download"

	if err := os.MkdirAll(download, 0755); err != nil {
		fmt.Printf("Failed to create temp directory: %v", err)
		os.Exit(1)
	}
	defer func(path string) {
		err := os.RemoveAll(path)
		if err != nil {
			fmt.Printf("Failed to remove temp directory: %v", err)
			os.Exit(1)
		}
	}(download)

	src := fmt.Sprintf("%s/%s", download, filepath.Base(url))

	f, err := os.Create(src)
	if err != nil {
		fmt.Printf("Failed to create file: %v", err)
		os.Exit(1)
	}
	defer f.Close()

	r, err := http.Get(url)
	if err != nil {
		fmt.Printf("Failed to download file: %v", err)
		os.Exit(1)
	}
	defer r.Body.Close()

	if r.StatusCode != http.StatusOK {
		fmt.Printf("Server returned %d status", r.StatusCode)
		os.Exit(1)
	}

	p := &progressReader{reader: r.Body, total: r.ContentLength}

	if _, err := io.Copy(f, p); err != nil {
		fmt.Printf("Failed to write file: %v", err)
		os.Exit(1)
	}

	if filepath.Ext(src) == ".zip" {
		if err := unzip.Unzip(src, download); err != nil {
			fmt.Printf("Failed to unzip: %v", err)
			os.Exit(1)
		}
	} else if filepath.Ext(src) == ".gz" {
		if err := unzip.Untargz(src, download); err != nil {
			fmt.Printf("Failed to untargz: %v", err)
			os.Exit(1)
		}
	}
	s, err := search(download, sel.Repo)
	if err != nil {
		fmt.Printf("Failed to search for the file: %s\n", err)
		os.Exit(1)
	}
	bin := filepath.Join(ya.GetBinDir(), filepath.Base(s))
	if err := mv(s, bin); err != nil {
		fmt.Printf("Failed to move the file: %s\n", err)
		os.Exit(1)
	}

	create, err := client.Provider.Create().
		SetName(sel.Repo).
		SetVersion(sel.Version).
		SetPath(bin).
		Save(ctx)
	if err != nil {
		fmt.Printf("Failed to create provider: %s\n", err)
		os.Exit(1)
	}

	U.Update().ClearSubscribe().SetProvider(create).ExecX(ctx)

	fmt.Printf("Provider %s %s initialized\n", create.Name, create.Version)

}
