package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func isSpecialChar(str string) bool {
	return !regexp.MustCompile(`^[0-9.]+$`).MatchString(str)
}

// number seekers
func getRightNumber(arrayLine []string, index int, value string, includeSpe bool) string {
	if index > len(arrayLine)-1 {
		return value
	}
	if _, err := strconv.Atoi(arrayLine[index]); err == nil {
		value = value + arrayLine[index]
		return getRightNumber(arrayLine, index+1, value, includeSpe)
	} else if includeSpe && isSpecialChar(arrayLine[index]) {
		value = value + arrayLine[index]
		return value
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
	} else if includeSpe && isSpecialChar(arrayLine[index]) {
		value = arrayLine[index] + value
		return value
	} else {
		return value
	}
}

// avoid-repeting-code methods
func addToTotalIfValidInt(val string, total int) int {
	if val != "" {
		if intVal, err := strconv.Atoi(val); err == nil && !isSpecialChar(val) {
			return total + intVal
		} else {
			return total
		}
	}
	return total
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

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		arrline := strings.Split(line, "")

		// shit happens here
		if currentLineNumber > 0 {
			var currentLineSymbols []int
			for i, char := range arrline {
				if char == "." {
					continue
				}
				if _, err := strconv.Atoi(char); err != nil {
					currentLineSymbols = append(currentLineSymbols, i)
					// right
					rightVal := getRightNumber(arrline, i+1, "", false)
					total = addToTotalIfValidInt(rightVal, total)
					// left
					leftVal := getLeftNumber(arrline, i-1, "", false)
					total = addToTotalIfValidInt(leftVal, total)
					// top
					if _, err := strconv.Atoi(previousLine[i]); err == nil {
						topLeftChars := getLeftNumber(previousLine, i-1, "", true)
						topRightChars := getRightNumber(previousLine, i+1, "", true)
						topVal := topLeftChars + previousLine[i] + topRightChars
						total = addToTotalIfValidInt(topVal, total)
					} else {
						// top left
						if _, err := strconv.Atoi(previousLine[i-1]); err == nil {
							topLeftNumbers := getLeftNumber(previousLine, i-2, "", true)
							topLeftVal := topLeftNumbers + previousLine[i-1]
							total = addToTotalIfValidInt(topLeftVal, total)
						}
						// top right
						if _, err := strconv.Atoi(previousLine[i+1]); err == nil {
							topRightNumbers := getRightNumber(previousLine, i+2, "", true)
							topRightVal := previousLine[i+1] + topRightNumbers
							total = addToTotalIfValidInt(topRightVal, total)
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
					total = addToTotalIfValidInt(bottomVal, total)
				} else {
					// bottom left
					if _, err := strconv.Atoi(arrline[i-1]); err == nil {
						bottomLeftNumbers := getLeftNumber(arrline, i-2, "", true)
						bottomLeftVal := bottomLeftNumbers + arrline[i-1]
						total = addToTotalIfValidInt(bottomLeftVal, total)
					}
					// bottom right
					if _, err := strconv.Atoi(arrline[i+1]); err == nil {
						bottomRightNumbers := getRightNumber(arrline, i+2, "", true)
						bottomRightVal := arrline[i+1] + bottomRightNumbers
						total = addToTotalIfValidInt(bottomRightVal, total)
					}
				}
			}
			previousLineSymbols = currentLineSymbols
		}

		// setup next iteration
		previousLine = arrline
		currentLineNumber = currentLineNumber + 1
	}
	fmt.Println(total)

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
