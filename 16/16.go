package quest16

func Solve(file string, part int) int {
	input := ParseInput(file)

	switch part {
	case 1:
		return I(input, 90)
	case 2:
		score, _ := II(input)
		return score
	case 3:
		return III(input, 202520252025000)
	}
	return 0
}
