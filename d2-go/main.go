package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file
	scanner := bufio.NewScanner(file)

	wordNumberMapping := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	lineNumber := 1
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		regex := regexp.MustCompile(`(one)|(two)|(three)|(four)|(five)|(six)|(seven)|(eight)|(nine)|\d`)
		matches := regex.FindAllString(line, -1)

		matchMap := make(map[string]string)
		for _, match := range matches {
			matchMap[match] = ""
		}

		// for every match, replace match in line with last char in the match
		newLine := line
		for match := range matchMap {
			lastChar := match[len(match)-1:]
			if lastChar == match {
				lastChar = ""
			}
			newLine = strings.Replace(newLine, match, match+"_"+lastChar, -1)
			// fmt.Println("line number ", lineNumber, " : ", newLine)
		}
		// fmt.Println("line number ", lineNumber, " : ", newLine)

		matches = regex.FindAllString(newLine, -1)

		firstNum := matches[0]
		lastNum := matches[len(matches)-1]

		if firstNumMapping, ok := wordNumberMapping[firstNum]; ok {
			firstNum = firstNumMapping
		}

		if lastNumMapping, ok := wordNumberMapping[lastNum]; ok {
			lastNum = lastNumMapping
		}

		num, err := strconv.Atoi(firstNum + lastNum)
		if err == nil {
			fmt.Println("line number ", lineNumber, " : ", matches, " = ", num)
			sum += num
		}

		lineNumber++
	}

	fmt.Println("final sum: ", sum)
}
