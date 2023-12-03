package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
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

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		regex := regexp.MustCompile(`\d`)
		matches := regex.FindAllString(line, -1)

		firstNumChar := matches[0]
		lastNumChar := matches[len(matches)-1]

		num, err := strconv.Atoi(firstNumChar + lastNumChar)

		if err == nil {
			sum += num
		}
	}

	fmt.Println("final sum: ", sum)
}
