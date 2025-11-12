package quest07

func Solve(file string, part int) any {
	names, rules := ParseInput(file)

	switch part {
	case 1:
		return I(names, rules, part)
	case 2:
		return I(names, rules, part)
	case 3:
		III(names, rules)
	}

	return ""
}
