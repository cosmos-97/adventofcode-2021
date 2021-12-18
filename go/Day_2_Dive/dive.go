package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func partOne() {
	var depth int = 0
	var horizontal int = 0

	file, err := os.Open("puzzle_input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		value := strings.Split(scanner.Text(), " ")
		intVar, err := strconv.Atoi(value[1])
		if err != nil {
			log.Fatal(err)
		}

		if value[0] == "forward" {
			horizontal = horizontal + intVar
		} else if value[0] == "down" {
			depth = depth + intVar
		} else if value[0] == "up" {
			depth = depth - intVar
		}
	}
	fmt.Println("Part One: ", depth*horizontal)

}

func partTwo() {
	var aim int = 0
	var depth int = 0
	var horizontal int = 0

	file, err := os.Open("puzzle_input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		value := strings.Split(scanner.Text(), " ")
		intVar, err := strconv.Atoi(value[1])
		if err != nil {
			log.Fatal(err)
		}

		if value[0] == "forward" {
			horizontal = horizontal + intVar
			depth = depth + (aim * intVar)
		} else if value[0] == "down" {
			aim = aim + intVar
		} else if value[0] == "up" {
			aim = aim - intVar
		}
	}
	fmt.Println("Part Two: ", depth*horizontal)
}

func main() {
	partOne()
	partTwo()
}
