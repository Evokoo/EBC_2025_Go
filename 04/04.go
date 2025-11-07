package quest04

func Solve(file string, part int) int {
	switch part {
	case 1:
		return int(I(ParseGears(file), 2025))
	case 2:
		return int(II(ParseGears(file), 10000000000000))
	case 3:
		return int(I(ParseGears(file), 100))
	}

	panic("Invaid Quest part")
}
