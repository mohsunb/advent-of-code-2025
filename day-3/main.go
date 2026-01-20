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

func partTwo() {
	banks := getInput("input.txt")

	var totalJoltage int
	for _, bank := range banks {
		joltages := make([]int, 0, 12)
		lastIndex := -1
		for i := 0; i < 12; i++ {
			biggestJoltage, index := biggestJoltageBetweenIndices(bank, lastIndex+1, len(bank)-1-(11-i))
			lastIndex = index
			joltages = append(joltages, biggestJoltage)
		}
		totalJoltage += calculateTotalJoltage(joltages)

	}

	fmt.Println("Part 2 result:", totalJoltage)
}

func biggestJoltageBetweenIndices(bank string, minIndex int, maxIndex int) (int, int) {
	var result int
	var index int

	for i := minIndex; i <= maxIndex; i++ {
		joltage, err := strconv.Atoi(string(bank[i]))
		if err != nil {
			log.Fatal("Failed to parse joltage:", err)
		}
		if joltage > result {
			result = joltage
			index = i
		}
	}

	return result, index
}

func calculateTotalJoltage(joltages []int) int {
	totalJoltageDigits := make([]rune, 0, len(joltages))
	for i := range joltages {
		totalJoltageDigits = append(totalJoltageDigits, rune(strconv.Itoa(joltages[i])[0]))
	}
	totalJoltage, _ := strconv.Atoi(string(totalJoltageDigits))
	return totalJoltage
}

func main() {
	partOne()
	partTwo()
}
