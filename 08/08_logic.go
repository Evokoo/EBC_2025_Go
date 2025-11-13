package quest08

import (
	"strconv"
	"strings"

	"github.com/Evokoo/EBC_2025_Go/utils"
)

// ========================
// PART I
// ========================
func I(list []int, pairs Pairs) int {
	count := 0

	for i := range len(list) - 1 {
		if pairs[list[i]] == list[i+1] {
			count++
		}
	}

	return count
}

// ========================
// PART II
// ========================
func II(list []int) int {
	seen := make(map[[2]int]struct{})
	knots := 0

	for i := range len(list) - 1 {
		a := list[i]
		b := list[i+1]

		for key := range seen {
			if Intersects(a, b, key[0], key[1]) {
				knots++
			}
		}
		seen[[2]int{a, b}] = struct{}{}
	}
	return knots
}

// ========================
// PART III
// ========================
func III(list []int, nails int) int {
	pattern := make(map[[2]int]int)
	for i := range len(list) - 1 {
		key := [2]int{list[i], list[i+1]}

		if key[0] > key[1] {
			key[0], key[1] = key[1], key[0]
		}

		pattern[key]++
	}

	best := 0

	for a := 1; a <= nails/2; a++ {
		for b := a + nails/2; b <= nails; b++ {
			count := 0
			for pair, layers := range pattern {
				if Intersects(a, b, pair[0], pair[1]) {
					count += layers
				}
			}
			best = max(best, count)
		}
	}

	return best
}

func Intersects(a, b, c, d int) bool {
	if a > b {
		a, b = b, a
	}
	if c > d {
		c, d = d, c
	}

	if a == c && b == d {
		return true
	}

	return (a < c && c < b && (d < a || d > b)) ||
		(a < d && d < b && (c < a || c > b))
}

// ========================
// PARSER
// ========================

type Pairs map[int]int

func ParseInput(file string) ([]int, Pairs, int) {
	data := utils.ReadFile(file)
	length := 0

	list := make([]int, 0)
	for number := range strings.SplitSeq(data, ",") {
		n, _ := strconv.Atoi(number)
		list = append(list, n)
		length = max(length, n)
	}

	pairs := make(Pairs)
	for i := range length / 2 {
		a := i + 1
		b := (i + 1) + (length / 2)
		pairs[a] = b
		pairs[b] = a
	}

	return list, pairs, length
}

// ========================
// ABS
// ========================

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
