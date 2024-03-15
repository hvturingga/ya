package main

import (
	"fmt"
	"os"
	"text/template"
)

type Arch struct {
	OS   string
	Arch string
	Exe  string
}

func main() {
	archs := []Arch{
		{OS: "linux", Arch: "amd64", Exe: ""},
		{OS: "linux", Arch: "arm64", Exe: ""},
		{OS: "windows", Arch: "amd64", Exe: ".exe"},
	}

	embedTmpl := `//go:build {{.OS}}
// +build {{.OS}}

package embed

import _ "embed"

//go:embed ya-daemon-{{.OS}}-{{.Arch}}{{.Exe}}
var Daemon []byte
`
	tmpl, err := template.New("embed").Parse(embedTmpl)
	if err != nil {
		panic(err)
	}

	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	for _, arch := range archs {
		name := "daemon_embed_" + arch.OS + "_" + arch.Arch + ".go"
		create, err := os.Create(fmt.Sprintf("%s/cmd/cli/embed/%s", wd, name))
		if err != nil {
			panic(err)
		}
		defer create.Close()
		if err := tmpl.Execute(create, arch); err != nil {
			panic(err)
		}
	}
}
