package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Rotation struct {
	Direction rune
	Distance  int
}

func readInput() ([]string, error) {
	stream, err := os.Open("input.txt")

	if err != nil {
		return nil, err
	}

	defer stream.Close()

	scanner := bufio.NewScanner(stream)
	scanner.Split(bufio.ScanLines)

	result := make([]string, 0)

	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	return result, nil
}

func parseInput(input []string) ([]Rotation, error) {
	result := make([]Rotation, 0, len(input))

	for _, line := range input {
		direction := rune(line[0])
		distance, err := strconv.Atoi(line[1:])
		if err != nil {
			return nil, err
		}
		rotation := Rotation{
			Direction: direction,
			Distance:  distance,
		}
		result = append(result, rotation)
	}

	return result, nil
}

func firstPart() {
	rawInput, err := readInput()
	if err != nil {
		log.Fatal("Failed to read the input: ", err)
	}
	input, err := parseInput(rawInput)
	if err != nil {
		log.Fatal("Failed to parse the input: ", err)
	}

	position := 50
	zeroCounter := 0
	for _, rotation := range input {
		simplifiedDistance := rotation.Distance % 100

		switch rotation.Direction {
		case 'L':
			position -= simplifiedDistance
		case 'R':
			position += simplifiedDistance
		}

		if position < 0 {
			position += 100
		} else if position > 99 {
			position -= 100
		}

		if position == 0 {
			zeroCounter++
		}
	}

	fmt.Println("Answer: " + strconv.Itoa(zeroCounter))
}

func main() {
	firstPart()
}
