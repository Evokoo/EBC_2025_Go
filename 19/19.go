package quest19

func Solve(file string, part int) int {
	course := ParseInput(file)

	switch part {
	case 1:
		return I(course)
	case 2:
		return I(course)
	}
	return 0
}
