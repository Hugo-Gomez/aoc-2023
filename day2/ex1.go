package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	total := 0

	red := 12
	green := 13
	blue := 14

	file, err := os.Open("./input")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		game := strings.Split(line, ": ")
		gameNb := strings.Split(game[0], " ")[1]
		sets := strings.Split(game[1], "; ")
		validGame := true
		for _, set := range sets {
			cubes := strings.Split(set, ", ")
			for _, cube := range cubes {
				strNumber := strings.Split(cube, " ")[0]
				number, err := strconv.Atoi(strNumber)
				check(err)
				color := strings.Split(cube, " ")[1]
				switch color {
				case "red":
					if number > red {
						validGame = false
					}
				case "green":
					if number > green {
						validGame = false
					}
				case "blue":
					if number > blue {
						validGame = false
					}
				}
			}
		}
		if validGame {
			gameNumber, err := strconv.Atoi(gameNb)
			check(err)
			total += gameNumber
		}
	}
	fmt.Println(total)

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
