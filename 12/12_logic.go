package quest12

import (
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

func (b *Barrels) ValidMoves(current [2]int, visited [][]bool) [][2]int {
	currentValue := b.GetValue(current)

	moves := make([][2]int, 0, 4)
	for _, dir := range DIRECTIONS {
		nx, ny := dir[0]+current[0], dir[1]+current[1]
		next := [2]int{nx, ny}

		if b.InRange(next) && !visited[ny][nx] && currentValue >= b.GetValue(next) {
			moves = append(moves, next)
		}
	}

	return moves
}

func (b *Barrels) GetValue(coord [2]int) int {
	return b.grid[coord[1]][coord[0]]
}

// ========================
// PART I & II
// ========================

func I(barrels Barrels, tracker Tracker, ignitionPoints [][2]int, maintain bool) int {
	queue := NewQueue[[2]int]()

	for _, point := range ignitionPoints {
		queue.Push(point)
	}

	count := 0

	for !queue.IsEmpty() {
		current := queue.Pop()
		x, y := current[0], current[1]

		if tracker.visited[y][x] {
			continue
		} else {
			tracker.visited[y][x] = true
			tracker.perRun = append(tracker.perRun, current)
			count++
		}

		for _, next := range barrels.ValidMoves(current, tracker.visited) {
			queue.Push(next)
		}
	}

	if !maintain {
		tracker.RemoveCurrentRun()
	}

	return count
}

// ========================
// PART III
// ========================

func III(barrels Barrels) int {
	tracker := NewTracker(barrels.rows, barrels.cols)

	count := 0

	for range 3 {
		maxCount := 0
		bestCoord := [2]int{}

		for y := 0; y < barrels.rows; y++ {
			for x := 0; x < barrels.cols; x++ {
				if tracker.visited[y][x] {
					continue
				}

				count := I(barrels, *tracker, [][2]int{{x, y}}, false)
				if count > maxCount {
					maxCount = count
					bestCoord = [2]int{x, y}
				}
			}
		}

		count += I(barrels, *tracker, [][2]int{bestCoord}, true)
	}

	return count
}

// ========================
// TRACKER
// ========================
type Tracker struct {
	visited [][]bool
	perRun  [][2]int
}

func NewTracker(rows, cols int) *Tracker {
	visited := make([][]bool, rows)
	for y := range visited {
		visited[y] = make([]bool, cols)
	}
	return &Tracker{
		visited: visited,
		perRun:  make([][2]int, 0, rows*cols),
	}
}

func (t *Tracker) RemoveCurrentRun() {
	for _, d := range t.perRun {
		t.visited[d[1]][d[0]] = false
	}
	t.perRun = t.perRun[:0]
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
