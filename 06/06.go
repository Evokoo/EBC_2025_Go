package quest06

func Solve(file string, part int) int {
	locations, len := ParseInput(file)

	switch part {
	case 1:
		return I(locations, true)
	case 2:
		return I(locations, false)
	case 3:
		return III(locations, 1000, 1000, len)
	}

	return 0
}
