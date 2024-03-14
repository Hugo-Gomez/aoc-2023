package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)


func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	var total int = 0

	file, err := os.Open("./input")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		arrline := strings.Split(line, "")
		var line_numbers []string
		for _, ch := range arrline  {
			if _, err := strconv.Atoi(ch); err == nil {
				line_numbers = append(line_numbers, ch)
			}
		}
		string_number := line_numbers[0] + line_numbers[len(line_numbers) - 1]
		line_number, err := strconv.Atoi(string_number)
		check(err)
		total = total + line_number
	}
	fmt.Println(total)

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
