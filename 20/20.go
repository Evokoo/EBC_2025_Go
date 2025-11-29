package quest20

func Solve(file string, part int) int {
	trampolines := ParseInput(file)

	switch part {
	case 1:
		return I(trampolines)
	case 2:
		return II(trampolines)
	case 3:
		return III(trampolines)
	}
	return 0
}
