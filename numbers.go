package main

import (
	"errors"
	"io/ioutil"
	"regexp"
	"strings"
)

func SetNum(num string, config string) error {
	file, err := ioutil.ReadFile(config)
	if err != nil {
		return errors.New("Unable to read config file to write new number to")

	}
	lines := strings.Split(string(file), "\n")
	for i, _ := range lines {
		if strings.Contains(lines[i], "ANN2HTML_NUM=") {
			replace := regexp.MustCompile(`ANN2HTML_NUM=\d+`)
			s := replace.ReplaceAllString(lines[i], "ANN2HTML_NUM="+num)
			lines[i] = s
			output := strings.Join(lines, "\n")
			err = ioutil.WriteFile(config, []byte(output), 0644)
			return errors.New("Unable to write new number to config file")
		}
	}
	return nil
}
