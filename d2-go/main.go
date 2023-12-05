package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	MaxRedCubes   = 12
	MaxGreenCubes = 13
	MaxBlueCubes  = 14
)

var MaxCubesMap = map[string]int{
	"red":   MaxRedCubes,
	"green": MaxGreenCubes,
	"blue":  MaxBlueCubes,
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
	sum := 0

LINELOOP:
	for scanner.Scan() {
		line := scanner.Text()

		maxGameCubeMap := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}

		splitLine := strings.Split(line, ":")
		// gameIdStr := splitLine[0]

		splitLine = strings.Split(splitLine[1], ";")
		for _, roundStr := range splitLine {
			roundRegex := regexp.MustCompile(`(\d+) (red|green|blue)`)
			roundPicks := roundRegex.FindAllString(roundStr, -1)

			for _, pick := range roundPicks {
				pickParts := strings.Split(pick, " ")
				count, err := strconv.Atoi(pickParts[0])
				if err != nil {
					fmt.Println("Error parsing count:", err)
					continue LINELOOP
				}
				color := pickParts[1]

				if count > maxGameCubeMap[color] {
					maxGameCubeMap[color] = count
				}

				// part 1 check
				// if count > MaxCubesMap[color] {
				// 	continue LINELOOP
				// }
			}

			//
		}

		power := 1
		for _, maxColorCount := range maxGameCubeMap {
			power *= maxColorCount
		}

		sum += power

		// parse game id for part 1
		// gameIdParts := strings.Split(gameIdStr, " ")
		// gameId, err := strconv.Atoi(gameIdParts[1])
		// if err != nil {
		// 	fmt.Println("Error parsing game id:", err)
		// 	continue
		// }

		// sum += gameId
	}

	fmt.Println("final sum: ", sum)
}
