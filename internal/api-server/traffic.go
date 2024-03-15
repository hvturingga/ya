package api_server

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/hvturingga/ya/conf"
	"log"
	"net/http"
	"os"
)

type TrafficData struct {
	Up   int `json:"up"`
	Down int `json:"down"`
}

func Traffic() {
	url := fmt.Sprintf("http://%s/traffic", conf.ClashAPI)

	r, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()

	reader := bufio.NewReader(r.Body)

	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			log.Fatal(err)
		}

		var data TrafficData
		if err := json.Unmarshal(line, &data); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("\rUp: %d bytes, Down: %d bytes    ", data.Up, data.Down)
		os.Stdout.Sync()
	}
}
