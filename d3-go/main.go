package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello, World!")
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
		// line := scanner.Text()
		// parse everything into a list of strings (lines)

	}

	// for each line
	//  find every number and its end and start index -> put [start end] pair into list [][2]int ? or something like that
	// for each pair
	//  is there a symbol before start or after end -> if yes great stop and add to sum
	//  for the line above it is there a symbol between start-1 and end+1 -> if yes great stop and add to sum
	//  for the line below it is there a symbol between start-1 and end+1 -> if yes great stop and add to sum

	// TODO min with length and max with 0 for ends and starts respectively
	// TODO symbol is not number character and not '.'
	// use regexs for things as can be for efficiency

	fmt.Println("final sum: ", sum)
}
