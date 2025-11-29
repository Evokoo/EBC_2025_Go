package quest20

import (
	"strings"

	"github.com/Evokoo/EBC_2025_Go/utils"
)

// ========================
// MOVES
// ========================
var EVEN_DIRS = [][2]int{{1, 0}, {-1, 0}, {0, -1}}
var ODD_DIRS = [][2]int{{1, 0}, {-1, 0}, {0, 1}}

func GetMoveSet(x, y int) [][2]int {
	if max(0, x-y)%2 == 0 {
		return EVEN_DIRS
	}
	return ODD_DIRS
}

// ========================
// GRID
// ========================
type Trampolines struct {
	start, exit   [2]int
	grid          [][]rune
	width, height int
}

func (t Trampolines) InRange(x, y int) bool {
	return x >= 0 && x < t.width && y >= 0 && y < t.height
}
func (t Trampolines) ValidJumps(x, y int) [][2]int {
	jumps := make([][2]int, 0, 3)
	for _, jump := range GetMoveSet(x, y) {
		nx, ny := x+jump[0], y+jump[1]

		if t.InRange(nx, ny) && (t.grid[ny][nx] == 'T' || t.grid[ny][nx] == 'E' || t.grid[ny][nx] == 'S') {
			jumps = append(jumps, [2]int{nx, ny})
		}
	}
	return jumps
}

// ========================
// PART I
// ========================
func I(trampolines Trampolines) int {
	seen := make(Set[[2]int])
	count := 0

	for y := 0; y < trampolines.height; y++ {
		for x := 0; x < trampolines.width; x++ {
			if trampolines.grid[y][x] == 'T' {
				for _, jump := range trampolines.ValidJumps(x, y) {
					if !seen.Has(jump) {
						count++
					}
				}
			}
			seen.Add([2]int{x, y})
		}
	}
	return count
}

// ========================
// PART II
// ========================
type State struct {
	position [2]int
	steps    int
}

func II(trampolines Trampolines) int {
	queue := make(Queue[State], 0)
	queue.Push(State{trampolines.start, 0})

	seen := make(Set[[2]int])
	seen.Add(trampolines.start)

	for !queue.IsEmpty() {
		current := queue.Pop()

		if current.position == trampolines.exit {
			return current.steps
		}

		for _, jump := range trampolines.ValidJumps(current.position[0], current.position[1]) {
			if !seen.Has(jump) {
				queue.Push(State{jump, current.steps + 1})
				seen.Add(jump)
			}
		}
	}

	return 0
}

// ========================
// PARSER
// ========================

func ParseInput(file string) Trampolines {
	data := utils.ReadFile(file)
	grid := make([][]rune, 0)

	var start, exit [2]int

	for y, line := range strings.Split(data, "\n") {
		row := make([]rune, len(line))

		for x, r := range line {
			if r == 'E' {
				exit = [2]int{x, y}
			}
			if r == 'S' {
				start = [2]int{x, y}
			}
			row[x] = r
		}

		grid = append(grid, row)
	}

	return Trampolines{
		grid:   grid,
		start:  start,
		exit:   exit,
		height: len(grid),
		width:  len(grid[0])}
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

// ========================
// QUEUE
// ========================
type Queue[T comparable] []T

func (q *Queue[T]) Pop() T {
	popped := (*q)[0]
	(*q) = (*q)[1:]
	return popped
}
func (q *Queue[T]) Push(value T) {
	(*q) = append((*q), value)
}
func (q *Queue[T]) IsEmpty() bool {
	return len(*q) == 0
}
