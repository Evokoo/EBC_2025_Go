package quest03

func Solve(file string, part int) int {
	boxes := ParseInput(file)

	switch part {
	case 1:
		return I(boxes)
	case 2:
		return II(boxes)
	case 3:
		return III(boxes)
	}

	return 0
}
