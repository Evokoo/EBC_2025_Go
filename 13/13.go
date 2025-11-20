package quest13

func Solve(file string, part int) int {

	switch part {
	case 1:
		return I(file, 2025)
	case 2:
		return II(file, 20252025)
	case 3:
		return II(file, 202520252025)
	}

	return 0
}
