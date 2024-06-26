package main

import (
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	configLocation, err := getConfig()
	ErrorCheck(err)
	// godotenv loads environment variables from a file but does
	//not replace them if they are in the environment
	err = godotenv.Load(configLocation)
	ErrorCheck(err)
	vocabDB := os.Getenv("ANN2HTML_VOCABDB")
	ErrorCheck(checkEnvVar(vocabDB, "ANN2HTML_VOCABDB"))
	lastNumber := os.Getenv("ANN2HTML_NUM")
	ErrorCheck(checkEnvVar(lastNumber, "ANN2HTML_NUM"))
	templateFile := os.Getenv("ANN2HTML_TEMPLATE")
	ErrorCheck(checkEnvVar(templateFile, "ANN2HTML_TEMPLATE"))
	outFile := os.Getenv("ANN2HTML_OUTPUT")
	ErrorCheck(checkEnvVar(outFile, "ANN2HTML_OUTPUT"))
	languages := os.Getenv("ANN2HTML_LNG")
	vocabMap, newNum, err := databaseToMap(vocabDB, languages, lastNumber)
	ErrorCheck(err)
	err = buildFile(vocabMap, templateFile, outFile)
	ErrorCheck(err)
	SetNum(newNum, configLocation)
	ErrorCheck(err)

}
func checkEnvVar(envVar, envName string) error {
	//Check if environmentable variable is set and if not return
	//error
	if envVar == "" {
		return errors.New(envName + " is not set")
	}
	return nil
}

func ErrorCheck(err error) {
	//check if function throws error and if it does display it and exit the program
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
