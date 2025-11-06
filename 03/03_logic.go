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
	seen := make(Set)

	for _, box := range boxes {
		if seen.Has(box) {
			continue
		}

		total += box
		seen.Add(box)
	}

	return total
}

// ========================
// PART II
// ========================
func II(boxes Boxes) int {
	boxset := make(Boxes, 0, 20)
	seen := make(Set)

	for i := len(boxes) - 1; i >= 0; i-- {
		box := boxes[i]
		if seen.Has(box) {
			continue
		}

		boxset = append(boxset, box)
		seen.Add(box)

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
// SET
// ========================
type Set map[int]struct{}

func (s *Set) Add(n int) {
	(*s)[n] = struct{}{}
}
func (s *Set) Has(n int) bool {
	_, found := (*s)[n]
	return found
}

// ========================
// PARSER
// ========================

type Boxes []int

func ParseInput(file string, part int) Boxes {
	data := utils.ReadFile(file)
	values := []int{}

	for _, value := range strings.Split(data, ",") {
		n, _ := strconv.Atoi(value)
		values = append(values, n)
	}

	if part == 2 {
		sort.Slice(values, func(i, j int) bool {
			return values[i] > values[j]
		})
	}

	return values
}
