package nodeswitch

import (
	"bytes"
	"fmt"
	"github.com/hvturingga/ya/conf"
	"io"
	"net/http"
	"os"
	"time"
)

type Node struct {
	group string
	name  string
}

func New(group, name string) *Node {
	return &Node{
		group: group,
		name:  name,
	}
}

func (n *Node) Switch() error {
	url := fmt.Sprintf("http://%s/proxies/%s", conf.ClashAPI, n.group)
	payload := bytes.NewBuffer([]byte(`{"name": "` + n.name + `"}`))
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	put, err := http.NewRequest(http.MethodPut, url, payload)
	if err != nil {
		return err
	}
	put.Header.Add("Content-Type", "application/json")

	r, err := client.Do(put)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	if r.StatusCode != http.StatusNoContent {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(string(body))
		return fmt.Errorf("status: %s", r.Status)
	}
	return nil
}
