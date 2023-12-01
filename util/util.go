package util

import (
	"log"
	"os"
	"strings"
)

func ReadLines(path string) ([]string, error) {
	// read file from path, return utf-8 string of contents
	contents, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(contents), "\n")

	return lines, nil
}
