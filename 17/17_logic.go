package quest17

import (
	"strings"

	"github.com/Evokoo/EBC_2025_Go/utils"
)

// ========================
// PART I
// ========================
func I(grid Grid, radius int) (sum int) {
	for y := range grid.height {
		for x := range grid.width {
			if grid.WithinRadius(x, y, radius) {
				sum += grid.GetValue(x, y)
			}
		}
	}
	return sum
}

// ========================
// PART II
// ========================

func II(grid Grid) int {
	cx, cy := grid.volcano[0], grid.volcano[1]

	visited := make([][]bool, grid.height)
	for y := range visited {
		visited[y] = make([]bool, grid.width)

		if y == cy {
			visited[cy][cx] = true
		}
	}

	best := [2]int{0, 0}
	radius := 0

	for cx+radius < grid.width {
		score := 0
		for y := cy - radius; y <= cy+radius; y++ {
			for x := cx - radius; x <= cx+radius; x++ {
				if !visited[y][x] && grid.WithinRadius(x, y, radius) {
					score += grid.GetValue(x, y)
					visited[y][x] = true
				}
			}

		}
		if score > best[0] {
			best[0] = score
			best[1] = radius
		}

		radius++
	}

	return best[0] * best[1]
}

// ========================
// PART III
// ========================

func III(grid Grid) {
	//Djikstra & minHeap?
}

// ========================
// GRID
// ========================

type Grid struct {
	cells          [][]int
	volcano, start [2]int
	width, height  int
}

func (g Grid) WithinRadius(x, y, r int) bool {
	dx := g.volcano[0] - x
	dy := g.volcano[1] - y
	return dx*dx+dy*dy <= r*r
}
func (g Grid) InRange(x, y int) bool {
	return x >= 0 && x < g.width && y >= 0 && y < g.height
}
func (g Grid) GetValue(x, y int) int {
	return g.cells[y][x]
}

// ========================
// PARSER
// ========================
func ParseInput(file string) Grid {
	data := utils.ReadFile(file)
	cells := make([][]int, 0)
	volcano := [2]int{-1, -1}
	start := [2]int{-1, -1}

	for y, line := range strings.Split(data, "\n") {
		row := make([]int, len(line))

		for x, value := range line {
			switch value {
			case '@':
				volcano = [2]int{x, y}
			case 'S':
				start = [2]int{x, y}
			default:
				row[x] = int(value - '0')
			}
		}

		cells = append(cells, row)
	}

	return Grid{
		cells:   cells,
		volcano: volcano,
		start:   start,
		height:  len(cells),
		width:   len(cells[0]),
	}
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
