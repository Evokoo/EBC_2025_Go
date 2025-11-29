package quest20

func Solve(file string, part int) int {
	trampolines := ParseInput(file)

	switch part {
	case 1:
		return I(trampolines)
	}
	return 0
}
