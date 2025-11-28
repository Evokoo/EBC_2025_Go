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
	gaps map[int][2]int // x, [yMin, yMax]
	exit int
}

func (c Course) IsWall(x, y int) bool {
	if yRange, found := c.gaps[x]; found {
		return y < yRange[0] || y > yRange[1]
	}

	return false
}

// ========================
// PART I
// ========================

var MOVES = [][3]int{{1, 1, 1}, {1, -1, 0}}

type Bird struct{ x, y, flaps int }

func QueueSorter(a, b Bird) bool {
	return a.flaps < b.flaps
}

func I(course Course) int {
	queue := utils.NewPriorityQueue(QueueSorter)
	queue.PushItem(Bird{x: 0, y: 0, flaps: 0})

	visited := make(map[[2]int]int)

	for !queue.IsEmpty() {
		bird := queue.PopItem()
		key := [2]int{bird.x, bird.y}

		if previous, found := visited[key]; found && previous <= bird.flaps {
			continue
		} else {
			visited[key] = bird.flaps
		}

		if bird.x == course.exit {
			return bird.flaps
		}

		for _, move := range MOVES {
			nx, ny, nf := move[0]+bird.x, move[1]+bird.y, move[2]+bird.flaps

			if !course.IsWall(nx, ny) {
				queue.PushItem(Bird{x: nx, y: ny, flaps: nf})
			}
		}
	}

	return 0
}

// ========================
// PARSER
// ========================

func ParseInput(file string) Course {
	data := utils.ReadFile(file)

	gaps := make(map[int][2]int, 0)
	exit := 0

	for line := range strings.SplitSeq(data, "\n") {
		value := strings.Split(line, ",")
		x, _ := strconv.Atoi(value[0])
		y1, _ := strconv.Atoi(value[1])
		y2, _ := strconv.Atoi(value[2])

		gaps[x] = [2]int{y1, y1 + y2}
		exit = x
	}

	return Course{gaps: gaps, exit: exit}
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
