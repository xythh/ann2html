package main

import (
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

//boundle script to automate this on linux, or figure out mounting solution

func main() {
	configLocation, err := getConfig()
	ErrorCheck(err)
	err = godotenv.Load(configLocation)
	//Valid config locations found in getConfig_{OS}
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
	if envVar == "" {
		return errors.New(envName + " is not set")
	}
	return nil
}

func ErrorCheck(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
