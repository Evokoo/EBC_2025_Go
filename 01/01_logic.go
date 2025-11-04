package quest01

import (
	"strconv"
	"strings"

	"github.com/Evokoo/EBC_2025_Go/utils"
)

// ========================
// PARSER
// ========================

type Names []string
type Moves []int

func ParseInput(file string) (Names, Moves) {
	data := utils.ReadFile(file)

	var names Names
	var directions Moves

	for i, line := range strings.Split(data, "\n") {
		if i == 0 {
			names = strings.Split(line, ",")
		}
		if i == 2 {
			for direction := range strings.SplitSeq(line, ",") {
				value, _ := strconv.Atoi(utils.QuickMatch(direction, `\d+`)[0])
				switch direction[0] {
				case 'R':
					directions = append(directions, value)
				case 'L':
					directions = append(directions, -value)
				}
			}
		}
	}
	return names, directions
}

func FindName(file string, part int) string {
	list, moves := ParseInput(file)
	length := len(list)
	index := 0

	for _, move := range moves {
		index += move

		switch part {
		case 1:
			if index < 0 {
				index = 0
			}
			if index >= length {
				index = length - 1
			}
		case 2:
			index = (index + length) % length
		case 3:
			target := ((move % length) + length) % length
			list[0], list[target] = list[target], list[0]
			index = 0
		}
	}

	return list[index]
}
