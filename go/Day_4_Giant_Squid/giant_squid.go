package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type BingoField struct {
	field_values  [][]int
	fields_checks [][]bool
}

func (bf *BingoField) ckeckFieldValue(value int) {
	for column := 0; column < 5; column++ {
		for row := 0; row < 5; row++ {

			fieldvalue := bf.field_values[row][column]
			if fieldvalue == value {
				bf.fields_checks[row][column] = true
			}
		}
	}
}

func (bf *BingoField) sumUncheckedFields() int {
	sum := 0
	for column := 0; column < 5; column++ {

		for row := 0; row < 5; row++ {
			fieldvalue := bf.fields_checks[row][column]
			if !fieldvalue {
				sum += bf.field_values[row][column]
			}
		}
	}
	return sum
}

func (bf *BingoField) printBoard() {
	for _, row := range bf.field_values {
		fmt.Println(row)
	}
	for _, row := range bf.fields_checks {
		fmt.Println(row)
	}
}

func (bf *BingoField) hasBingo() bool {
	for _, row := range bf.fields_checks {
		isBingo := true

		for _, fieldvalue := range row {
			if !fieldvalue {
				isBingo = false
			}
		}
		if isBingo {
			return true
		}
	}

	for column := 0; column < 5; column++ {
		isBingo := true

		for row := 0; row < 5; row++ {
			fieldvalue := bf.fields_checks[row][column]
			if !fieldvalue {
				isBingo = false
			}
		}
		if isBingo {
			return true
		}
	}
	return false
}

func lineToBoardLine(line string) []int {
	arr := make([]int, 5)

	for i := 0; i < 5; i = i + 1 {
		arr[i], _ = strconv.Atoi(strings.Trim(line[(i*3):(i*3+2)], " "))
	}
	return arr
}

func main() {
	var bingoFields []BingoField
	var arrayNumbers []string

	dataByte, err := os.ReadFile("puzzle_input.txt")
	if err != nil {
		log.Fatal(err)
	}

	dataStr := string(dataByte)
	lines := strings.Split(dataStr, "\n")

	arrayNumbers = strings.Split(lines[0], ",")
	lines = lines[2:]

	for {

		bingoFieldData := make([][]int, 5)
		bingoFieldBoolData := make([][]bool, 5)

		for i := 0; i < 5; i++ {
			bingoFieldData[i] = lineToBoardLine(lines[i])
			bingoFieldBoolData[i] = []bool{false, false, false, false, false}
		}

		bingofield := BingoField{field_values: bingoFieldData, fields_checks: bingoFieldBoolData}
		bingoFields = append(bingoFields, bingofield)
		if len(lines) >= 6 {
			lines = lines[6:]
		} else {
			break
		}
	}

	for _, numberStr := range arrayNumbers {
		endLoop := false

		numberInt, err := strconv.Atoi(numberStr)
		if err != nil {
			log.Fatal(err)
		}

		for _, bingo := range bingoFields {
			bingo.ckeckFieldValue(numberInt)
			if bingo.hasBingo() {
				fmt.Println("Part One: ", bingo.sumUncheckedFields()*numberInt)
				endLoop = true
			}
		}

		if endLoop {
			break
		}
	}
}
