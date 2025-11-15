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
	currentRound := NewQueue(grid.dragon)
	nextRound := make(Queue, 0)
	seen := make(Set)

	for range rounds {
		for _, dragon := range currentRound {
			for _, direction := range DIRECTIONS {
				next := Point{dragon.x + direction.x, dragon.y + direction.y}

				if grid.InRange(next) && !seen.Has(next) {
					nextRound.Push(next)
				}
			}
			seen.Add(dragon)
		}
		currentRound = nextRound
		nextRound.Clear()
	}

	count := 0
	for _, sheep := range grid.sheep {
		if seen.Has(sheep) {
			count++
		}
	}
	return count
}

// ========================
// PART II
// ========================
func II(grid Grid, rounds int) int {
	dragonCurrent := NewQueue(grid.dragon)
	dragonNext := make(Queue, 0)

	sheepCurrent := grid.sheep
	sheepNext := make(Queue, 0)

	eaten := 0

	for range rounds {
		dragons := make(Set)

		// Dragon Moves
		for _, dragon := range dragonCurrent {
			for _, direction := range DIRECTIONS {
				next := Point{dragon.x + direction.x, dragon.y + direction.y}

				if grid.InRange(next) && !dragons.Has(next) {
					dragonNext.Push(next)
					dragons.Add(next)
				}
			}
		}

		//Sheep Moves
		for _, sheep := range sheepCurrent {
			if dragons.Has(sheep) && !grid.hut.Has(sheep) {
				eaten++
				continue
			}

			next := Point{sheep.x, sheep.y + 1}

			if next.y >= grid.rows {
				continue
			} else {
				if dragons.Has(next) && !grid.hut.Has(next) {
					eaten++
				} else {
					sheepNext.Push(next)
				}
			}
		}

		sheepCurrent = sheepNext
		sheepNext.Clear()
		dragonCurrent = dragonNext
		dragonNext.Clear()

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
func (q *Queue) Clear() {
	(*q) = make(Queue, 0)
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

//	func (g *Grid) IsSheep(point Point) bool {
//		return g.sheep.Has(point)
//	}
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
