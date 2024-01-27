package utils

import (
	"os"
	"strings"
)

func ReadInputFile(file string) []byte {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	dat, err := os.ReadFile(path + file)
	if err != nil {
		panic(err)
	}

	return dat
}

func ReadLinesFromFile(file string) []string {
	data := ReadInputFile(file)
	lines := strings.Split(string(data), "\n")

	return lines
}
