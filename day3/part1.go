package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"unicode"
	u "utils"
)

// type rangePosition struct {
// 	X [2]int
// 	Y int
// }

// type partNumber struct {
// 	Value      string
// 	Position   rangePosition
// 	IsAdjacent bool
// }

// type position struct {
// 	X int
// 	Y int
// }

// type partSymbol struct {
// 	Value    rune
// 	Position position
// }

// func isAdjacent(numP rangePosition, symP position) bool {
// 	if math.Abs(float64(symP.Y-numP.Y)) > 1 {
// 		return false
// 	}

// 	if symP.X >= numP.X[0]-1 && symP.X <= numP.X[1]+1 {
// 		return true
// 	}

// 	return false
// }

func part1() {
	input := u.ReadLinesFromFile("/input.txt")

	// sum := 0
	numbersStack := []partNumber{}
	partsStack := []partSymbol{}

	for yIndex, line := range input {
		lastCharWasNumber := false

		for xIndex, char := range []rune(line) {
			if char == '.' {
				lastCharWasNumber = false
				continue
			}

			if unicode.IsDigit(char) {
				var lastNumber *partNumber

				if len(numbersStack) == 0 || !lastCharWasNumber {
					newPartNumber := partNumber{Value: string(char), Position: rangePosition{X: [...]int{xIndex, xIndex}, Y: yIndex}}
					numbersStack = append(numbersStack, newPartNumber)
					lastNumber = &newPartNumber
				} else {
					lastNumber = &numbersStack[len(numbersStack)-1]
					lastNumber.Value += string(char)
					lastNumber.Position.X[1] = xIndex
				}
				lastCharWasNumber = true

				continue
			}

			newPart := partSymbol{Value: char, Position: position{X: xIndex, Y: yIndex}}
			partsStack = append(partsStack, newPart)

			lastCharWasNumber = false
		}
		// done constructing stacks for line y

		for i := len(numbersStack) - 1; i >= 0; i-- {
			num := &numbersStack[i]
			if num.Position.Y < yIndex {
				break
			}

			for i := len(partsStack) - 1; i >= 0; i-- {
				sym := partsStack[i]
				if sym.Position.Y < yIndex-1 {
					break
				}

				if !num.IsAdjacent && isAdjacent(num.Position, sym.Position) {
					num.IsAdjacent = true
				}

			}
		}

		for i := len(partsStack) - 1; i >= 0; i-- {
			sym := partsStack[i]
			if sym.Position.Y < yIndex {
				break
			}

			for i := len(numbersStack) - 1; i >= 0; i-- {
				num := &numbersStack[i]
				if num.Position.Y < yIndex-1 {
					break
				}

				if !num.IsAdjacent && isAdjacent(num.Position, sym.Position) {
					num.IsAdjacent = true
				}
			}
		}

	}

	sum := 0
	for _, num := range numbersStack {
		if num.IsAdjacent {
			valueAsNum, _ := strconv.Atoi(num.Value)
			sum += valueAsNum
		}
	}

	jsonsed, err := json.MarshalIndent(numbersStack, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(jsonsed))
	fmt.Println(sum)

}
