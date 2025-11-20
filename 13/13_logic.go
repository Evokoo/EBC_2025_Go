package quest13

import (
	"strconv"
	"strings"

	"github.com/Evokoo/EBC_2025_Go/utils"
)

// ========================
// PART I
// ========================
func I(file string, rotations int) int {
	data := utils.ReadFile(file)
	lines := strings.Split(data, "\n")
	length := len(lines) + 1

	dial := make([]int, length)
	dial[0] = 1

	offset := 1
	for i, value := range lines {
		n, _ := strconv.Atoi(value)

		if i%2 == 0 {
			dial[offset] = n
		} else {
			dial[length-offset] = n
			offset++
		}
	}

	return dial[rotations%length]
}

// ========================
// PART II
// ========================
type Section struct {
	start, end, length, step int
}

func II(file string, rotations int) int {
	data := utils.ReadFile(file)
	lines := strings.Split(data, "\n")
	sections := len(lines) + 1

	dial := make([]Section, sections)
	dial[0] = Section{1, 1, 1, 1}

	offset := 1
	totalLength := 1

	for i, values := range lines {
		digits := strings.Split(values, "-")
		a, _ := strconv.Atoi(digits[0])
		b, _ := strconv.Atoi(digits[1])
		c := b - a + 1

		totalLength += c

		if i%2 == 0 {
			dial[offset] = Section{a, b, c, 1}
		} else {
			dial[sections-offset] = Section{b, a, c, -1}
			offset++
		}
	}

	target := rotations % totalLength
	current := 0

	for _, section := range dial {
		if target >= current && target < current+section.length {
			for i := section.start; i != section.end; i += section.step {

				if current == target {
					return i
				}
				current++
			}
		} else {
			current += section.length
		}
	}

	return 0
}
