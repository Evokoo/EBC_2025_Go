package quest09

import (
	"sort"
	"strconv"
	"strings"

	"github.com/Evokoo/EBC_2025_Go/utils"
)

// ========================
// I
// ========================
func I(dna []DNA) int {
	childIndex := FindChildIndex(dna)

	if childIndex == -1 {
		return -1
	}

	child := dna[childIndex].sequence
	var likeness []int
	for i, member := range dna {
		if i == childIndex {
			continue
		}

		score := 0
		for j, r := range member.sequence {
			if r == child[j] {
				score++
			}
		}

		likeness = append(likeness, score)
	}

	return likeness[0] * likeness[1]
}

func FindChildIndex(dna []DNA) int {
	sequences := make([][]rune, 0)
	for _, member := range dna {
		sequences = append(sequences, member.sequence)
	}

	common := make(map[int]int)
	for i := range len(sequences[0]) {
		matches := make(map[rune][]int)

		for j, sequence := range sequences {
			r := sequence[i]
			matches[r] = append(matches[r], j)
		}

		for _, indices := range matches {
			if len(indices) > 1 {
				for _, index := range indices {
					common[index]++
				}
			}
		}
	}

	var result [2]int
	for index, count := range common {
		if count > result[1] {
			result[0] = index
			result[1] = count
		}
	}

	if result[1] == len(sequences[0]) {
		return result[0]
	} else {
		return -1
	}
}

// ========================
// II
// ========================

func II(dna []DNA) int {
	seen := make(Set)
	sum := 0

	for a := 0; a < len(dna); a++ {
		for b := a + 1; b < len(dna); b++ {
			for c := b + 1; c < len(dna); c++ {
				key := [3]int{a, b, c}
				sort.Ints(key[:])

				if seen.Has(key) {
					continue
				} else {
					seen.Add(key)
				}

				score := I([]DNA{dna[a], dna[b], dna[c]})
				if score != -1 {
					sum += I([]DNA{dna[a], dna[b], dna[c]})
				}
			}
		}
	}

	return sum
}

// ========================
// SET
// ========================
type Set map[[3]int]struct{}

func (s *Set) Add(key [3]int) {
	(*s)[key] = struct{}{}
}
func (s *Set) Has(key [3]int) bool {
	_, found := (*s)[key]
	return found
}

// ========================
// PARSER
// ========================
type DNA struct {
	id       int
	sequence []rune
}

func ParseInput(file string) []DNA {
	data := utils.ReadFile(file)

	var output []DNA

	for line := range strings.SplitSeq(data, "\n") {
		id, _ := strconv.Atoi(string(line[0]))
		sequence := make([]rune, 0, len(line)-2)

		for _, r := range line[2:] {
			sequence = append(sequence, r)
		}
		output = append(output, DNA{id, sequence})
	}

	return output
}
