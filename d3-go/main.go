package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

// assumes str is length 1
func isSymbol(str string) bool {
	char := str[0]
	return !((char == 46) || (char >= 48 && char <= 57) || (char >= 65 && char <= 90) || (char >= 97 && char <= 122))
}

func containsSymbol(str string) bool {
	for _, char := range str {
		if isSymbol(string(char)) {
			return true
		}
	}
	return false
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file
	scanner := bufio.NewScanner(file)

	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	sum := 0
	for lineIndex, line := range lines {
		numberRegexp := regexp.MustCompile(`(\d+)`)
		numberMatches := numberRegexp.FindAllStringIndex(line, -1)
		// [pair is inclusive of start and exclusive of end]

	PARTNUMBERLOOP:
		for _, numberMatch := range numberMatches {
			start := numberMatch[0] // inclusive
			end := numberMatch[1]   // exclusive

			partNumber, err := strconv.Atoi(line[start:end])
			if err != nil {
				fmt.Println("Error parsing part number:", err)
				partNumber = 0
			}

			// todo issue - always true
			if start > 0 && isSymbol(string(line[start-1])) {
				sum += partNumber
				continue
			} else if end < len(line) && isSymbol(string(line[end])) {
				sum += partNumber
				continue
			}

			symbolRegexp := regexp.MustCompile(`[^0-9A-Za-z.]`)
			if lineIndex > 0 {
				symbolMatches := symbolRegexp.FindAllStringIndex(lines[lineIndex-1], -1)
				// [pair is inclusive of start and exclusive of end]

				for _, symbolMatch := range symbolMatches {
					symbolStart := symbolMatch[0] // inclusive

					if symbolStart >= start-1 && symbolStart <= end {
						sum += partNumber
						continue PARTNUMBERLOOP
					}
				}
			}
			if lineIndex < len(lines)-1 { // same function but for line below
				symbolMatches := symbolRegexp.FindAllStringIndex(lines[lineIndex+1], -1)

				for _, symbolMatch := range symbolMatches {
					symbolStart := symbolMatch[0] // inclusive

					if symbolStart >= start-1 && symbolStart <= end {
						sum += partNumber
						continue PARTNUMBERLOOP
					}
				}
			}
		}

		// fmt.Println(string(line[numberMatches[0][0]]))
	}

	fmt.Println("final sum: ", sum)
}
