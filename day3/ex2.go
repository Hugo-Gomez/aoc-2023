package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// number seekers
func getRightNumber(arrayLine []string, index int, value string, includeSpe bool) string {
	if index > len(arrayLine)-1 {
		return value
	}
	if _, err := strconv.Atoi(arrayLine[index]); err == nil {
		value = value + arrayLine[index]
		return getRightNumber(arrayLine, index+1, value, includeSpe)
	} else {
		return value
	}
}

func getLeftNumber(arrayLine []string, index int, value string, includeSpe bool) string {
	if index < 0 {
		return value
	}
	if _, err := strconv.Atoi(arrayLine[index]); err == nil {
		value = arrayLine[index] + value
		return getLeftNumber(arrayLine, index-1, value, includeSpe)
	} else {
		return value
	}
}

func main() {
	total := 0

	file, err := os.Open("./input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	currentLineNumber := 0
	var previousLine []string
	var previousLineSymbols []int
	gears := make(map[int]map[int][]int) // {currentLineNumber: {indexLine : [987, 894], ...}, ... }

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		arrline := strings.Split(line, "")

		// shit happens here
		if currentLineNumber > 0 {
			gears[currentLineNumber] = make(map[int][]int)
			var currentLineSymbols []int
			for i, char := range arrline {
				if char == "." {
					continue
				}
				if char == "*" {
					currentLineSymbols = append(currentLineSymbols, i)

					// right
					rightVal := getRightNumber(arrline, i+1, "", false)
					if rVal, err := strconv.Atoi(rightVal); err == nil {
						gears[currentLineNumber][i] = append(gears[currentLineNumber][i], rVal)
					}
					// left
					leftVal := getLeftNumber(arrline, i-1, "", false)
					if lVal, err := strconv.Atoi(leftVal); err == nil {
						gears[currentLineNumber][i] = append(gears[currentLineNumber][i], lVal)
					}

					// top
					if _, err := strconv.Atoi(previousLine[i]); err == nil {
						topLeftChars := getLeftNumber(previousLine, i-1, "", true)
						topRightChars := getRightNumber(previousLine, i+1, "", true)
						topVal := topLeftChars + previousLine[i] + topRightChars
						if tVal, err := strconv.Atoi(topVal); err == nil {
							gears[currentLineNumber][i] = append(gears[currentLineNumber][i], tVal)
						}
					} else {
						// top left
						if _, err := strconv.Atoi(previousLine[i-1]); err == nil {
							topLeftNumbers := getLeftNumber(previousLine, i-2, "", true)
							topLeftVal := topLeftNumbers + previousLine[i-1]
							if tlVal, err := strconv.Atoi(topLeftVal); err == nil {
								gears[currentLineNumber][i] = append(gears[currentLineNumber][i], tlVal)
							}
						}
						// top right
						if _, err := strconv.Atoi(previousLine[i+1]); err == nil {
							topRightNumbers := getRightNumber(previousLine, i+2, "", true)
							topRightVal := previousLine[i+1] + topRightNumbers
							if trVal, err := strconv.Atoi(topRightVal); err == nil {
								gears[currentLineNumber][i] = append(gears[currentLineNumber][i], trVal)
							}
						}
					}
				}
			}
			for _, i := range previousLineSymbols {
				// bottom
				if _, err := strconv.Atoi(arrline[i]); err == nil {
					bottomLeftChars := getLeftNumber(arrline, i-1, "", true)
					bottomRightChars := getRightNumber(arrline, i+1, "", true)
					bottomVal := bottomLeftChars + arrline[i] + bottomRightChars
					if bVal, err := strconv.Atoi(bottomVal); err == nil {
						gears[currentLineNumber-1][i] = append(gears[currentLineNumber-1][i], bVal)
					}
				} else {
					// bottom left
					if _, err := strconv.Atoi(arrline[i-1]); err == nil {
						bottomLeftNumbers := getLeftNumber(arrline, i-2, "", true)
						bottomLeftVal := bottomLeftNumbers + arrline[i-1]
						if blVal, err := strconv.Atoi(bottomLeftVal); err == nil {
							gears[currentLineNumber-1][i] = append(gears[currentLineNumber-1][i], blVal)
						}
					}
					// bottom right
					if _, err := strconv.Atoi(arrline[i+1]); err == nil {
						bottomRightNumbers := getRightNumber(arrline, i+2, "", true)
						bottomRightVal := arrline[i+1] + bottomRightNumbers
						if brVal, err := strconv.Atoi(bottomRightVal); err == nil {
							gears[currentLineNumber-1][i] = append(gears[currentLineNumber-1][i], brVal)
						}
					}
				}
			}
			previousLineSymbols = currentLineSymbols
		}

		// setup next iteration
		previousLine = arrline
		currentLineNumber = currentLineNumber + 1
	}
	for _, lineGears := range gears {
		for _, numbers := range lineGears {
			if len(numbers) == 2 {
				total = total + (numbers[0] * numbers[1])
			}
		}
	}
	fmt.Println(total)

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
