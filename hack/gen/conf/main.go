package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/hvturingga/ya/conf"
)

const constantsTpl = `// Code generated by go generate; DO NOT EDIT.
package conf

const (
	ClashAPI = "{{ .ClashAPI }}"
)
`

const providerListTpl = `// Code generated by go generate; DO NOT EDIT.
package conf

func GetProviderList() []Provider {
	return []Provider{
{{- range . }}
		{
			Repo: "{{ .Repo }}", 
			Owner: "{{ .Owner }}", 
			Version: []string{
				{{- range .Version }}
				"{{ . }}",
				{{- end }}
			},
		},
{{- end }}
	}
}
`

func gen(wd, path, tpl string, data interface{}) error {
	t, err := template.New("").Parse(tpl)
	if err != nil {
		return err
	}

	fp := filepath.Join(wd, path)
	f, err := os.Create(fp)
	if err != nil {
		return err
	}
	defer f.Close()

	return t.Execute(f, data)
}

func read[T any](path string) (T, error) {
	var cfg T
	b, err := os.ReadFile(path)
	if err != nil {
		return cfg, err
	}
	err = json.Unmarshal(b, &cfg)
	return cfg, err
}

func main() {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("Failed to get working directory:", err)
		return
	}

	if err := genConstants(wd); err != nil {
		fmt.Printf("generating constants failed: %w", err)
		return
	}
	if err := genProviderLists(wd); err != nil {
		fmt.Printf("generating providers failed: %w", err)
		return
	}
}

func genConstants(wd string) error {
	config, err := read[conf.Conf](filepath.Join(wd, "conf", "meta", "conf.json"))
	if err != nil {
		return err
	}

	return gen(wd, "conf/constants.go", constantsTpl, config)
}

func genProviderLists(wd string) error {
	providers, err := read[[]conf.Provider](filepath.Join(wd, "conf", "meta", "provider.json"))
	if err != nil {
		return err
	}

	return gen(wd, "conf/provider_list.go", providerListTpl, providers)
}
