package quest16

import (
	"sort"
	"strconv"
	"strings"

	"github.com/Evokoo/EBC_2025_Go/utils"
)

// ========================
// WALL
// ========================
type Wall struct {
	blocks map[int]int
	count  int
}

func NewWall(length int) Wall {
	return Wall{
		blocks: make(map[int]int, length),
		count:  0,
	}
}
func (w *Wall) AddBlock(column int) {
	w.blocks[column]++
	w.count++
}

// ========================
// PART I
// ========================

func I(spell []int, length int) int {
	wall := NewWall(length)

	for _, n := range spell {
		for i := n; i <= length; i += n {
			wall.AddBlock(i)
		}
	}

	return wall.count
}

// ========================
// PART II
// ========================

func II(columns []int) (int, []int) {
	score := 1
	values := make(Set[int])

	array := make([]int, len(columns))
	copy(array, columns)

	for i, n := range array {
		index := i + 1

		if n > 0 {
			score *= index
			values.Add(index)

			for j := i; j < len(array); j += index {
				array[j]--
			}
		}
	}

	spell := make([]int, 0, len(values))
	for n := range values {
		spell = append(spell, n)
	}

	sort.Ints(spell)

	return score, spell
}

// ========================
// PART III
// ========================

func III(section []int, target int) int {
	_, spell := II(section)

	BlockCount := func(m int) (sum int) {
		for _, n := range spell {
			sum += m / n
		}
		return
	}

	upper, lower := target, 0
	for lower != upper {
		middle := (lower + upper + 1) / 2

		if BlockCount(middle) > target {
			upper = middle - 1
		} else {
			lower = middle
		}
	}

	return lower
}

// ========================
// PARSER
// ========================

func ParseInput(file string) []int {
	data := utils.ReadFile(file)
	values := make([]int, 0)

	for number := range strings.SplitSeq(data, ",") {
		n, _ := strconv.Atoi(number)
		values = append(values, n)
	}

	return values
}

// ========================
// SET
// ========================
type Set[T comparable] map[T]struct{}

func (s Set[T]) Add(v T) {
	s[v] = struct{}{}
}
func (s Set[T]) Has(v T) bool {
	_, ok := s[v]
	return ok
}
