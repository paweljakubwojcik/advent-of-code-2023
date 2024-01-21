package main

import (
	"fmt"
	"os"
	"strings"
)

const FILE = "/input.txt"

var digitsToIntMap = map[string]int{
	"one": 1,
	"two": 2,
	"three": 3,
	"four": 4,
	"five": 5,
	"six": 6,
	"seven": 7,
	"eight": 8,
	"nine": 9,
	"1": 1,
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
}

// day 1
func main() {

	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	dat, err := os.ReadFile(path + FILE)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(dat), "\n")
	
	sum := 0
	for _, elem := range lines {
		
		// this approach accounts for overlapping strings
		// ex. threetwone should yeld 31, but using regex approach it produces 32

		var first int
		var last int
		var firstIndex int = len(elem)
		var lastIndex int = 0

		for key, val := range digitsToIntMap {
			
			newLastIndex := strings.LastIndex(elem, key)
			if newLastIndex != -1 && newLastIndex >= lastIndex {
				lastIndex = newLastIndex
				last = val
			}

			newFirstIndex := strings.Index(elem, key)
			if newFirstIndex != -1 && newFirstIndex <= firstIndex {
				firstIndex = newFirstIndex
				first = val
			}
		}

		fmt.Println(elem, first, last)

		code := first * 10 + last

		sum += code
	}

	fmt.Println(sum)

}
