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
		return "", errors.New("unable to find config directory")
	}
	path := configDir + "/ann2html/config"
	if fileExists(path) != false {
		return path, nil
	}
	return "", errors.New("No configuration file found at " + path)

}
