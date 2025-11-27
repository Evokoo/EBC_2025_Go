package quest18

func Solve(file string, part int) int {
	tree := ParseInput(file)

	switch part {
	case 1:
		return I(tree)
	}
	return 0
}
