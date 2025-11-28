package quest19

import (
	"strconv"
	"strings"

	"github.com/Evokoo/EBC_2025_Go/utils"
)

// ========================
// COURSE
// ========================
type Course struct {
	gaps map[int][][2]int // x, [yMin, yMax]
	exit int
}

func (c Course) IsFreeSpace(x, y int) bool {
	if gaps, found := c.gaps[x]; found {
		for _, yRange := range gaps {
			if y >= yRange[0] && y <= yRange[1] {
				return true
			}
		}
		return false
	}

	return true
}

// ========================
// PART I
// ========================

var MOVES = [][3]int{{1, 1, 1}, {1, -1, 0}}

func QueueSorter(a, b [3]int) bool {
	return a[2] < b[2]
}

func I(course Course) int {
	queue := utils.NewPriorityQueue(QueueSorter)
	queue.PushItem([3]int{0, 0, 0})

	visited := make(map[[2]int]int)

	for !queue.IsEmpty() {
		bird := queue.PopItem()
		key := [2]int{bird[0], bird[1]}

		if previous, found := visited[key]; found && previous <= bird[2] {
			continue
		} else {
			visited[key] = bird[2]
		}

		if bird[0] == course.exit {
			return bird[2]
		}

		for _, move := range MOVES {
			nx, ny, nf := move[0]+bird[0], move[1]+bird[1], move[2]+bird[2]

			if course.IsFreeSpace(nx, ny) {
				queue.PushItem([3]int{nx, ny, nf})
			}
		}
	}

	return 0
}

// ========================
// PART III
// ========================

// Base on maneatingape's solution
// https://github.com/maneatingape/everybody-codes-rust/blob/main/src/event2025/quest19.rs

func III(file string) int {
	x, f := 0, 0

	for line := range strings.SplitSeq(utils.ReadFile(file), "\n") {
		value := strings.Split(line, ",")
		x1, _ := strconv.Atoi(value[0])
		y1, _ := strconv.Atoi(value[1])

		if x < x1 {
			x = x1
			f = max(f, (x1+y1+1)/2)
		}
	}

	return f
}

// ========================
// PARSER
// ========================

func ParseInput(file string) Course {
	data := utils.ReadFile(file)

	gaps := make(map[int][][2]int, 0)
	exit := 0

	for line := range strings.SplitSeq(data, "\n") {
		value := strings.Split(line, ",")
		x, _ := strconv.Atoi(value[0])
		y1, _ := strconv.Atoi(value[1])
		y2, _ := strconv.Atoi(value[2])

		if array, found := gaps[x]; found {
			array = append(array, [2]int{y1, y1 + y2})
			gaps[x] = array
		} else {
			gaps[x] = [][2]int{{y1, y1 + y2}}
		}

		exit = x
	}

	return Course{gaps: gaps, exit: exit}
}
