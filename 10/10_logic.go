package quest10

import (
	"strings"

	"github.com/Evokoo/EBC_2025_Go/utils"
)

// ========================
// PART I
// ========================
func I(dragon Dragon, sheep Queue, grid Grid, rounds int) int {
	queue := NewQueue(dragon)
	occupied := make(Set[Point])

	for range rounds {
		length := len(queue)

		for range length {
			current := queue.Pop()

			for _, move := range current.ValidMoves(grid) {
				if !occupied.Has(move) {
					queue.Push(move)
				}
			}
			occupied.Add(current)
		}
	}

	count := 0
	for _, sheep := range sheep {
		if occupied.Has(sheep) {
			count++
		}
	}

	return count
}

// ========================
// PART II
// ========================
func II(dragon Dragon, sheep Queue, grid Grid, rounds int) int {
	dQueue := NewQueue(dragon)
	sQueue := sheep
	eaten := 0

	for range rounds {
		occupied := make(Set[Point])

		// Dragon Moves
		length := len(dQueue)

		for range length {
			current := dQueue.Pop()

			for _, move := range current.ValidMoves(grid) {
				if !occupied.Has(move) {
					dQueue.Push(move)
					occupied.Add(move)
				}
			}
		}

		//Sheep Moves
		length = len(sQueue)
		for range length {
			sheep := sQueue.Pop()

			if occupied.Has(sheep) && !grid.hut.Has(sheep) {
				eaten++
				continue
			}

			sheep.y++

			if sheep.y == grid.rows {
				continue
			}

			if occupied.Has(sheep) && !grid.hut.Has(sheep) {
				eaten++
				continue
			}

			sQueue.Push(sheep)
		}
	}
	return eaten
}

// ========================
// PART III
// ========================
type Sheep [8]int

type State struct {
	sheepCache  map[uint64]uint64
	dragonCache map[uint64]uint64
}

func NewState() *State {
	return &State{
		sheepCache:  make(map[uint64]uint64),
		dragonCache: make(map[uint64]uint64),
	}
}

func key(dragon Point, sheep Sheep) uint64 {
	var key uint64
	// Dragon: 6 bits for x (3 bits) + y (3 bits)
	key = uint64(dragon.x)<<3 | uint64(dragon.y)
	// Sheep: 4 bits per column (if rows ≤ 15)
	for i := 0; i < 8; i++ {
		s := sheep[i] + 1 // shift -1 -> 0
		key = (key << 4) | uint64(s)
	}
	return key
}

func ConvertQueue(sheepQueue Queue) Sheep {
	var sheep Sheep
	for i := range sheep {
		sheep[i] = -1
	}
	for _, s := range sheepQueue {
		sheep[s.x] = s.y
	}
	return sheep
}

func III(dragon Dragon, sheepQueue Queue, grid Grid) uint64 {
	sheep := ConvertQueue(sheepQueue)
	state := NewState()
	return sheepMove(state, grid, dragon, sheep)
}

func sheepMove(state *State, grid Grid, dragon Point, sheep Sheep) uint64 {
	k := key(dragon, sheep)
	if val, ok := state.sheepCache[k]; ok {
		return val
	}

	result := uint64(0)
	moved := false

	for i, s := range sheep {
		if s == -1 { // inactive sheep
			continue
		}

		nextY := s + 1
		point := Point{x: i, y: nextY}

		// Moved off board or found instant exit square
		if !grid.InRange(point) || grid.IsExit(point) {
			moved = true
			continue
		}

		// Can't move onto dragon (unless hut)
		if point == dragon && !grid.IsHut(point) {
			continue
		}

		// Sheep moves forward
		moved = true
		nextSheep := sheep
		nextSheep[i] = nextY

		result += dragonMove(state, grid, dragon, nextSheep)
	}

	// If no sheep moved at all, dragon moves
	if !moved {
		result = dragonMove(state, grid, dragon, sheep)
	}

	state.sheepCache[k] = result
	return result
}
func dragonMove(state *State, grid Grid, dragon Point, sheep Sheep) uint64 {
	k := key(dragon, sheep)
	if val, ok := state.dragonCache[k]; ok {
		return val
	}

	result := uint64(0)
	for _, next := range dragon.ValidMoves(grid) {
		currentSheep := sheep

		// Eat a sheep if present and not in hut
		if currentSheep[next.x] == next.y && !grid.IsHut(next) {
			currentSheep[next.x] = -1
		}

		if AllSheepEaten(currentSheep) {
			result += 1
		} else {
			result += sheepMove(state, grid, next, currentSheep)
		}
	}

	state.dragonCache[k] = result
	return result
}

func AllSheepEaten(sheep Sheep) bool {
	for _, s := range sheep {
		if s != -1 {
			return false
		}
	}
	return true
}

// ========================
// POINT
// ========================
type Point struct{ x, y int }

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
type Queue []Point

func NewQueue(value Point) Queue {
	q := Queue{}
	q.Push(value)
	return q
}
func (q *Queue) Pop() Point {
	popped := (*q)[0]
	(*q) = (*q)[1:]
	return popped
}
func (q *Queue) Push(value Point) {
	(*q) = append((*q), value)
}
func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

// ========================
// GRID
// ========================
type Grid struct {
	cols int
	rows int
	hut  Set[Point]
	exit Set[Point]
}

func (g *Grid) InRange(point Point) bool {
	return point.x >= 0 && point.x < g.cols && point.y >= 0 && point.y < g.rows
}
func (g *Grid) IsHut(point Point) bool {
	return g.hut.Has(point)
}
func (g *Grid) IsExit(point Point) bool {
	return g.exit.Has(point)
}

// ========================
// DRAGON
// ========================
type Dragon = Point

var DIRECTIONS = []Point{{2, 1}, {2, -1}, {-2, 1}, {-2, -1}, {1, 2}, {1, -2}, {-1, 2}, {-1, -2}}

func (d *Dragon) ValidMoves(grid Grid) []Point {
	moves := make([]Point, 0, len(DIRECTIONS))

	for _, direction := range DIRECTIONS {
		next := Point{d.x + direction.x, d.y + direction.y}

		if grid.InRange(next) {
			moves = append(moves, next)
		}
	}

	return moves
}

// ========================
// PARSER
// ========================
func ParseInput(file string) (Grid, Dragon, Queue) {
	data := utils.ReadFile(file)
	matrix := strings.Split(data, "\n")
	grid := Grid{
		rows: len(matrix),
		cols: len(matrix[0]),
		hut:  make(Set[Point]),
		exit: make(Set[Point]),
	}

	var dragon Dragon
	var sheep Queue

	for y, row := range matrix {
		for x, r := range row {
			if r == 'S' {
				sheep.Push(Point{x, y})
			}
			if r == 'D' {
				dragon = Point{x, y}
			}
			if r == '#' {
				grid.hut.Add(Point{x, y})
			}
			if r == 'V' {
				grid.exit.Add(Point{x, y})
			}
		}
	}

	return grid, dragon, sheep
}
