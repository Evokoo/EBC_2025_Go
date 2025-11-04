package quest01

func Solve(file string, part int) string {
	// names, directions := ParseInput(file)
	// length := len(names)

	// i := 0
	// for _, step := range directions {
	// 	i += step

	// 	if i < 0 {
	// 		i = 0
	// 	}

	// 	if i >= length {
	// 		i = length - 1
	// 	}
	// }

	// fmt.Println(i)

	return FindName(file, part)
}
