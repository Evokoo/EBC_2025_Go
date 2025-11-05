package quest02

import (
	"strconv"

	"github.com/Evokoo/EBC_2025_Go/utils"
)

// ========================
// PARSER
// ========================
func ParseInput(file string) Pair {
	data := utils.ReadFile(file)

	var output Pair

	for i, value := range utils.QuickMatch(data, `-*\d+`) {
		n, _ := strconv.Atoi(value)
		output[i] = n
	}

	return output
}

// ========================
// OPERATION
// ========================
type Pair [2]int

func NewPair(x, y int) Pair {
	return Pair{x, y}
}

func (p *Pair) Add(a Pair) {
	x1, y1 := p[0], p[1]
	x2, y2 := a[0], a[1]
	(*p)[0] = x1 + x2
	(*p)[1] = y1 + y2
}

func (p *Pair) Multiply(a Pair) {
	x1, y1 := p[0], p[1]
	x2, y2 := a[0], a[1]
	(*p)[0] = x1*x2 - y1*y2
	(*p)[1] = x1*y2 + y1*x2
}

func (p *Pair) Divide(a Pair) {
	x1, y1 := p[0], p[1]
	x2, y2 := a[0], a[1]
	(*p)[0] = x1 / x2
	(*p)[1] = y1 / y2
}

func (p *Pair) InRange(min, max int) bool {
	if p[0] < min || p[0] > max || p[1] < min || p[1] > max {
		return false
	}
	return true
}

// ========================
// PART I
// ========================

func SimpleLoop(file string) Pair {
	input := ParseInput(file)
	result := NewPair(0, 0)
	divisor := NewPair(10, 10)

	for range 3 {
		result.Multiply(result)
		result.Divide(divisor)
		result.Add(input)
	}

	return result
}

// ========================
// PART II & III
// ========================

func CountPoints(file string, size int) Pair {
	a := ParseInput(file)
	b := NewPair(a[0]+1000, a[1]+1000)
	d := NewPair(100000, 100000)

	xStep := (b[0]-a[0])/size + 1
	yStep := (b[1]-a[1])/size + 1
	count := 0

	for y := range size {
	row:
		for x := range size {
			coord := NewPair(a[0]+x*xStep, a[1]+y*yStep)
			result := NewPair(0, 0)

			for range 100 {
				result.Multiply(result)
				result.Divide(d)
				result.Add(coord)

				if !result.InRange(-1000000, 1000000) {
					continue row
				}
			}
			count++
		}
	}
	return NewPair(count, 0)
}
