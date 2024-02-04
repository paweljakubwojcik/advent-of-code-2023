package main

import (
	"encoding/json"
	"fmt"
	"math"
	"strconv"
	"unicode"
	u "utils"

	"github.com/golang-collections/collections/set"
)

type rangePosition struct {
	X [2]int
	Y int
}

type partNumber struct {
	Value      string
	Position   rangePosition
	IsAdjacent bool
}

type position struct {
	X int
	Y int
}

type partSymbol struct {
	Value           rune
	Position        position
	AdjacentIndexes *set.Set
}

func isAdjacent(numP rangePosition, symP position) bool {
	if math.Abs(float64(symP.Y-numP.Y)) > 1 {
		return false
	}

	if symP.X >= numP.X[0]-1 && symP.X <= numP.X[1]+1 {
		return true
	}

	return false
}

func main() {
	input := u.ReadLinesFromFile("/input.txt")

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

			newPart := partSymbol{Value: char, Position: position{X: xIndex, Y: yIndex}, AdjacentIndexes: set.New()}
			partsStack = append(partsStack, newPart)

			lastCharWasNumber = false
		}
		// done constructing stacks for line y

		for i := len(numbersStack) - 1; i >= 0; i-- {
			num := &numbersStack[i]
			if num.Position.Y < yIndex {
				break
			}

			for j := len(partsStack) - 1; j >= 0; j-- {
				sym := partsStack[j]
				if sym.Position.Y < yIndex-1 {
					break
				}

				if isAdjacent(num.Position, sym.Position) {
					num.IsAdjacent = true
					sym.AdjacentIndexes.Insert(i)
				}

			}
		}

		for i := len(partsStack) - 1; i >= 0; i-- {
			sym := partsStack[i]
			if sym.Position.Y < yIndex {
				break
			}

			for j := len(numbersStack) - 1; j >= 0; j-- {
				num := &numbersStack[j]
				if num.Position.Y < yIndex-1 {
					break
				}

				if isAdjacent(num.Position, sym.Position) {
					num.IsAdjacent = true
					sym.AdjacentIndexes.Insert(j)
				}
			}
		}

	}

	sum := 0
	for _, part := range partsStack {
		if part.Value == '*' && part.AdjacentIndexes.Len() == 2 {
			product := 1
			part.AdjacentIndexes.Do(func(i interface{}) {
				index := i.(int)
				n := numbersStack[index]
				value, _ := strconv.Atoi(n.Value)
				product *= value
			})
			sum += product
		}
	}

	jsonsed, err := json.MarshalIndent(partsStack, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(jsonsed))
	fmt.Println(sum)

}
