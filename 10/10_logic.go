package quest10

import (
	"strings"

	"github.com/Evokoo/EBC_2025_Go/utils"
)

// ========================
// DIRECTIONS
// ========================
var DIRECTIONS = []Point{{2, 1}, {2, -1}, {-2, 1}, {-2, -1}, {1, 2}, {1, -2}, {-1, 2}, {-1, -2}}

// ========================
// PART I
// ========================

func I(grid Grid, turns int) int {
	queue := NewQueue(Dragon{grid.dragon, 0})
	seen := make(Set)

	for !queue.IsEmpty() {
		current := queue.Pop()

		if current.moves == turns {
			continue
		}

		for _, direction := range DIRECTIONS {
			next := Point{current.position.x + direction.x, current.position.y + direction.y}

			if grid.InRange(next) && !seen.Has(next) {
				queue.Push(Dragon{next, current.moves + 1})
			}
		}

		seen.Add(current.position)
	}

	count := 0
	for coord := range seen {
		if grid.IsSheep(coord) {
			count++
		}
	}
	return count
}

// ========================
// POINT
// ========================
type Point struct{ x, y int }

// ========================
// DRAGON
// ========================
type Dragon struct {
	position Point
	moves    int
}

// ========================
// SET
// ========================
type Set map[Point]struct{}

func (s *Set) Add(coord Point) {
	(*s)[coord] = struct{}{}
}
func (s *Set) Has(coord Point) bool {
	_, found := (*s)[coord]
	return found
}

// ========================
// QUEUE
// ========================
type Queue []Dragon

func NewQueue(value Dragon) Queue {
	q := Queue{}
	q.Push(value)
	return q
}
func (q *Queue) Pop() Dragon {
	popped := (*q)[0]
	(*q) = (*q)[1:]
	return popped
}
func (q *Queue) Push(value Dragon) {
	(*q) = append((*q), value)
}
func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

// ========================
// PARSER
// ========================

type Grid struct {
	cols   int
	rows   int
	sheep  Set
	dragon Point
}

func (g *Grid) InRange(point Point) bool {
	return point.x >= 0 && point.x < g.cols && point.y >= 0 && point.y < g.rows
}
func (g *Grid) IsSheep(point Point) bool {
	return g.sheep.Has(point)
}

func ParseInput(file string) Grid {
	data := utils.ReadFile(file)
	matrix := strings.Split(data, "\n")
	grid := Grid{
		rows:  len(matrix),
		cols:  len(matrix[0]),
		sheep: make(Set),
	}

	for y, row := range matrix {
		for x, r := range row {
			if r == 'S' {
				grid.sheep.Add(Point{x, y})
			}
			if r == 'D' {
				grid.dragon = Point{x, y}
			}
		}
	}

	return grid
}
