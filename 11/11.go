package quest11

func Solve(file string, part int) int {
	columns := ParseInput(file)

	switch part {
	case 1:
		return I(columns, 10)
	case 2:
		return II(columns)
	case 3:
		return III(columns)
	}
	return 0
}
