package quest05

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/Evokoo/EBC_2025_Go/utils"
)

// ========================
// LAYER
// ========================
type Layer [3]int

func NewLayer(value int) Layer {
	return Layer{-1, value, -1}
}
func (l *Layer) GetQuality() int {
	var digits strings.Builder

	for _, value := range *l {
		if value != -1 {
			digits.WriteString(fmt.Sprintf("%d", value))
		}
	}

	result, _ := strconv.Atoi(digits.String())
	return result
}

// ========================
// FISHBONE
// ========================
type FishBone []Layer

func NewFishBone(value int) FishBone {
	return FishBone{NewLayer(value)}
}
func (s *FishBone) AddLayer(value int) {
	(*s) = append((*s), NewLayer(value))
}
func (s *FishBone) AddValue(value int) {
	for i, layer := range *s {
		if value < layer[1] && layer[0] == -1 {
			(*s)[i][0] = value
			return
		}
		if value > layer[1] && layer[2] == -1 {
			(*s)[i][2] = value
			return
		}
	}
	(*s).AddLayer(value)
}
func (s *FishBone) GetQuality() int {
	var digits strings.Builder

	for _, layer := range *s {
		digits.WriteString(fmt.Sprintf("%d", layer[1]))
	}

	result, _ := strconv.Atoi(digits.String())
	return result
}

// ========================
// PART I
// ========================

func I(list []int) (int, FishBone) {
	fishBone := NewFishBone(list[1])

	for _, value := range list[2:] {
		fishBone.AddValue(value)
	}

	return fishBone.GetQuality(), fishBone
}

// ========================
// PART II
// ========================

func II(lists [][]int) int {
	best := math.MinInt
	worst := math.MaxInt

	for _, list := range lists {
		quality, _ := I(list)
		best = max(quality, best)
		worst = min(quality, worst)
	}

	return best - worst
}

// ========================
// PART III
// ========================

type Sword struct {
	id        int
	quality   int
	structure FishBone
}

func (s Sword) LesserThan(b Sword) bool {
	if s.quality != b.quality {
		return s.quality < b.quality
	}

	for i := range max(len(b.structure), len(s.structure)) {
		var q1, q2 int

		if i < len(s.structure) {
			q1 = s.structure[i].GetQuality()
		}
		if i < len(b.structure) {
			q2 = b.structure[i].GetQuality()
		}

		if q1 != q2 {
			return q1 < q2
		}
	}

	return s.id < b.id
}

func III(lists [][]int) int {
	var swords []Sword

	for _, list := range lists {
		quality, structure := I(list)
		swords = append(swords, Sword{list[0], quality, structure})
	}

	sort.Slice(swords, func(i, j int) bool {
		return swords[j].LesserThan(swords[i])
	})

	checksum := 0

	for i, sword := range swords {
		checksum += (i + 1) * sword.id
	}

	return checksum
}

// ========================
// PARSER
// ========================

func ParseInput(file string) [][]int {
	data := utils.ReadFile(file)
	output := make([][]int, 0)

	for line := range strings.SplitSeq(data, "\n") {
		row := make([]int, 0)
		for _, value := range utils.QuickMatch(line, `-*\d+`) {
			n, _ := strconv.Atoi(value)
			row = append(row, n)
		}
		output = append(output, row)
	}

	return output
}
