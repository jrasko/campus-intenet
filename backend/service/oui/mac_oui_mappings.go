package oui

import (
	_ "embed"
	"encoding/json"
	"log"
)

//go:embed oui_mappings.json
var ouiFile json.RawMessage

var Mappings map[string]string

func init() {
	err := json.Unmarshal(ouiFile, &Mappings)
	if err != nil {
		log.Fatalf("could not parse oui file: %w", err)
	}
}
