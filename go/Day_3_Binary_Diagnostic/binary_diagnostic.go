package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func filterArray(arr []string, position int, value string) []string {
	index := 0
	for {
		if index >= len(arr) {
			break
		}
		currentelement := arr[index]
		if string(currentelement[position]) != value {
			arr = append(arr[:index], arr[index+1:]...)
			continue
		}
		index = index + 1
	}
	return arr
}

func getFrequentValue(arr []string, position int, tie string) string {
	trues := 0
	falses := 0
	for _, v := range arr {
		if string(v[position]) == "1" {
			trues = trues + 1
		} else {
			falses = falses + 1
		}
	}

	if trues == falses {
		return tie
	}

	if trues > falses {
		return "1"
	}

	return "0"
}

func convertBinaryToDecimal(binary string) int64 {
	decimal, err := strconv.ParseInt(binary, 2, 64)
	if err != nil {
		log.Fatal(err)
	}

	return decimal
}

func partOne() {
	var gammaRate int64 = 0
	var epsilonRate int64 = 0
	var count int = 0
	var countEachNumber int = 0
	countBit := map[int]int{}

	file, err := os.Open("puzzle_input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		value := scanner.Text()
		countEachNumber = len(value)

		for pos, char := range value {
			if string(char) == "0" {
				countBit[pos]++
			}
		}

		count++
	}

	sampleArray := make([]string, countEachNumber)
	for index := range sampleArray {
		sampleArray[index] = "0"
	}

	binaryGammaRate := string(strings.Join(sampleArray, ""))
	binaryEpsilonRate := string(strings.Join(sampleArray, ""))

	for key, element := range countBit {
		if count-element > element {
			binaryGammaRate = binaryGammaRate[:key] + "1" + binaryGammaRate[key+1:]
		} else {
			binaryEpsilonRate = binaryEpsilonRate[:key] + "1" + binaryEpsilonRate[key+1:]
		}
	}

	gammaRate = convertBinaryToDecimal(binaryGammaRate)
	epsilonRate = convertBinaryToDecimal(binaryEpsilonRate)
	fmt.Println("Part One: ", gammaRate*epsilonRate)
}

func partTwo() {
	var oxygenGeneratorRating int64 = 0
	var co2ScrubberRating int64 = 0
	var index int = 0

	inputAsBytes, err := os.ReadFile("puzzle_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	inputAsStr := string(inputAsBytes)
	lines := strings.Split(inputAsStr, "\n")

	cpLines := make([]string, len(lines))
	copy(cpLines, lines)

	for {
		if len(cpLines) == 1 {
			oxygenGeneratorRating = convertBinaryToDecimal(cpLines[0])
			break
		}
		mostFrequent := getFrequentValue(cpLines, index, "1")
		cpLines = filterArray(cpLines, index, mostFrequent)
		index = index + 1
	}

	cpLines = make([]string, len(lines))
	copy(cpLines, lines)
	index = 0

	for {
		if len(cpLines) == 1 {
			co2ScrubberRating = convertBinaryToDecimal(cpLines[0])
			break
		}
		mostFrequent := getFrequentValue(cpLines, index, "tie")
		if mostFrequent == "1" || mostFrequent == "tie" {
			cpLines = filterArray(cpLines, index, "0")
		} else {
			cpLines = filterArray(cpLines, index, "1")
		}
		index = index + 1
	}

	fmt.Println("Part Two: ", co2ScrubberRating*oxygenGeneratorRating)

}

func main() {
	partOne()
	partTwo()
}
