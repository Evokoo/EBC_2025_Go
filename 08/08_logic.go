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
func II(list []int, pairs Pairs, nails int) int {
	seen := make(map[[2]int]struct{})
	knots := 0

	for i := range len(list) - 1 {
		a := list[i]
		b := list[i+1]

		for key, _ := range seen {
			if IsIntersection(a, b, key[0], key[1]) {
				knots++
			}
		}

		seen[[2]int{a, b}] = struct{}{}
	}

	return knots
}

func IsIntersection(a, b, c, d int) bool {
	if a > b {
		a, b = b, a
	}

	if c < d {
		c, d = d, c
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
