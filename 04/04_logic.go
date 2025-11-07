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
	turns := 1.0
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

func II(gears [][]float64, target float64) float64 {
	turns := I(gears, 1)
	return math.Ceil(target / turns)
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
			n, _ := strconv.ParseFloat(gear, 64)
			set = append(set, n)
		}

		sets = append(sets, set)
	}

	return sets
}
