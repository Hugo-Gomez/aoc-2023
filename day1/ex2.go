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
	m := map[string]string{
		"one": "1",
		"two": "2",
		"three": "3",
		"four": "4",
		"five": "5",
		"six": "6",
		"seven": "7",
		"eight": "8",
		"nine": "9",
	}

	file, err := os.Open("./input")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		arrline := strings.Split(line, "")
		digits := []string{}
		for k, v := range m {
			digits = append(digits, k, v)
		}
		var f_digit, l_digit string
		var f_tmp, l_tmp string
		// f
		out:
		for _, char := range arrline  {
			f_tmp = f_tmp + char
			for _, digit := range digits {
				if strings.Contains(f_tmp, digit) {
					if len(digit) == 1 {
						f_digit = digit
					} else {
						f_digit = m[digit]
					}
					break out
				}
			}
		}
		// l
		lout:
		for i := len(arrline)-1; i >= 0; i-- {
			l_tmp = arrline[i] + l_tmp
			for _, digit := range digits {
				if strings.Contains(l_tmp, digit) {
					if len(digit) == 1 {
						l_digit = digit
					} else {
						l_digit = m[digit]
					}
					break lout
				}
			}
		}

		string_number := f_digit + l_digit
		line_number, err := strconv.Atoi(string_number)
		check(err)
		total = total + line_number
	}
	fmt.Println(total)

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
