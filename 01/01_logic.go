package quest01

import "github.com/Evokoo/EBC_2025_Go/utils"

// ========================
// COST
// ========================

var Units = map[rune]int{
	'A': 0, 'B': 1, 'C': 3, 'D': 5,
}

func BattleCost(battles []string) int {
	total := 0

	for _, battle := range battles {
		monsters := 0
		potions := 0

		for _, monster := range battle {
			if monster != 'x' {
				potions += Units[monster]
				monsters++
			}
		}

		switch monsters {
		case 2:
			potions += 2
		case 3:
			potions += 6
		}

		total += potions
	}

	return total
}

// ========================
// PARSER
// ========================

func ParseInput(file string, part int) []string {
	data := utils.ReadFile(file)
	output := make([]string, 0)

	for i := 0; i < len(data); i += part {
		battle := data[i : i+part]
		output = append(output, battle)
	}

	return output
}
