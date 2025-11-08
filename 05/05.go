package quest05

func Solve(file string, part int) int {
	lists := ParseInput(file)

	switch part {
	case 1:
		quality, _ := I(lists[0])
		return quality
	case 2:
		return II(lists)
	case 3:
		return III(lists)
	}

	panic("Invalid Quest part")
}
