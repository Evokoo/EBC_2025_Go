package quest10

func Solve(file string, part int) int {
	grid := ParseInput(file)

	switch part {
	case 1:
		return I(grid, 5)
	}
	return 0
}
