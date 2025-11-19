package quest12

import (
	"strings"

	"github.com/Evokoo/EBC_2025_Go/utils"
)

// ========================
// BARRELS
// ========================

type BarrelMap struct {
	rows, cols int
	grid       [][]int
	cache      [][][][2]int
	moves      [][2]int
}

func NewBarrelMap(grid [][]int) BarrelMap {
	rows := len(grid)
	cols := len(grid[0])

	cache := make([][][][2]int, rows)
	for y := range rows {
		cache[y] = make([][][2]int, cols)
		for x := range cols {
			cache[y][x] = make([][2]int, 0, 4)
		}
	}

	barrels := BarrelMap{
		rows:  rows,
		cols:  cols,
		grid:  grid,
		cache: cache,
		moves: [][2]int{},
	}

	barrels.CacheMoves()

	return barrels
}

var DIRECTIONS = [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func (b *BarrelMap) CacheMoves() {
	for y := 0; y < b.rows; y++ {
		for x := 0; x < b.cols; x++ {
			current := [2]int{x, y}

			for _, dir := range DIRECTIONS {
				nx, ny := dir[0]+current[0], dir[1]+current[1]
				next := [2]int{nx, ny}

				if nx >= 0 && nx < b.cols && ny >= 0 && ny < b.rows && b.grid[y][x] >= b.grid[ny][nx] {
					b.cache[y][x] = append(b.cache[y][x], next)
				}
			}
		}
	}
}

func (b *BarrelMap) ValidMoves(current [2]int, visited [][]bool) [][2]int {
	x, y := current[0], current[1]
	b.moves = b.moves[:0]

	for _, move := range b.cache[y][x] {
		if !visited[move[1]][move[0]] {
			b.moves = append(b.moves, move)
		}
	}
	return b.moves
}

// ========================
// PART I & II
// ========================

func I(barrels BarrelMap, tracker *Tracker, ignitionPoints [][2]int, maintain bool) int {
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
			tracker.temporary = append(tracker.temporary, current)
			count++
		}

		for _, next := range barrels.ValidMoves(current, tracker.visited) {
			queue.Push(next)
		}
	}

	if !maintain {
		tracker.ClearCurrentRun()
	} else {
		tracker.ClearTemporary()
	}

	return count
}

// ========================
// PART III
// ========================

func III(barrels BarrelMap) int {
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

				count := I(barrels, tracker, [][2]int{{x, y}}, false)

				if count > maxCount {
					maxCount = count
					bestCoord = [2]int{x, y}
				}
			}
		}

		count += I(barrels, tracker, [][2]int{bestCoord}, true)
	}

	return count
}

// ========================
// TRACKER
// ========================
type Tracker struct {
	visited   [][]bool
	temporary [][2]int
}

func NewTracker(rows, cols int) *Tracker {
	visited := make([][]bool, rows)
	for y := range visited {
		visited[y] = make([]bool, cols)
	}
	return &Tracker{
		visited:   visited,
		temporary: make([][2]int, 0, rows*cols),
	}
}

func (t *Tracker) ClearCurrentRun() {
	for _, d := range t.temporary {
		t.visited[d[1]][d[0]] = false
	}
	t.ClearTemporary()
}

func (t *Tracker) ClearTemporary() {
	t.temporary = t.temporary[:0]
}

// ========================
// PARSER
// ========================

func ParseInput(file string) BarrelMap {
	data := utils.ReadFile(file)

	var grid [][]int

	for line := range strings.SplitSeq(data, "\n") {
		row := make([]int, len(line))
		for i, digit := range line {
			row[i] = int(digit - '0')
		}
		grid = append(grid, row)
	}
	return NewBarrelMap(grid)
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
