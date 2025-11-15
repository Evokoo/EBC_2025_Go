package quest10

func Solve(file string, part int) int {
	grid := ParseInput(file)

	switch part {
	case 1:
		return I(grid, 5)
	case 2:
		return II(grid, 20)
	}
	return 0
}
