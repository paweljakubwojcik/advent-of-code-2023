package main

import (
	"strconv"
	"strings"
	u "utils"
)

func getGameMinimalNumberOfBalls(game string) map[string]int {
	config := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}

	for _, set := range strings.Split(game, ";") {
		for _, ballData := range strings.Split(set, ",") {
			n := strings.Split(strings.TrimSpace(ballData), " ")

			num, err := strconv.Atoi(n[0])
			if err != nil {
				panic(err)
			}

			color := n[1]

			if config[color] < num {
				config[color] = num
			}
		}
	}
	return config
}

func main() {
	games := u.ReadLinesFromFile("/input.txt")

	sum := 0
	for _, game := range games {
		g := strings.Split(game, ":")
		gameContent := g[1]

		gameMinimalNumber := getGameMinimalNumberOfBalls(gameContent)

		power := gameMinimalNumber["red"] * gameMinimalNumber["blue"] * gameMinimalNumber["green"]
		sum += power
	}
	println(sum)

}
