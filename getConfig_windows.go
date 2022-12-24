package main

import (
	"errors"
	"os"
)

func getConfig() (string, error) {
	if os.Getenv("ANN2HTML_CONFIG") != "" {
		path := os.Getenv("ANN2HTML_CONFIG")
		if fileExists(path) != false {
			//	file found in ANN2HTML_CONFIG
			return path, nil
		} else {
			return "", errors.New("ANN2HTML_CONFIG set but no config file found")
		}
	}
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", errors.New("Could not find user config directory")
	}
	path := configDir + filepath.FromSlash("/ann2html/config")
	if fileExists(path) != false {
		return path, nil
	}
	return "", errors.New("Configuration file not found in " + path)

}
