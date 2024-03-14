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

	file, err := os.Open("./input")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		redMin := 0
		greenMin := 0
		blueMin := 0
		line := scanner.Text()
		game := strings.Split(line, ": ")
		sets := strings.Split(game[1], "; ")
		for _, set := range sets {
			cubes := strings.Split(set, ", ")
			for _, cube := range cubes {
				strNumber := strings.Split(cube, " ")[0]
				number, err := strconv.Atoi(strNumber)
				check(err)
				color := strings.Split(cube, " ")[1]
				switch color {
				case "red":
					if number > redMin {
						redMin = number
					}
				case "green":
					if number > greenMin {
						greenMin = number
					}
				case "blue":
					if number > blueMin {
						blueMin = number
					}
				}
			}
		}
		total = total + (redMin * greenMin * blueMin)
	}
	fmt.Println(total)

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
