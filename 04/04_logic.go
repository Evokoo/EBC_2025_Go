package quest04

import (
	"math"
	"strconv"
	"strings"

	"github.com/Evokoo/EBC_2025_Go/utils"
)

// ========================
// PART I & III
// ========================

func I(gearSets [][]float64, rotations float64) float64 {
	turns := float64(1)
	last := gearSets[0][0]

	for _, set := range gearSets[1:] {
		turns *= last / set[0]

		if len(set) > 1 {
			last = set[1]
		} else {
			last = set[0]
		}
	}
	return turns * rotations
}

// ========================
// PART II
// ========================

func II(gears [][]float64, target int) float64 {
	turns := I(gears, 1)
	return math.Ceil(float64(target) / turns)
}

// ========================
// PARSER
// ========================

func ParseGears(file string) [][]float64 {
	data := utils.ReadFile(file)
	sets := make([][]float64, 0)

	for line := range strings.SplitSeq(data, "\n") {
		var set []float64

		for _, gear := range strings.Split(line, "|") {
			n, _ := strconv.Atoi(gear)
			set = append(set, float64(n))
		}

		sets = append(sets, set)
	}

	return sets
}
