package quest06

import (
	"unicode"

	"github.com/Evokoo/EBC_2025_Go/utils"
)

// ========================
// PART I & II
// ========================

func I(locations LocationMap, filter bool) int {
	count := 0

	for student, indices := range locations.students {
		if student != 'a' && filter {
			continue
		}

		mentors := locations.mentors[unicode.ToUpper(student)]
		for _, i := range indices {
			for _, j := range mentors {
				if i > j {
					count++
				}
			}
		}
	}

	return count
}

// ========================
// PARSER
// ========================
type LocationMap struct {
	mentors  map[rune][]int
	students map[rune][]int
}

func ParseInput(file string) LocationMap {
	data := utils.ReadFile(file)
	locations := LocationMap{
		mentors:  make(map[rune][]int, 0),
		students: make(map[rune][]int, 0),
	}

	for i, r := range data {
		if unicode.IsUpper(r) {
			current := locations.mentors[r]
			current = append(current, i)
			locations.mentors[r] = current
		} else {
			current := locations.students[r]
			current = append(current, i)
			locations.students[r] = current
		}
	}

	return locations
}
