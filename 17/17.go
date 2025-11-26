package quest17

func Solve(file string, part int) int {
	grid := ParseInput(file)

	switch part {
	case 1:
		return I(grid, 10)
	case 2:
		return II(grid)
	case 3:
		III(grid)
	}

	return 0
}
