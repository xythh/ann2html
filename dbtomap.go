package main

import (
	"database/sql"
	"errors"
	_ "modernc.org/sqlite"
	"os"
	"strings"
)

func databaseToMap(databaseFile string, lang string, lastTime string) ([]string, string, error) {
	if fileExists(databaseFile) != true {
		return nil, "", errors.New("your vocab.db was not found")
		os.Exit(1)
	}
	db, err := sql.Open("sqlite", databaseFile)
	if err != nil {
		return nil, "", errors.New("error accesing vocab.db")
	}
	defer db.Close()
	rows, err := db.Query("SELECT timestamp,word_key,usage from LOOKUPS where timestamp > " + lastTime + getLanguages(lang))
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
		final = append(final, formatSent(word_key, usage))
	}
	if len(final) == 0 {
		return final, count, errors.New("No new annotations")

	}
	return final, count, nil
}

func formatSent(word_key, sentence string) string {
	//searches for word and in sentence and replaces it with </b>
	// IMPLEMENT TABBING to make mining even easier
	word := strings.Split(word_key, ":")

	replacement := "<b tabindex=\"0\">" + word[1] + "</b>"
	newSentence := strings.Replace(sentence, word[1], replacement, 1)
	return newSentence
}
func getLanguages(annLanguages string) string {
	var andWord string
	var extra string
	if annLanguages == "" {
		return ""
	}

	languages := strings.Split(annLanguages, ",")
	andWord = " AND word_key LIKE '" + languages[0] + ":%'"
	if len(languages) == 1 {
		return andWord
	}

	for _, value := range languages[1:] {
		extra = extra + " OR word_key LIKE '" + value + ":%'"

	}
	return andWord + extra
}
