package quest07

func Solve(file string, part int) any {
	names, rules, tests := ParseInput(file)

	switch part {
	case 1:
		return I(names, tests)
	case 2:
		return II(names, tests)
	case 3:
		return III(names, rules, tests)
	}
	return ""
}
