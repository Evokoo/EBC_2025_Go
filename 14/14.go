package quest14

func Solve(file string, part int) int {

	switch part {
	case 1:
		return I(ParseInput(file), 10)
	case 2:
		return I(ParseInput(file), 2025)
	case 3:
		return III(ParseInputIII(file), 1000000000)
	}
	return 0
}
