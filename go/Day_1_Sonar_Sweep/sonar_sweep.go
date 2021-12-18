package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func removeIndex(array []int, index int) []int {
	return append(array[:index], array[index+1:]...)
}

func sumArray(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func partOne() {
	var counter int = 0
	var prevVar int = 0
	var indexOne bool = true

	file, err := os.Open("puzzle_input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		intVar, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		if indexOne {
			indexOne = false
			fmt.Println(intVar, "(N/A - no previous measurement)")
		} else if intVar > prevVar {
			fmt.Println(intVar, "(increased)")
			counter++
		} else {
			fmt.Println(intVar, "(decreased)")
		}
		prevVar = intVar
	}

	fmt.Println("Answer Part One: ", counter)
}

func partTwo() {
	arrayA := []int{}
	var counter int = 0
	var sum int = 0
	var prevSum int = 0
	var indexOne bool = true

	file, err := os.Open("puzzle_input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		intVar, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		arrayA = append(arrayA, intVar)

		if len(arrayA) == 3 {
			sum = sumArray(arrayA)

			if indexOne {
				indexOne = false
				fmt.Println(sum, "(N/A - no previous measurement)")
			} else if sum > prevSum {
				fmt.Println(sum, "(increased)")
				counter++
			} else {
				fmt.Println(sum, "(decreased)")
			}

			arrayA = removeIndex(arrayA, 0)
			prevSum = sum
		}
	}

	fmt.Println("Answer Part Two: ", counter)
}

func main() {
	partOne()
	partTwo()
}
