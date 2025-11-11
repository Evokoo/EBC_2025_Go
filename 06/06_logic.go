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

	for novice, indices := range locations.novice {
		if novice != 'a' && filter {
			continue
		}

		for i := range indices {
			for j := range locations.mentor[unicode.ToUpper(novice)] {
				if i > j {
					count++
				}
			}
		}
	}

	return count
}

// ========================
// PART III
// ========================

func III(locations LocationMap, distance, repeat, length int) int {
	count := 0
	total := length * repeat

	for novice, nSet := range locations.novice {
		mSet := locations.mentor[unicode.ToUpper(novice)]

		for n := range nSet {
			for offset := -distance; offset <= distance; offset++ {
				iAbs := n + offset

				first := (0 - iAbs + length - 1) / length
				last := (total - 1 - iAbs) / length

				if first < 0 {
					first = 0
				}

				if last >= repeat {
					last = repeat - 1
				}

				if last >= first {
					if mSet.Has(((iAbs % length) + length) % length) {
						count += last - first + 1
					}
				}
			}
		}
	}

	return count
}

// ========================
// PARSER
// ========================
type Set map[int]struct{}

func (s Set) Add(value int) {
	s[value] = struct{}{}
}
func (s *Set) Has(value int) bool {
	_, found := (*s)[value]
	return found
}

type LocationMap struct {
	novice map[rune]Set
	mentor map[rune]Set
}

func ParseInput(file string) (LocationMap, int) {
	data := utils.ReadFile(file)
	len := len(data)
	locations := LocationMap{
		mentor: make(map[rune]Set),
		novice: make(map[rune]Set),
	}

	for i, r := range data {
		if unicode.IsUpper(r) {
			_, found := locations.mentor[r]
			if !found {
				locations.mentor[r] = make(Set)
			}

			locations.mentor[r].Add(i)
		} else {
			_, found := locations.novice[r]
			if !found {
				locations.novice[r] = make(Set)
			}

			locations.novice[r].Add(i)
		}
	}
	return locations, len
}
