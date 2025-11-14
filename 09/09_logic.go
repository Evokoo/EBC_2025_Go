package quest09

import (
	"strconv"
	"strings"

	"github.com/Evokoo/EBC_2025_Go/utils"
)

// ========================
// I
// ========================
func I(dna []DNA) (int, [3]int) {
	childIndex := FindChildIndex(dna)

	if childIndex == -1 {
		return -1, [3]int{}
	}

	child := dna[childIndex].sequence
	product := 1

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
		product *= score
	}

	var family [3]int
	switch childIndex {
	case 0:
		family = [3]int{0, 1, 2}
	case 1:
		family = [3]int{1, 0, 2}
	case 2:
		family = [3]int{2, 0, 1}
	}

	return product, family
}

func FindChildIndex(dna []DNA) int {
	length := len(dna[0].sequence)
	common := [3]int{0, 0, 0}

	for i := range length {
		a := dna[0].sequence[i]
		b := dna[1].sequence[i]
		c := dna[2].sequence[i]

		if a != b && a != c && b != c {
			return -1
		}

		if a == b && b == c {
			common[0]++
			common[1]++
			common[2]++
			continue
		}
		if a == b {
			common[0]++
			common[1]++
		}
		if a == c {
			common[0]++
			common[2]++
		}
		if b == c {
			common[1]++
			common[2]++
		}
	}

	for index, value := range common {
		if value == length {
			return index
		}
	}

	return -1
}

// ========================
// II
// ========================

func II(dna []DNA) int {
	sum := 0

	for a := 0; a < len(dna); a++ {
		for b := a + 1; b < len(dna); b++ {
			for c := b + 1; c < len(dna); c++ {
				score, _ := I([]DNA{dna[a], dna[b], dna[c]})

				if score != -1 {
					sum += score
				}
			}
		}
	}

	return sum
}

// ========================
// III
// ========================

type Node struct {
	id       int
	parents  []*Node
	children []*Node
}
type Tree map[int]*Node

func (t *Tree) AddNode(id int) {
	(*t)[id] = &Node{id: id + 1}
}
func (t *Tree) GetNode(id int) *Node {
	node, found := (*t)[id]

	if !found {
		t.AddNode(id)
		return t.GetNode(id)
	}

	return node
}

func (t *Tree) AddFamily(family [3]int) {
	c := t.GetNode(family[0])
	p1 := t.GetNode(family[1])
	p2 := t.GetNode(family[2])

	c.parents = append(c.parents, p1, p2)
	p1.children = append(p1.children, c)
	p2.children = append(p2.children, c)
}

func III(dna []DNA) int {
	tree := make(Tree)
	children := make(Set)

	for a := 0; a < len(dna); a++ {
		for b := a + 1; b < len(dna); b++ {
			for c := b + 1; c < len(dna); c++ {
				score, indexes := I([]DNA{dna[a], dna[b], dna[c]})

				if score != -1 {
					family := [3]int{0, 0, 0}

					for i, index := range indexes {
						family[i] = []int{a, b, c}[index]
					}

					tree.AddFamily(family)
					children.Add(family[0])
				}
			}
		}
	}

	visited := make(Set)
	result := [2]int{0, 0}

	for i := range children {
		node := tree.GetNode(i)

		if visited.Has(node.id) {
			continue
		} else {
			score, size := 0, 0
			DFS(node, &score, &size, &visited)

			if size > result[0] {
				result[0] = size
				result[1] = score
			}
		}
	}

	return result[1]
}

func DFS(node *Node, score, size *int, visited *Set) {
	if visited.Has(node.id) {
		return
	}

	visited.Add(node.id)
	*score += node.id
	*size++

	for _, parent := range node.parents {
		DFS(parent, score, size, visited)
	}
	for _, child := range node.children {
		DFS(child, score, size, visited)
	}
}

// ========================
// PARSER
// ========================
type DNA struct {
	id       int
	sequence []string
}

func ParseInput(file string) []DNA {
	data := utils.ReadFile(file)

	var output []DNA

	for line := range strings.SplitSeq(data, "\n") {
		sections := strings.Split(line, ":")
		id, _ := strconv.Atoi(sections[0])
		sequence := strings.Split(sections[1], "")
		output = append(output, DNA{id, sequence})
	}

	return output
}

// ========================
// SET
// ========================
type Set map[int]struct{}

func (s *Set) Add(id int) {
	(*s)[id] = struct{}{}
}
func (s *Set) Has(id int) bool {
	_, found := (*s)[id]
	return found
}
