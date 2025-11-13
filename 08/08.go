package quest08

func Solve(file string, part int) int {
	list, pairs, nails := ParseInput(file)

	switch part {
	case 1:
		return I(list, pairs)
	case 2:
		return II(list)
	case 3:
		return III(list, nails)
	}
	return 0
}
