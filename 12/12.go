package quest12

func Solve(file string, part int) int {
	barrels := ParseInput(file)

	a := [2]int{0, 0}
	b := [2]int{barrels.cols - 1, barrels.rows - 1}

	tracker := NewTracker(barrels.rows, barrels.cols)

	switch part {
	case 1:
		return I(barrels, tracker, [][2]int{a}, false)
	case 2:
		return I(barrels, tracker, [][2]int{a, b}, false)
	case 3:
		return III(barrels)
	}
	return 0
}
