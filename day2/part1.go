package main

import (
	"strconv"
	"strings"
	u "utils"
)

var GAME_CONFIG = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func isGamePossible(game string) bool {
	for _, set := range strings.Split(game, ";") {
		for _, ballData := range strings.Split(set, ",") {
			n := strings.Split(strings.TrimSpace(ballData), " ")
			
			num, err := strconv.Atoi(n[0])
			if err != nil {
				panic(err)
			}

			color := n[1]
			if GAME_CONFIG[color] < num {
				return false
			}
		}
	}
	return true

}

func part1() {
	games := u.ReadLinesFromFile("/input.txt")

	sum := 0
	for _, game := range games {
		g := strings.Split(game, ":")
		gameName := g[0]
		gameContent := g[1]
		if isGamePossible(gameContent) {
			n := strings.Split(strings.TrimSpace(gameName), " ")
			gameIndex, err := strconv.Atoi(n[1])
			if err != nil {
				panic(err)
			}
			sum += gameIndex
		}
	}
	println(sum)

}
