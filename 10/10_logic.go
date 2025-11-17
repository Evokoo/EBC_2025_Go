package quest10

import (
	// "strconv"
	"strings"

	"github.com/Evokoo/EBC_2025_Go/utils"
)

// ========================
// PART I
// ========================
func I(dragon Dragon, sheep Sheep, grid Grid, rounds int) int {
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
func II(dragon Dragon, sheep Sheep, grid Grid, rounds int) int {
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

			sheep.MoveDown()

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

func III(dragon Dragon, sheep Sheep, grid Grid) {

}

// ========================
// POINT
// ========================
type Point struct{ x, y int }

func (p *Point) MoveDown() {
	(*p).y++
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
}

func (g *Grid) InRange(point Point) bool {
	return point.x >= 0 && point.x < g.cols && point.y >= 0 && point.y < g.rows
}
func (g *Grid) IsHut(point Point) bool {
	return g.hut.Has(point)
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
// SHEEP
// ========================
type Sheep = Queue

// ========================
// PARSER
// ========================
func ParseInput(file string) (Grid, Dragon, Sheep) {
	data := utils.ReadFile(file)
	matrix := strings.Split(data, "\n")
	grid := Grid{
		rows: len(matrix),
		cols: len(matrix[0]),
		hut:  make(Set[Point]),
	}

	var dragon Dragon
	var sheep Sheep

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
		}
	}

	return grid, dragon, sheep
}
