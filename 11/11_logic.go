package quest11

import (
	"strconv"
	"strings"

	"github.com/Evokoo/EBC_2025_Go/utils"
)

// ========================
// PARSER
// ========================

func ParseInput(file string) []int {
	data := utils.ReadFile(file)

	var output []int
	for _, value := range strings.Split(data, "\n") {
		n, _ := strconv.Atoi(value)
		output = append(output, n)
	}

	return output
}

// ========================
// PART I
// ========================

func I(columns []int, rounds int) int {
	//Phase one
	for !IsAcending(columns) {
		for i := range len(columns) - 1 {
			if columns[i+1] < columns[i] {
				columns[i]--
				columns[i+1]++
			}
		}
		rounds--
	}

	//Phase two
	for i := 0; i < rounds; i++ {
		for i := range len(columns) - 1 {
			if columns[i+1] > columns[i] {
				columns[i]++
				columns[i+1]--
			}
		}
	}

	checksum := 0
	for i, value := range columns {
		checksum += (i + 1) * value
	}

	return checksum
}

func IsAcending(arr []int) bool {
	for i := range len(arr) - 1 {
		a := arr[i]
		b := arr[i+1]

		if a > b {
			return false
		}
	}
	return true
}

// ========================
// PART II
// ========================

func II(columns []int) (rounds int) {

	for !IsAcending(columns) {
		for i := range len(columns) - 1 {
			if columns[i+1] < columns[i] {
				columns[i]--
				columns[i+1]++
			}
		}
		rounds++
	}

	for !IsBalanced(columns) {
		for i := range len(columns) - 1 {
			if columns[i+1] > columns[i] {
				columns[i]++
				columns[i+1]--
			}
		}
		rounds++
	}

	return rounds
}

func IsBalanced(arr []int) bool {
	target := arr[0]
	for _, value := range arr[1:] {
		if value != target {
			return false
		}
	}
	return true
}

// ========================
// PART III
// ========================

func III(columns []int) int {
	avg := Sum(columns) / len(columns)
	rounds := 0

	for _, n := range columns {
		if n < avg {
			rounds += avg - n
		}
	}

	return rounds
}

func Sum(numbers []int) (sum int) {
	for _, n := range numbers {
		sum += n
	}
	return
}
