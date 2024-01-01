package main

import (
	"fmt"
	"os"
	"regexp"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	contents, err := os.ReadDir(".")
	check(err)

	fmt.Println("Adding files...")

	// Define a regular expression to match .mkv files
	mkvRegex := regexp.MustCompile(`\.mkv$`)

 	srtRegex := regexp.MustCompile(`\.srt$`)

	var videoNames, subtitleNames []string

	for _, entry := range contents {
		if entry.IsDir() {
			continue // Skip directories
		}

		fileName := entry.Name()

		// Check if the file name matches the pattern
		if mkvRegex.MatchString(fileName) {
			videoNames = append(videoNames, fileName)
		} else if srtRegex.MatchString(fileName) {
			subtitleNames = append(subtitleNames, fileName)
		}
	}

	for i := 0; i < len(videoNames); i++ {
		base := videoNames[i][:len(videoNames[i]) - len(".mkv")]
		newName := base + ".srt"
		os.Rename(subtitleNames[i], newName)
	}
}

