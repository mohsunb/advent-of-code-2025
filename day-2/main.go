package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type ProductRange struct {
	LowerBound int
	UpperBound int
}

func getInput(fileName string) []ProductRange {
	stream, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}

	defer stream.Close()

	scanner := bufio.NewScanner(stream)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	input := scanner.Text()

	rawProductRanges := strings.Split(input, ",")

	productRanges := make([]ProductRange, 0, len(rawProductRanges))
	for _, rawProductRange := range rawProductRanges {
		bounds := strings.Split(rawProductRange, "-")
		lowerBound, err := strconv.Atoi(bounds[0])
		if err != nil {
			log.Fatal(fmt.Sprintf("Failed to derive lower bound from product range %s:", rawProductRange), err)
		}
		upperBound, err := strconv.Atoi(bounds[1])
		if err != nil {
			log.Fatal(fmt.Sprintf("Failed to derive upper bound from product range %s:", rawProductRange), err)
		}
		productRanges = append(productRanges, ProductRange{
			LowerBound: lowerBound,
			UpperBound: upperBound,
		})
	}

	return productRanges
}

func partOne() {
	productRanges := getInput("input.txt")

	var sum int
	for _, productRange := range productRanges {
		for productId := productRange.LowerBound; productId <= productRange.UpperBound; productId++ {
			stringifiedProductId := strconv.Itoa(productId)
			length := len(stringifiedProductId)
			if length%2 == 0 {
				firstHalf := stringifiedProductId[:length/2]
				secondHalf := stringifiedProductId[length/2:]
				if firstHalf == secondHalf {
					sum += productId
				}
			}
		}
	}

	fmt.Println("Part 1 result:", sum)
}

func partTwo() {
	productRanges := getInput("input.txt")

	var sum int
	for _, productRange := range productRanges {
		for productId := productRange.LowerBound; productId <= productRange.UpperBound; productId++ {
			stringifiedProductId := strconv.Itoa(productId)
			if !isValid(stringifiedProductId) {
				sum += productId
			}
		}
	}

	fmt.Println("Part 2 result:", sum)
}

func isValid(productId string) bool {
	valid := true

	length := len(productId)
	maxLengthPerRepetition := length / 2

	for i := 1; i <= maxLengthPerRepetition; i++ {

		if length%i == 0 {
			repetitions := length / i
			if strings.Repeat(productId[:i], repetitions) == productId {
				valid = false
			}
		}

		if !valid {
			break
		}
	}

	return valid
}

func main() {
	partOne()
	partTwo()
}
