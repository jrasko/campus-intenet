package confwriter

import (
	"bytes"
	"log"
	"net/http"
)

const (
	baseURL = "http://dhcp4-server"
	port    = "8000"
)

func reloadConfig() {
	resp, err := http.DefaultClient.Post(
		baseURL+":"+port,
		"application/json",
		bytes.NewBufferString(`{"command":"config-reload", "service": ["dhcp4"]}`),
	)

	if err != nil {
		log.Printf("error when sending 'config-reload' to kea: %v", err)
		return
	}

	if resp.StatusCode != http.StatusOK {
		log.Printf("received unexpected status %d", resp.StatusCode)
	}
}
