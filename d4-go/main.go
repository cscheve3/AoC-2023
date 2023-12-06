package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func getLineScore(line string) float64 {
	splitLine := strings.Split(line, ":")
	// gameIdStr := splitLine[0]

	splitLine = strings.Split(splitLine[1], "|")
	winningNumbers := strings.Split(strings.Trim(splitLine[0], " "), " ")
	playingNumbers := strings.Split(strings.Trim(splitLine[1], " "), " ")

	winningNumbersMap := make(map[string]bool)
	for _, number := range winningNumbers {
		if number != "" {
			winningNumbersMap[number] = true
		}
	}

	winningCount := 0.0
	// winnerMatches := make([]string, 0)
	for _, number := range playingNumbers {
		if number == "" {
			continue
		}

		if _, ok := winningNumbersMap[number]; ok {
			winningCount++
			// winnerMatches = append(winnerMatches, number)
		}
	}

	return winningCount
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

	sum := 0.0
	for scanner.Scan() {
		line := scanner.Text()

		winningCount := getLineScore(line)

		// fmt.Println("line data:", winnerMatches, winningCount)

		if winningCount > 0 {
			sum += math.Pow(2, float64(winningCount-1))
		}
	}

	fmt.Println("final sum:", sum)
}

func mainp2() {
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file
	scanner := bufio.NewScanner(file)

	scores := make([]float64, 0)
	sum := 0.0
	for scanner.Scan() {
		line := scanner.Text()

		scores = append(scores, getLineScore(line))

	}

	// for lineIndex, line := range lines {

	// }

	fmt.Println("final sum:", sum)
}
