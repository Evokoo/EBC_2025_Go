package quest09

func Solve(file string, part int) int {
	dna := ParseInput(file)

	switch part {
	case 1:
		score, _ := I(dna)
		return score
	case 2:
		return II(dna)
	case 3:
		return III(dna)
	}

	return 0
}
