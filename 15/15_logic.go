package quest15

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/Evokoo/EBC_2025_Go/utils"
)

// ========================
// LOOKUP
// ========================
type Lookup struct {
	cx, cy map[int]int
	rx, ry map[int]int
}

func GenerateLookup(xSet, ySet Set[int]) Lookup {
	xs := make([]int, 0, len(xSet))
	ys := make([]int, 0, len(ySet))

	for x := range xSet {
		xs = append(xs, x)
	}
	for y := range ySet {
		ys = append(ys, y)
	}

	sort.Ints(xs)
	sort.Ints(ys)

	cxMap := make(map[int]int)
	rxMap := make(map[int]int)
	for i, x := range xs {
		cxMap[x] = i
		rxMap[i] = x
	}

	cyMap := make(map[int]int)
	ryMap := make(map[int]int)
	for i, y := range ys {
		cyMap[y] = i
		ryMap[i] = y
	}

	return Lookup{cx: cxMap, cy: cyMap, rx: rxMap, ry: ryMap}
}
func (l Lookup) CompressXY(xy [2]int) [2]int {
	return [2]int{l.CompressX(xy[0]), l.CompressY(xy[1])}
}
func (l Lookup) CompressX(x int) int {
	return l.cx[x]
}
func (l Lookup) CompressY(y int) int {
	return l.cy[y]
}
func (l Lookup) ExpandXY(xy [2]int) [2]int {
	return [2]int{l.ExpandX(xy[0]), l.ExpandY(xy[1])}
}
func (l Lookup) ExpandX(x int) int {
	return l.rx[x]
}
func (l Lookup) ExpandY(y int) int {
	return l.ry[y]
}
func (l Lookup) GetDistance(a, b [2]int) int {
	from := l.ExpandXY(a)
	to := l.ExpandXY(b)

	return abs(from[0]-to[0]) + abs(from[1]-to[1])
}

// ========================
// WALLS
// ========================
type Walls struct {
	horizontal map[int][][2]int
	vertical   map[int][][2]int
}

func (w *Walls) AddPoints(from, to [2]int) {
	x1, y1, x2, y2 := from[0], from[1], to[0], to[1]

	if x1 == x2 { // Vertical
		a := min(y1, y2)
		b := max(y1, y2)
		w.vertical[x1] = append(w.vertical[x1], [2]int{a, b})
	} else { // Horizontal
		a := min(x1, x2)
		b := max(x1, x2)
		w.horizontal[y1] = append(w.horizontal[y1], [2]int{a, b})
	}
}
func (w Walls) UniquePoints(start, end [2]int) (Set[int], Set[int]) {
	ySet := make(Set[int])
	xSet := make(Set[int])

	for y, points := range w.horizontal {
		ySet.Add(y - 1)
		ySet.Add(y)
		ySet.Add(y + 1)

		for _, pair := range points {
			xSet.Add(pair[0] - 1)
			xSet.Add(pair[0])
			xSet.Add(pair[0] + 1)
			xSet.Add(pair[1] - 1)
			xSet.Add(pair[1])
			xSet.Add(pair[1] + 1)
		}
	}
	for x, points := range w.vertical {
		xSet.Add(x - 1)
		xSet.Add(x)
		xSet.Add(x + 1)

		for _, pair := range points {
			ySet.Add(pair[0] - 1)
			ySet.Add(pair[0])
			ySet.Add(pair[0] + 1)
			ySet.Add(pair[1] - 1)
			ySet.Add(pair[1])
			ySet.Add(pair[1] + 1)
		}
	}

	//Add Start
	ySet.Add(start[1])
	xSet.Add(start[0])

	//Add End
	ySet.Add(end[1])
	xSet.Add(end[0])

	return xSet, ySet
}
func (w Walls) CompressWalls(lookup Lookup) Set[[2]int] {
	walls := make(Set[[2]int])

	for y, points := range w.horizontal {
		ny := lookup.CompressY(y)

		for _, point := range points {
			x1, x2 := lookup.CompressX(point[0]), lookup.CompressX(point[1])
			if x1 > x2 {
				x1, x2 = x2, x1
			}

			for nx := x1; nx <= x2; nx++ {
				walls.Add([2]int{nx, ny})
			}
		}
	}

	for x, points := range w.vertical {
		nx := lookup.CompressX(x)
		for _, point := range points {
			y1, y2 := lookup.CompressY(point[0]), lookup.CompressY(point[1])
			if y1 > y2 {
				y1, y2 = y2, y1
			}

			for y := y1; y <= y2; y++ {
				walls.Add([2]int{nx, y})
			}
		}
	}

	return walls
}

// ========================
// TUNNEL
// ========================
type Tunnel struct {
	entry, exit   [2]int
	width, height int
	walls         Set[[2]int]
	lookup        Lookup
}

func NewTunnel(moves [][2]int) Tunnel {
	entry := [2]int{0, 0}
	exit := entry
	bearing := 0

	walls := Walls{
		horizontal: make(map[int][][2]int),
		vertical:   make(map[int][][2]int),
	}

	for _, move := range moves {
		bearing = (bearing + move[0] + 360) % 360
		from := exit

		switch bearing {
		case 0:
			exit[1] -= move[1]
		case 90:
			exit[0] += move[1]
		case 180:
			exit[1] += move[1]
		case 270:
			exit[0] -= move[1]
		}

		walls.AddPoints(from, exit)
	}

	start := [2]int{0, 0}
	lookup := GenerateLookup(walls.UniquePoints(start, exit))

	return Tunnel{
		entry:  lookup.CompressXY(entry),
		exit:   lookup.CompressXY(exit),
		width:  len(lookup.cx),
		height: len(lookup.cy),
		walls:  walls.CompressWalls(lookup),
		lookup: lookup,
	}
}
func (t Tunnel) InRange(coord [2]int) bool {
	return coord[0] >= 0 && coord[0] < t.width && coord[1] >= 0 && coord[1] < t.height
}
func (t Tunnel) PrintGraph() {
	// Create empty grid
	grid := make([][]rune, t.height)
	for y := 0; y < t.height; y++ {
		grid[y] = make([]rune, t.width)
		for x := 0; x < t.width; x++ {
			grid[y][x] = '.' // empty space
		}
	}

	// Place walls
	for wall := range t.walls {
		x, y := wall[0], wall[1]
		grid[y][x] = '#'
	}

	// Place start/end
	grid[t.entry[1]][t.entry[0]] = 'S'
	grid[t.exit[1]][t.exit[0]] = 'E'

	// Print the grid
	for y := 0; y < t.height; y++ {
		for x := 0; x < t.width; x++ {
			fmt.Printf("%c", grid[y][x])
		}
		fmt.Println()
	}
}

// ========================
// PART I, II & III
// ========================
var DIRECTIONS = [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

type State struct {
	position [2]int
	steps    int
}

func I(tunnel Tunnel) int {
	queue := make(Queue[State], 0)
	queue.Push(State{tunnel.entry, 0})

	seen := make(Set[[2]int])
	seen.Add(tunnel.entry)

	for !queue.IsEmpty() {
		current := queue.Pop()

		for _, dir := range DIRECTIONS {
			next := [2]int{current.position[0] + dir[0], current.position[1] + dir[1]}
			dist := tunnel.lookup.GetDistance(current.position, next)

			if next == tunnel.exit {
				return current.steps + dist
			}

			if tunnel.InRange(next) && !tunnel.walls.Has(next) && !seen.Has(next) {
				seen.Add(next)
				queue.Push(State{next, current.steps + dist})
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

// ========================
// ABS
// ========================
func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
