package quest12

func Solve(file string, part int) int {
	barrels := ParseInput(file)

	a := [2]int{0, 0}
	b := [2]int{barrels.cols - 1, barrels.rows - 1}

	switch part {
	case 1:
		destoryed := I(barrels, make(Set[[2]int]), [][2]int{a})
		return len(destoryed)
	case 2:
		destoryed := I(barrels, make(Set[[2]int]), [][2]int{a, b})
		return len(destoryed)
	case 3:
		III(barrels)
	}
	return 0
}
