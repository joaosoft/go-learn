package config

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/gommon/log"
	"os"
)

// LoadConfigFromFile loads a configuration from a file path.
func LoadConfigFromFile(contextName string, obj interface{}) error {
	root := getWorkingDir()

	var err error
	file, err := os.Open(fmt.Sprintf("%s/examples/52_nsq/config/%s.json", root, contextName))
	if err != nil {
		return err
	}

	decoder := json.NewDecoder(file)
	if err = decoder.Decode(obj); err != nil {
		return err
	}

	return nil
}

// GetWorkingDir returns the current working directory
func getWorkingDir() string {
	dir, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	return dir
}
