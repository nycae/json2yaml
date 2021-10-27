package main

import (
	"encoding/json"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

func main() {
	var c interface{}
	if err := json.NewDecoder(os.Stdin).Decode(&c); err != nil {
		log.Fatal(err)
	}
	if err := yaml.NewEncoder(os.Stdout).Encode(&c); err != nil {
		log.Fatal(err)
	}
}
