package main

import (
	"database/sql"
	"errors"
	"golang.org/x/text/language"
	_ "modernc.org/sqlite"
	"os"
	"strconv"
	"strings"
)

func databaseToMap(databaseFile string, lang string, lastTime string) ([]string, string, error) {
	//takes the path to your vocab.db,languages you want to mine from and the highest unix timestamp from the last time
	//the program was ran then returns a map with all your annotations, the highest unix timestamp found and an error.
	if fileExists(databaseFile) != true {
		return nil, "", errors.New("your vocab.db was not found")
		os.Exit(1)
	}
	db, err := sql.Open("sqlite", databaseFile)
	if err != nil {
		return nil, "", errors.New("error accesing vocab.db")
	}
	// make sure lastTime is a number
	_, err = strconv.ParseInt(lastTime, 10, 64)
	if err != nil {
		return nil, "", err
	}

	defer db.Close()
	languages, err := getLanguages(lang)
	if err != nil {
		return nil, "", err
	}
	rows, err := db.Query("SELECT timestamp,word_key,usage from LOOKUPS where timestamp > " + lastTime + languages)
	ErrorCheck(err)
	defer rows.Close()

	var final []string
	var count string
	defer rows.Close()
	for rows.Next() {
		var timestamp, word_key, usage string
		err = rows.Scan(&timestamp, &word_key, &usage)
		if err != nil {
			return nil, "", errors.New("error forming map")

		}
		count = timestamp
		// this count is used to keep track of the last unix timestamp in your kindle database
		final = append(final, formatSent(word_key, usage))
	}
	if len(final) == 0 {
		return final, count, errors.New("No new annotations")

	}
	return final, count, nil
}

func formatSent(word_key, sentence string) string {
	//finds your mined words in your sentences and replace it
	//with <b>word</b>
	word := strings.Split(word_key, ":")

	replacement := "<b>" + word[1] + "</b>"
	newSentence := strings.Replace(sentence, word[1], replacement, 1)
	return newSentence
}
func getLanguages(annLanguages string) (string, error) {
	//splits your ANN2HTML_LNG string into a map with all the languages you want to mine from,
	//then it return a appropriate string to add to the database query it only queries from those
	//languages
	var andWord string
	var extra string
	if annLanguages == "" {
		return "", nil
		//if nothing is set than it mines from all languages.
	}

	languages := strings.Split(annLanguages, ",")
	// make sure our string of languages are valid languages
	for _, v := range languages {
		_, err := language.ParseBase(v)
		if err != nil {
			return "", errors.New(v + "is not a correct ISO 639 language code")
		}
	}
	andWord = " AND word_key LIKE '" + languages[0] + ":%'"
	if len(languages) == 1 {
		return andWord, nil
	}

	for _, value := range languages[1:] {
		extra = extra + " OR word_key LIKE '" + value + ":%'"

	}
	return andWord + extra, nil
}
