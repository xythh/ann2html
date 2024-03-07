package main

import (
	"errors"
	"os"
	"fmt"
	"path/filepath"
)

func getConfig() (string, error) {
	//Get config from environment variable
	if os.Getenv("ANN2HTML_CONFIG") != "" {
		path := os.Getenv("ANN2HTML_CONFIG")
		if fileExists(path) != false {
			return path, nil
		} else {
			return "", errors.New("ANN2HTML_CONFIG set but no config file found")
		}
	}
	// Get location of executable, and search for config in that location.
	ex, err := os.Executable()
	if err != nil {
		return "", errors.New("Unable to find executable")
	}
	configDir := filepath.Dir(ex)
	fmt.Println(configDir)
	path := configDir + filepath.FromSlash("/config")
	if fileExists(path) != false {
		return path, nil
	}
	return "", errors.New("No configuration file found at " + path)

}
