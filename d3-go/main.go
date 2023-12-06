package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var gearCoordinateMap map[[2]int][]int = make(map[[2]int][]int)

// assumes str is length 1
func isSymbol(str string) bool {
	char := str[0]
	return !((char == 46) || (char >= 48 && char <= 57) || (char >= 65 && char <= 90) || (char >= 97 && char <= 122))
}

func isGearSymbol(str string) bool {
	return str == "*"
}

func addNumberToGearCoordinates(gearCoordinate [2]int, number int) {
	if _, ok := gearCoordinateMap[gearCoordinate]; !ok {
		gearCoordinateMap[gearCoordinate] = make([]int, 0)
	}
	gearCoordinateMap[gearCoordinate] = append(gearCoordinateMap[gearCoordinate], number)
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

	// sum := 0
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

			// use for part 1
			// if start > 0 && isSymbol(string(line[start-1])) {
			// 	// todo for p2 change to add an entry to the map at coordinates - if no entry, make int literal with new num
			// 	sum += partNumber
			// 	continue
			// } else if end < len(line) && isSymbol(string(line[end])) {
			// 	sum += partNumber
			// 	continue
			// }

			if start > 0 && isGearSymbol(string(line[start-1])) {
				addNumberToGearCoordinates([2]int{start - 1, lineIndex}, partNumber)
				continue
			} else if end < len(line) && isGearSymbol(string(line[end])) {
				addNumberToGearCoordinates([2]int{end, lineIndex}, partNumber)
				continue
			}

			symbolRegexp := regexp.MustCompile(`[*]`)
			// use for part 1
			// symbolRegexp := regexp.MustCompile(`[^0-9A-Za-z.]`)
			if lineIndex > 0 {
				symbolMatches := symbolRegexp.FindAllStringIndex(lines[lineIndex-1], -1)
				for _, symbolMatch := range symbolMatches {
					symbolStart := symbolMatch[0] // inclusive
					if symbolStart >= start-1 && symbolStart <= end {
						addNumberToGearCoordinates([2]int{symbolStart, lineIndex - 1}, partNumber)
						// use for part 1
						// sum += partNumber
						continue PARTNUMBERLOOP
					}
				}
			}
			if lineIndex < len(lines)-1 { // same function but for line below
				symbolMatches := symbolRegexp.FindAllStringIndex(lines[lineIndex+1], -1)
				for _, symbolMatch := range symbolMatches {
					symbolStart := symbolMatch[0] // inclusive
					if symbolStart >= start-1 && symbolStart <= end {
						addNumberToGearCoordinates([2]int{symbolStart, lineIndex + 1}, partNumber)
						// use for part 1
						// sum += partNumber
						continue PARTNUMBERLOOP
					}
				}
			}
		}
	}

	// fmt.Println("gear map", gearCoordinateMap)

	// part 2
	gearRatioSum := 0
	for _, values := range gearCoordinateMap {
		if len(values) == 2 {
			gearRatio := values[0] * values[1]
			gearRatioSum += gearRatio
		}
	}

	// fmt.Println("final sum: ", sum)
	fmt.Println("final gear ratio sum:", gearRatioSum)
}
