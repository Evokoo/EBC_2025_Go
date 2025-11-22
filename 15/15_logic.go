package quest15

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/Evokoo/EBC_2025_Go/utils"
)

// ========================
// TUNNEL
// ========================
type Tunnel struct {
	start, end    [2]int
	width, height [2]int
	walls         Set[[2]int]
}

func NewTunnel(moves [][2]int) Tunnel {
	tunnel := Tunnel{
		start:  [2]int{0, 0},
		end:    [2]int{0, 0},
		width:  [2]int{math.MaxInt, math.MinInt},
		height: [2]int{math.MaxInt, math.MinInt},
		walls:  make(Set[[2]int]),
	}

	bearing := 0
	tunnel.walls.Add(tunnel.start)

	for _, move := range moves {
		bearing = (bearing + move[0] + 360) % 360

		for range move[1] {
			switch bearing {
			case 0:
				tunnel.end[1]--
			case 90:
				tunnel.end[0]++
			case 180:
				tunnel.end[1]++
			case 270:
				tunnel.end[0]--
			}

			tunnel.walls.Add(tunnel.end)
		}

		tunnel.width[0] = min(tunnel.width[0], tunnel.end[0])
		tunnel.width[1] = max(tunnel.width[1], tunnel.end[0])
		tunnel.height[0] = min(tunnel.height[0], tunnel.end[1])
		tunnel.height[1] = max(tunnel.height[1], tunnel.end[1])
	}

	return tunnel
}
func (t Tunnel) InRange(coord [2]int) bool {
	return coord[0] >= t.width[0] && coord[0] < t.width[1] && coord[1] >= t.height[0] && coord[1] < t.height[1]
}
func (t Tunnel) DrawTunnel() {
	cols := t.width[1] - t.width[0] + 1
	rows := t.height[1] - t.height[0] + 1

	var grid [][]string
	for y := 0; y < rows; y++ {
		grid = append(grid, make([]string, cols))
		for x := 0; x < cols; x++ {
			grid[y][x] = "."
		}
	}

	for wall := range t.walls {
		x, y := wall[0], wall[1]
		grid[y-t.height[0]][x-t.width[0]] = "#"
	}

	grid[t.start[1]-t.height[0]][t.start[0]-t.width[0]] = "S"
	grid[t.end[1]-t.height[0]][t.end[0]-t.width[0]] = "E"

	for y := range grid {
		var row strings.Builder
		for x := range grid[y] {
			row.WriteString(grid[y][x])
		}
		fmt.Println(row.String())
	}

	fmt.Println()
}

// ========================
// PART I
// ========================
var DIRECTIONS = [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

type State struct {
	position [2]int
	steps    int
}

func I(tunnel Tunnel) int {
	queue := make(Queue[State], 0)
	queue.Push(State{tunnel.start, 0})
	seen := make(Set[[2]int])

	for !queue.IsEmpty() {
		current := queue.Pop()

		if seen.Has(current.position) {
			continue
		} else {
			seen.Add(current.position)
		}

		for _, dir := range DIRECTIONS {
			next := [2]int{current.position[0] + dir[0], current.position[1] + dir[1]}

			if next == tunnel.end {
				return current.steps + 1
			}

			if tunnel.InRange(next) && !tunnel.walls.Has(next) && !seen.Has(next) {
				queue.Push(State{next, current.steps + 1})
			}
		}
	}

	return 0
}

// ========================
// PARSER
// ========================

func ParseInput(file string) Tunnel {
	data := utils.ReadFile(file)

	var moves [][2]int
	for instruction := range strings.SplitSeq(data, ",") {
		turn := instruction[0]
		step, _ := strconv.Atoi(instruction[1:])

		switch turn {
		case 'R':
			moves = append(moves, [2]int{90, step})
		case 'L':
			moves = append(moves, [2]int{-90, step})
		default:
			panic("Invalid direction")
		}
	}

	return NewTunnel(moves)
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
