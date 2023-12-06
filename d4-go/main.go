package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func getLineScore(line string) int {
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

	winningCount := 0
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

func mainp1() {
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

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file
	scanner := bufio.NewScanner(file)

	scores := make([]int, 0)
	cardInstances := make([]int, 0)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()

		scores = append(scores, getLineScore(line))
		cardInstances = append(cardInstances, 1)
	}

	for card, score := range scores {
		instances := cardInstances[card]
		for i := 1; i <= score; i++ {
			cardInstances[card+i] += instances
		}

		sum += instances
	}

	// fmt.Println("final instances:", cardInstances)

	fmt.Println("final sum:", sum)
}
