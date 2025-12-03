package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	startPosition := 50

	inputLines, err := readInput("../inputs/day1-input.txt")
	if err != nil {
		fmt.Println("Error reading file", err)
	}
	password := processInput(startPosition, inputLines)
	// password := processInput(startPosition, []string{
	// 	"L68",
	// 	"L30",
	// 	"R48",
	// 	"L5",
	// 	"R60",
	// 	"L55",
	// 	"L1",
	// 	"L99",
	// 	"R14",
	// 	"L82",
	// })
	fmt.Printf("Password is %v", password)
}

func processInput(startPosition int, inputs []string) int {
	minNumber := 0
	maxNumber := 99
	var currentPos int = startPosition
	passwordHit := 0

	for _, input := range inputs {
		var direction = input[0:1]
		var distance, _ = strconv.Atoi(input[1:])
		if direction == "R" {
			currentPos += distance
		} else {
			currentPos -= distance
		}

		for currentPos < minNumber || currentPos > maxNumber {
			if currentPos > maxNumber {
				currentPos = currentPos - maxNumber - 1
			} else if currentPos < minNumber {
				currentPos = maxNumber + currentPos + 1
			}
		}

		if currentPos == 0 {
			passwordHit++
		}
	}

	return passwordHit
}

func readInput(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}
