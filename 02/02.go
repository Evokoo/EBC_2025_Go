package quest02

func Solve(file string, part int) Pair {
	switch part {
	case 1:
		return SimpleLoop(file)
	case 2, 3:
		size := 101

		if part == 3 {
			size = 1001
		}

		return PlotPoints(file, size)
	}

	return NewPair(-1, -1)
}
