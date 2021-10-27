package main

import (
	"encoding/json"
	"os"

	"gopkg.in/yaml.v2"
)

func main() {
	var c interface{}
	json.NewDecoder(os.Stdin).Decode(&c)
	yaml.NewEncoder(os.Stdout).Encode(&c)
}
