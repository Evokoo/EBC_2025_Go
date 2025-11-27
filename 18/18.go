package quest18

func Solve(file string, part int) int {
	tree, activity := ParseInput(file)

	switch part {
	case 1:
		return I(tree)
	case 2:
		return II(tree, activity)
	case 3:
		return III(tree, activity)
	}
	return 0
}
