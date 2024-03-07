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
	// copies a file by opening your input file and then creating a file exactly like your input file
	// in the directory of your output file.
	i, err := os.Open(inputFile)
	if err != nil {
		return 0, err
	}
	defer i.Close()
		os.Remove(outputFile)// not neccesary on linux but might be on windows 
	o, err := os.Create(outputFile)
	if err != nil {
		return 0, err
	}
	defer o.Close()
	return o.ReadFrom(i)
}
func fileExists(filename string) bool {
	// Check stats of the file to make sure it exists.
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
