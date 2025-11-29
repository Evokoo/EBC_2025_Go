package quest20

import (
	"strings"

	"github.com/Evokoo/EBC_2025_Go/utils"
)

// ========================
// GRID
// ========================
type Trampolines struct {
	grid          [][]rune
	width, height int
}

var JUMPS = [][2]int{{-1, 1}, {0, 1}, {1, 1}, {-1, -1}, {0, -1}, {1, 1}}

func (t Trampolines) InRange(x, y int) bool {
	return x >= 0 && x < t.width && y >= 0 && y < t.height
}

func (t Trampolines) CountPairs(x, y int) int {
	count := 0
	for _, jump := range JUMPS[:3] {
		nx, ny := x+jump[0], y+jump[1]

		if t.InRange(nx, ny) && t.grid[ny][nx] == 'T' {
			count++
		}
	}
	return count
}

// ========================
// PART I
// ========================
func I(trampolines Trampolines) int {
	count := 0

	for y := 0; y < trampolines.height; y++ {
		for x := 0; x < trampolines.width; x++ {
			if trampolines.grid[y][x] == 'T' {
				count += trampolines.CountPairs(x, y)
			}
		}
	}

	return count
}

// ========================
// PARSER
// ========================

func ParseInput(file string) Trampolines {
	data := utils.ReadFile(file)
	grid := make([][]rune, 0)
	offset := 0

	for line := range strings.SplitSeq(data, "\n") {
		upper := make([]rune, len(line))
		lower := make([]rune, len(line))

		for i, r := range line {
			switch r {
			case '.':
				upper[i] = '.'
				lower[i] = '.'
			case 'T':
				if (i+offset)%2 == 0 {
					upper[i] = 'T'
					lower[i] = '#'
				} else {
					upper[i] = '#'
					lower[i] = 'T'
				}
			case '#':
				upper[i] = '#'
				lower[i] = '#'
			}
		}

		offset++
		grid = append(grid, [][]rune{upper, lower}...)
	}

	return Trampolines{grid: grid, height: len(grid), width: len(grid[0])}
}
