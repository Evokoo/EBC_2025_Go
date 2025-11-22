package quest15

func Solve(file string, part int) int {
	tunnel := ParseInput(file)

	switch part {
	case 1:
		return I(tunnel)
	case 2:
		return I(tunnel)
	}
	return 0
}
