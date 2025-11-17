package quest10

func Solve(file string, part int) int {
	grid, dragon, sheep := ParseInput(file)

	switch part {
	case 1:
		return I(dragon, sheep, grid, 5)
	case 2:
		return II(dragon, sheep, grid, 20)
	case 3:
		return int(III(dragon, sheep, grid))
	}
	return 0
}
