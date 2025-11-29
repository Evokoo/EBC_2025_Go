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
// TRAMPOLINES
// ========================
type Trampolines struct {
	start, exit   [2]int
	grids         map[int][][]rune
	width, height int
}

func (t Trampolines) InRange(x, y int) bool {
	return x >= 0 && x < t.width && y >= 0 && y < t.height
}

func (t Trampolines) ValidJumps(x, y, r int, neutral bool) [][2]int {
	grid := t.grids[r]
	jumps := make([][2]int, 0, 4)

	if neutral && (grid[y][x] == 'T' || grid[y][x] == 'E' || grid[y][x] == 'S') {
		jumps = append(jumps, [2]int{x, y})
	}

	for _, jump := range GetMoveSet(x, y) {
		nx, ny := x+jump[0], y+jump[1]
		if t.InRange(nx, ny) && (grid[ny][nx] == 'T' || grid[ny][nx] == 'E' || grid[ny][nx] == 'S') {
			jumps = append(jumps, [2]int{nx, ny})
		}
	}

	return jumps
}

func (t *Trampolines) ComputeRotations() {
	for i := 0; i < 2; i++ {
		current := t.grids[i]
		rotation := make([][]rune, t.height)
		for y := 0; y < t.height; y++ {
			rotation[y] = make([]rune, t.width)
			copy(rotation[y], current[y])
		}

		for y := 0; y < t.height; y++ {
			cx, cy := t.start[0]+y, t.start[1]-y
			px, py := y, y

			for x := 0; x < t.width-(y*2); x++ {
				rotation[py][px] = current[cy][cx]

				if x%2 == 0 {
					cy--
				} else {
					cx--
				}
				px++
			}
		}
		t.grids[i+1] = rotation
	}
}

// ========================
// STATE
// ========================
type State struct {
	pos      [2]int
	steps    int
	rotation int
}

// ========================
// PART I
// ========================
func I(trampolines Trampolines) int {
	grid := trampolines.grids[0]
	seen := make(Set[[2]int])
	count := 0

	for y := 0; y < trampolines.height; y++ {
		for x := 0; x < trampolines.width; x++ {
			if grid[y][x] == 'T' {
				for _, jump := range trampolines.ValidJumps(x, y, 0, false) {
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

func II(trampolines Trampolines) int {
	queue := make(Queue[State], 0)
	queue.Push(State{trampolines.start, 0, 0})

	seen := make(Set[[2]int])
	seen.Add(trampolines.start)

	for !queue.IsEmpty() {
		current := queue.Pop()

		if current.pos == trampolines.exit {
			return current.steps
		}

		for _, jump := range trampolines.ValidJumps(current.pos[0], current.pos[1], 0, false) {
			if !seen.Has(jump) {
				queue.Push(State{jump, current.steps + 1, 0})
				seen.Add(jump)
			}
		}
	}

	return 0
}

// ========================
// PART III
// ========================

func III(t Trampolines) int {
	//Compute rotations
	t.ComputeRotations()

	queue := make(Queue[State], 0)
	queue.Push(State{t.start, 0, 0})

	seen := make(Set[[3]int])

	for !queue.IsEmpty() {
		current := queue.Pop()
		key := [3]int{current.pos[0], current.pos[1], current.rotation}
		rotation := (current.rotation + 1) % 3

		if current.pos == t.exit {
			return current.steps
		}

		if seen.Has(key) {
			continue
		} else {
			seen.Add(key)
		}

		for _, jump := range t.ValidJumps(current.pos[0], current.pos[1], rotation, true) {
			queue.Push(State{jump, current.steps + 1, rotation})
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

	grids := make(map[int][][]rune, 3)
	grids[0] = grid

	return Trampolines{
		grids:  grids,
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
