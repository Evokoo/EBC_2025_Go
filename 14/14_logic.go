package quest14

import (
	"fmt"
	"strings"

	"github.com/Evokoo/EBC_2025_Go/utils"
)

// ========================
// TILES
// ========================
type Target struct {
	x, y   int
	status bool
}

type Tiles struct {
	grid       [][]bool
	rows, cols int
	updates    Queue[[2]int]
	pattern    []Target
}

var DIRECTIONS = [][2]int{{1, 1}, {1, -1}, {-1, 1}, {-1, -1}}

func (t Tiles) GetNewStatus(x, y int) bool {
	active := 0
	for _, dir := range DIRECTIONS {
		nx, ny := x+dir[0], y+dir[1]

		if t.InRange(nx, ny) && t.IsActive(nx, ny) {
			active++
		}
	}

	if t.IsActive(x, y) {
		return active%2 != 0
	} else {
		return active%2 == 0
	}
}
func (t Tiles) IsActive(x, y int) bool {
	return t.grid[y][x]
}
func (t Tiles) InRange(x, y int) bool {
	return x >= 0 && x < t.cols && y >= 0 && y < t.rows
}
func (t *Tiles) UpdateTiles() {
	for !t.updates.IsEmpty() {
		xy := t.updates.Pop()
		t.grid[xy[1]][xy[0]] = !t.grid[xy[1]][xy[0]]
	}
}
func (t Tiles) CountActive() int {
	count := 0
	for y := 0; y < t.rows; y++ {
		for x := 0; x < t.cols; x++ {
			if t.IsActive(x, y) {
				count++
			}
		}
	}
	return count
}
func (t Tiles) HasPattern() bool {
	for _, tile := range t.pattern {
		if t.grid[tile.y][tile.x] != tile.status {
			return false
		}
	}

	for y := 0; y < t.rows/2; y++ {
		for x := 0; x < t.cols/2; x++ {
			a := t.grid[y][x]
			b := t.grid[t.rows-1-y][x]
			c := t.grid[y][t.cols-1-x]
			d := t.grid[t.rows-1-y][t.cols-1-x]

			if a != b || a != c || a != d {
				return false

			}
		}
	}

	return true
}
func (t Tiles) PrintPattern() {
	var line strings.Builder

	for _, row := range t.grid {
		for _, value := range row {
			if value {
				line.WriteString("#")
			} else {
				line.WriteString(".")
			}
		}

		fmt.Println(line.String())
		line.Reset()
	}

	fmt.Println()
}

// ========================
// PART I & II
// ========================
func I(tiles Tiles, rounds int) int {
	total := 0

	for range rounds {
		for y := 0; y < tiles.rows; y++ {
			for x := 0; x < tiles.cols; x++ {
				newStatus := tiles.GetNewStatus(x, y)
				if tiles.IsActive(x, y) != newStatus {
					tiles.updates.Push([2]int{x, y})
				}
			}
		}

		tiles.UpdateTiles()
		total += tiles.CountActive()
	}

	return total
}

// ========================
// PART III
// ========================
func III(tiles Tiles, rounds int) int {
	seen := make(map[int]struct{})
	cycles := make([][2]int, 0)
	total := 0

	for round := range rounds {
		for y := 0; y < tiles.rows; y++ {
			for x := 0; x < tiles.cols; x++ {
				newStatus := tiles.GetNewStatus(x, y)
				if tiles.IsActive(x, y) != newStatus {
					tiles.updates.Push([2]int{x, y})
				}
			}
		}

		tiles.UpdateTiles()

		if tiles.HasPattern() {
			count := tiles.CountActive()
			cycles = append(cycles, [2]int{count, round})

			if _, found := seen[count]; found {
				return PredictCumulative(cycles, rounds)
			}

			seen[count] = struct{}{}
			total += count
		}
	}

	return 0
}

func PredictCumulative(cycles [][2]int, round int) int {
	var cycleLength int

	for _, c := range cycles[1:] {
		cycleLength++
		if c[0] == cycles[0][0] {
			break
		}
	}

	c1 := cycles[:cycleLength]

	fullCycleRounds := cycles[cycleLength][1] - c1[0][1]
	roundsToCover := round - c1[0][1]
	fullCycles := roundsToCover / fullCycleRounds

	totalTiles := fullCycles * SumCycle(c1)
	remainderRounds := roundsToCover % fullCycleRounds

	for _, step := range c1 {
		recordedRoundInCycle := step[1]
		tileCount := step[0]

		if recordedRoundInCycle <= remainderRounds {
			totalTiles += tileCount
		} else {
			break
		}
	}

	return totalTiles
}

func SumCycle(cycle [][2]int) int {
	sum := 0
	for _, c := range cycle {
		sum += c[0]
	}
	return sum
}

// ========================
// PARSER
// ========================

func ParseInput(file string) Tiles {
	data := utils.ReadFile(file)
	lines := strings.Split(data, "\n")

	rows := len(lines)
	cols := len(lines[0])

	grid := make([][]bool, rows)
	for y, row := range lines {
		for _, col := range row {
			switch col {
			case '.':
				grid[y] = append(grid[y], false)
			case '#':
				grid[y] = append(grid[y], true)
			}
		}
	}

	updates := make(Queue[[2]int], 0, rows*cols)
	return Tiles{grid: grid, updates: updates, rows: rows, cols: cols}
}

func ParseInputIII(file string) Tiles {
	data := utils.ReadFile(file)
	lines := strings.Split(data, "\n")
	rows := len(lines)
	cols := len(lines[0])

	target := make([]Target, 0, 64)
	grid := make([][]bool, 34)
	for i := range 34 {
		grid[i] = make([]bool, 34)
	}

	xOffset := (34 - cols) / 2
	yOffset := (34 - rows) / 2

	for y := range rows {
		for x := range cols {
			nx, ny := x+xOffset, y+yOffset

			if lines[y][x] == '#' {
				target = append(target, Target{nx, ny, true})
			} else {
				target = append(target, Target{nx, ny, false})
			}
		}
	}

	return Tiles{
		grid:    grid,
		updates: make(Queue[[2]int], 0, 34*34),
		pattern: target,
		rows:    34,
		cols:    34,
	}
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
