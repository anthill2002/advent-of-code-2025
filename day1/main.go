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
	passwordHit := 0
	currentPos := startPosition

	for _, input := range inputs {
		direction := input[0:1]
		distance, _ := strconv.Atoi(input[1:])

		if distance >= 100 {
			newDistance, extraHits := processLargeDistance(distance)
			distance = newDistance
			passwordHit += extraHits
		}

		newPos, dialHits := clicker(distance, currentPos, direction)
		currentPos = newPos
		passwordHit += dialHits
	}

	return passwordHit
}

func processLargeDistance(distance int) (int, int) {
	passwordHit := 0
	temp := distance / 100
	passwordHit += temp
	distance = distance - (temp * 100)
	return distance, passwordHit
}

func clicker(distance int, startPosition int, direction string) (int, int) {
	currentPos := startPosition
	minNumber := 0
	maxNumber := 99
	passwordHit := 0

	for range max(distance, distance*-1) {
		if direction == "R" {
			currentPos++
		} else {
			currentPos--
		}

		if currentPos == maxNumber+1 {
			currentPos = 0
			passwordHit++
		} else if currentPos < minNumber {
			currentPos = 99
		} else if currentPos == 0 {
			passwordHit++
		}
	}
	return currentPos, passwordHit
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
