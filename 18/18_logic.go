package quest18

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Evokoo/EBC_2025_Go/utils"
)

// ========================
// PLANT, BRANCH & TREE
// ========================
type Plant struct {
	id, weight int
	children   []Child
	parents    []*Plant
}
type Child struct {
	plant  *Plant
	weight int
}

func NewPlant(id, weight int) *Plant {
	return &Plant{
		id:       id,
		weight:   weight,
		children: make([]Child, 0), // Weight of connection and branch it goes to
		parents:  make([]*Plant, 0),
	}
}
func (p *Plant) AddChild(plant *Plant, weight int) {
	(*p).children = append((*p).children, Child{plant, weight})
}
func (p *Plant) AddParent(plant *Plant) {
	(*p).parents = append((*p).parents, plant)
}

type Tree map[int]*Plant

func (t *Tree) AddPlant(plant *Plant) {
	(*t)[plant.id] = plant
}
func (t Tree) GetPlant(id int) *Plant {
	if plant, found := t[id]; found {
		return plant
	}

	panic("No plant found!")
}
func (t Tree) ConnectPlants(a, b, weight int) {
	child := t.GetPlant(a)
	parent := t.GetPlant(b)

	parent.AddChild(child, weight)
	child.AddParent(parent)
}

// ========================
// PART I
// ========================

func I(tree Tree) int {
	queue := make([]*Plant, 0)
	store := make(map[int]int)        // total incoming energy for each plant
	readyParents := make(map[int]int) // number of parents processed per plant

	// Initialize entry nodes
	for _, plant := range tree {
		if len(plant.parents) == 0 {
			queue = append(queue, plant)
			store[plant.id] = 1 // free branch energy
		}
	}

	// BFS / topological processing
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		incoming := store[current.id]

		// Threshold: if energy < thickness, plant dies, propagate 0
		var energy int
		if incoming < current.weight {
			energy = 0
		} else {
			energy = incoming
		}

		// Propagate to children
		for _, childBranch := range current.children {
			child := childBranch.plant
			outgoing := energy * childBranch.weight
			store[child.id] += outgoing

			// track processed parents
			readyParents[child.id]++
			if readyParents[child.id] == len(child.parents) {
				queue = append(queue, child)
			}
		}
	}

	fmt.Println("Energy per plant:", store)

	// If there is a single final plant (no children), return its energy
	for _, plant := range tree {
		if len(plant.children) == 0 {
			return store[plant.id]
		}
	}

	return 0
}

// ========================
// PARSER
// ========================

func ParseInput(file string) Tree {
	data := utils.ReadFile(file)
	sections := strings.Split(data, "\n\n")

	tree := make(Tree, len(sections))
	connections := [][3]int{}

	for _, section := range sections {
		var id, weight int

		for i, line := range strings.Split(section, "\n") {
			values := utils.QuickMatch(line, `\d+`)

			if i == 0 {
				a, _ := strconv.Atoi(values[0])
				b, _ := strconv.Atoi(values[1])
				id, weight = a, b
			} else {
				if !strings.Contains(line, "free") {
					a, _ := strconv.Atoi(values[0])
					b, _ := strconv.Atoi(values[1])
					connections = append(connections, [3]int{id, a, b})
				}
			}
		}

		tree.AddPlant(NewPlant(id, weight))
	}

	for _, connection := range connections {
		a, b, weight := connection[0], connection[1], connection[2]
		tree.ConnectPlants(a, b, weight)
	}

	return tree
}

// ========================
// SET
// ========================
type Set[T comparable] map[T]struct{}

func (s Set[T]) Add(v T) {
	s[v] = struct{}{}
}
func (s Set[T]) Has(v T) bool {
	_, ok := s[v]
	return ok
}

// ========================
// QUEUE
// ========================
type Queue[T comparable] []T

func (q *Queue[T]) Pop() T {
	popped := (*q)[0]
	(*q) = (*q)[1:]
	return popped
}
func (q *Queue[T]) Push(value T) {
	(*q) = append((*q), value)
}
func (q *Queue[T]) IsEmpty() bool {
	return len(*q) == 0
}
