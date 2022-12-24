package main

import (
	"fmt"
	"os"
	"errors"
)

func buildFile(databaseMap []string, template string, outFile string) error {
	//erase old file, copy template file to replace it and add new annotations
	//this is kinda scuffed, probably want to change how this is done, in a way
	// that the program is aware of html format.
	os.Remove(outFile)
	_, err := copyFile(template, outFile)
	if err != nil {
		return errors.New("Failure to copy template to output file")
	}
	f, err := os.OpenFile(outFile, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return errors.New("Failure to open newly created output file")
	}
	for _, k := range databaseMap {
		_, err = fmt.Fprintln(f, "<p>", k, "</p>")
		if err != nil {
		return errors.New("Failure to append to output file")
		}
	}
	fmt.Fprintln(f, "</body>")
	fmt.Fprintln(f, "</html>")
	return nil
}

func copyFile(inputFile, outputFile string) (int64, error) {
	i, e := os.Open(inputFile)
	if e != nil {
		return 0, e
	}
	defer i.Close()
	//	os.Remove(outputFile)// possible fix
	o, e := os.Create(outputFile)
	if e != nil {
		return 0, e
	}
	defer o.Close()
	return o.ReadFrom(i)
}
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
