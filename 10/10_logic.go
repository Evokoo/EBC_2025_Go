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

func I(grid Grid, rounds int) int {
	dQueue := NewQueue(grid.dragon)
	occupied := make(Set)

	for range rounds {
		dLength := len(dQueue)

		for range dLength {
			dragon := dQueue.Pop()

			for _, direction := range DIRECTIONS {
				next := Point{dragon.x + direction.x, dragon.y + direction.y}

				if grid.InRange(next) && !occupied.Has(next) {
					dQueue.Push(next)
				}
			}

			occupied.Add(dragon)
		}
	}

	count := 0
	for _, sheep := range grid.sheep {
		if occupied.Has(sheep) {
			count++
		}
	}
	return count
}

// ========================
// PART II
// ========================
func II(grid Grid, rounds int) int {
	dQueue := NewQueue(grid.dragon)
	sQueue := grid.sheep
	eaten := 0

	for range rounds {
		occupied := make(Set)

		// Dragon Moves
		dLength := len(dQueue)
		for range dLength {
			dragon := dQueue.Pop()
			for _, direction := range DIRECTIONS {
				next := Point{dragon.x + direction.x, dragon.y + direction.y}

				if grid.InRange(next) && !occupied.Has(next) {
					dQueue.Push(next)
					occupied.Add(next)
				}
			}
		}

		//Sheep Moves
		sLength := len(sQueue)
		for range sLength {
			sheep := sQueue.Pop()

			if occupied.Has(sheep) && !grid.hut.Has(sheep) {
				eaten++
				continue
			}

			next := Point{sheep.x, sheep.y + 1}

			if next.y == grid.rows {
				continue
			} else {
				if occupied.Has(next) && !grid.hut.Has(next) {
					eaten++
				} else {
					sQueue.Push(next)
				}
			}
		}
	}

	return eaten
}

// ========================
// POINT
// ========================
type Point struct{ x, y int }

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
func (s *Set) Remove(coord Point) {
	delete(*s, coord)
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
// PARSER
// ========================

type Grid struct {
	cols   int
	rows   int
	sheep  Queue
	hut    Set
	dragon Point
}

func (g *Grid) InRange(point Point) bool {
	return point.x >= 0 && point.x < g.cols && point.y >= 0 && point.y < g.rows
}
func (g *Grid) IsHut(point Point) bool {
	return g.hut.Has(point)
}

func ParseInput(file string) Grid {
	data := utils.ReadFile(file)
	matrix := strings.Split(data, "\n")
	grid := Grid{
		rows: len(matrix),
		cols: len(matrix[0]),
		hut:  make(Set),
	}

	for y, row := range matrix {
		for x, r := range row {
			if r == 'S' {
				grid.sheep.Push(Point{x, y})
			}
			if r == 'D' {
				grid.dragon = Point{x, y}
			}
			if r == '#' {
				grid.hut.Add(Point{x, y})
			}
		}
	}

	return grid
}
