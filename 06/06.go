package quest06

func Solve(file string, part int) int {
	locations := ParseInput(file)

	switch part {
	case 1:
		return I(locations, true)
	case 2:
		return I(locations, false)
	case 3:
		// locations = ParseInput(file, 1000)
		// return III()
	}

	return 0
}
