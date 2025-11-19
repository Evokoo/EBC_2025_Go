package quest12

import (
	"fmt"
	"strings"

	"github.com/Evokoo/EBC_2025_Go/utils"
)

// ========================
// BARRELS
// ========================

type Barrels struct {
	rows int
	cols int
	grid [][]int
}

var DIRECTIONS = [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func (b *Barrels) InRange(coord [2]int) bool {
	x := coord[0]
	y := coord[1]
	return x >= 0 && x < b.cols && y >= 0 && y < b.rows
}

func (b *Barrels) ValidMoves(current [2]int, seen Set[[2]int]) [][2]int {
	currentValue := b.GetValue(current)

	var moves [][2]int
	for _, dir := range DIRECTIONS {
		next := [2]int{dir[0] + current[0], dir[1] + current[1]}

		if b.InRange(next) && !seen.Has(next) && currentValue >= b.GetValue(next) {
			moves = append(moves, next)
		}
	}
	return moves
}

func (b *Barrels) GetValue(coord [2]int) int {
	return b.grid[coord[1]][coord[0]]
}

func (b *Barrels) IsPeak(current [2]int) bool {
	value := b.GetValue(current)

	for _, dir := range DIRECTIONS {
		next := [2]int{dir[0] + current[0], dir[1] + current[1]}

		var adjacentValue int
		if !b.InRange(next) {
			adjacentValue = 0
		} else {
			adjacentValue = b.GetValue(next)
		}

		if value < adjacentValue {
			return false
		}
	}
	return true
}

// ========================
// PART I & II
// ========================

func I(barrels Barrels, seen Set[[2]int], ignitionPoints [][2]int) Set[[2]int] {
	queue := NewQueue[[2]int]()

	for _, point := range ignitionPoints {
		queue.Push(point)
	}

	for !queue.IsEmpty() {
		current := queue.Pop()

		if seen.Has(current) {
			continue
		} else {
			seen.Add(current)
		}

		for _, next := range barrels.ValidMoves(current, seen) {
			queue.Push(next)
		}
	}

	return seen
}

// ========================
// PART III
// ========================

type Best struct {
	coord     [2]int
	count     int
	destoryed Set[[2]int]
}

func III(barrels Barrels) {
	roundOne := Best{
		coord:     [2]int{},
		count:     0,
		destoryed: make(Set[[2]int]),
	}

	// B1
	for y := 0; y < barrels.rows; y++ {
		for x := 0; x < barrels.cols; x++ {
			current := [2]int{x, y}

			destoryed := I(barrels, make(Set[[2]int]), [][2]int{current})

			if destoryed.Length() > roundOne.count {
				roundOne.coord = current
				roundOne.count = destoryed.Length()
				roundOne.destoryed = destoryed
			}
		}
	}

	roundTwo := Best{
		coord:     [2]int{},
		count:     0,
		destoryed: make(Set[[2]int]),
	}

	// B2
	for y := 0; y < barrels.rows; y++ {
		for x := 0; x < barrels.cols; x++ {
			current := [2]int{x, y}

			if roundOne.destoryed.Has(current) {
				continue
			}

			destoryed := I(barrels, roundOne.destoryed.Clone(), [][2]int{current})

			if destoryed.Length() > roundTwo.count {
				roundTwo.coord = current
				roundTwo.count = destoryed.Length()
				roundTwo.destoryed = destoryed
			}
		}
	}

	roundThee := Best{
		coord:     [2]int{},
		count:     0,
		destoryed: make(Set[[2]int]),
	}

	// B3
	for y := 0; y < barrels.rows; y++ {
		for x := 0; x < barrels.cols; x++ {
			current := [2]int{x, y}

			if roundTwo.destoryed.Has(current) {
				continue
			}

			destoryed := I(barrels, roundTwo.destoryed.Clone(), [][2]int{current})

			if destoryed.Length() > roundThee.count {
				roundThee.coord = current
				roundThee.count = destoryed.Length()
				roundThee.destoryed = destoryed
			}
		}
	}

	fmt.Println(roundThee)
}

// ========================
// PARSER
// ========================

func ParseInput(file string) Barrels {
	data := utils.ReadFile(file)

	var grid [][]int

	for line := range strings.SplitSeq(data, "\n") {
		row := make([]int, len(line))

		for i, digit := range line {
			row[i] = int(digit - '0')
		}

		grid = append(grid, row)
	}

	return Barrels{rows: len(grid), cols: len(grid[0]), grid: grid}
}

// ========================
// QUEUE
// ========================
type Queue[T comparable] []T

func NewQueue[T comparable]() Queue[T] {
	return Queue[T]{}
}

func (q *Queue[T]) Pop() T {
	popped := (*q)[0]
	*q = (*q)[1:]
	return popped
}

func (q *Queue[T]) Push(value T) {
	*q = append(*q, value)
}

func (q *Queue[T]) IsEmpty() bool {
	return len(*q) == 0
}

// ========================
// SET
// ========================
type Set[T comparable] map[T]struct{}

func (s *Set[T]) Add(v T) {
	(*s)[v] = struct{}{}
}

func (s Set[T]) Has(v T) bool {
	_, ok := (s)[v]
	return ok
}

func (s Set[T]) Length() int {
	return len(s)
}

func (s Set[T]) Clone() Set[T] {
	newSet := make(Set[T], s.Length())
	for k := range s {
		newSet[k] = struct{}{}
	}
	return newSet
}
