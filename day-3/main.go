package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func getInput(fileName string) []string {
	stream, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}

	defer stream.Close()

	scanner := bufio.NewScanner(stream)

	input := make([]string, 0)

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	return input
}

func partOne() {
	banks := getInput("input.txt")

	var totalJoltage int
	for _, bank := range banks {
		count := len(bank)
		var highestJoltage int
		for i := 0; i < count-1; i++ {
			for j := i + 1; j < count; j++ {
				joltageString := string(bank[i]) + string(bank[j])
				joltage, err := strconv.Atoi(joltageString)
				if err != nil {
					log.Fatal(fmt.Sprintf("Failed to derive joltage from combination %s", joltageString), err)
				}
				if joltage > highestJoltage {
					highestJoltage = joltage
				}
			}
		}
		totalJoltage += highestJoltage
	}

	fmt.Println("Part 1 result:", totalJoltage)
}

func main() {
	partOne()
}
