package quest03

import (
	"sort"
	"strconv"
	"strings"

	"github.com/Evokoo/EBC_2025_Go/utils"
)

// ========================
// PART I
// ========================

func I(boxes []int) int {
	total := 0
	seen := make(map[int]struct{})

	for _, box := range boxes {
		if _, found := seen[box]; found {
			continue
		}

		total += box
		seen[box] = struct{}{}
	}

	return total
}

// ========================
// PART II
// ========================
func II(boxes Boxes) int {
	boxset := make(Boxes, 0, 20)
	seen := make(map[int]struct{})

	for i := len(boxes) - 1; i >= 0; i-- {
		box := boxes[i]
		if _, found := seen[box]; found {
			continue
		}

		boxset = append(boxset, box)
		seen[box] = struct{}{}

		if len(boxset) == 20 {
			break
		}
	}

	return I(boxset)
}

// ========================
// PART III
// ========================
func III(boxes Boxes) int {
	count := make(map[int]int)
	for _, box := range boxes {
		count[box]++
	}

	max := 0
	for _, c := range count {
		if c > max {
			max = c
		}
	}

	return max
}

// ========================
// PARSER
// ========================

type Boxes []int

func ParseInput(file string) Boxes {
	data := utils.ReadFile(file)
	values := []int{}

	for _, value := range strings.Split(data, ",") {
		n, _ := strconv.Atoi(value)
		values = append(values, n)
	}

	sort.Slice(values, func(i, j int) bool {
		return values[i] > values[j]
	})

	return values
}
