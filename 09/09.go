package quest09

func Solve(file string, part int) int {
	dna := ParseInput(file)

	switch part {
	case 1:
		return I(dna)
	case 2:
		return II(dna)
	}

	return 0
}
