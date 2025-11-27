package quest18

import (
	"strconv"
	"strings"

	"github.com/Evokoo/EBC_2025_Go/utils"
)

// ========================
// PLANT, CHILD & TREE
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
// ENERGY CALCULATOR
// ========================

func Calculate(tree Tree, configuration []int) int {
	queue, store := QueueAndStore(tree, configuration)
	targetID := len(tree)
	parentsProcessed := make(map[int]int)

	for !queue.IsEmpty() {
		current := queue.Pop()
		incoming := store[current.id]

		energy := 0
		if incoming >= current.weight {
			energy = incoming
		}

		for _, childBranch := range current.children {
			child := childBranch.plant
			outgoing := energy * childBranch.weight
			store[child.id] += outgoing

			parentsProcessed[child.id]++
			if parentsProcessed[child.id] == len(child.parents) {
				queue.Push(child)
			}
		}
	}

	result := store[targetID]
	if result < tree.GetPlant(targetID).weight {
		return 0
	}

	return result
}
func QueueAndStore(tree Tree, configuration []int) (Queue[*Plant], map[int]int) {
	queue := make(Queue[*Plant], 0)
	store := make(map[int]int)

	for _, plant := range tree {
		if len(plant.parents) == 0 {
			queue = append(queue, plant)
			store[plant.id] = configuration[plant.id-1]
		}
	}

	return queue, store
}

// ========================
// PART I
// ========================

func I(tree Tree) int {
	configuration := make([]int, 0, len(tree))

	for _, plant := range tree {
		if len(plant.parents) == 0 {
			configuration = append(configuration, 1)
		}
	}

	return Calculate(tree, configuration)
}

// ========================
// PART II
// ========================

func II(tree Tree, configurations [][]int) int {
	total := 0

	for _, configuration := range configurations {
		total += Calculate(tree, configuration)
	}

	return total
}

// ========================
// PART III
// ========================

func III(tree Tree, configurations [][]int) int {
	optimal := Calculate(tree, FindOptimalConfiguration(tree, len(configurations[0])))

	totalDifference := 0
	for _, config := range configurations {
		result := Calculate(tree, config)

		if result > 0 {
			totalDifference += optimal - result
		}
	}
	return totalDifference
}

func FindOptimalConfiguration(tree Tree, length int) []int {
	optimalConfiguration := make([]int, length)

	for id, plant := range tree {
		if len(plant.parents) == 0 {
			flipBit := true

			for _, child := range plant.children {
				if child.weight < 0 {
					flipBit = false
					break
				}
			}

			if flipBit {
				optimalConfiguration[id-1] = 1
			}
		}
	}

	improved := true
	for improved {

		improved = false
		currentResult := Calculate(tree, optimalConfiguration)

		for i := range len(optimalConfiguration) {
			newConfig := make([]int, len(optimalConfiguration))
			copy(newConfig, optimalConfiguration)

			newConfig[i] = 1 - newConfig[i]
			newResult := Calculate(tree, newConfig)

			if newResult > currentResult {
				optimalConfiguration[i] = newConfig[i]
				improved = true
				currentResult = newResult
			}
		}
	}

	return optimalConfiguration
}

// ========================
// PARSER
// ========================

func ParseInput(file string) (Tree, [][]int) {
	data := utils.ReadFile(file)
	sections := strings.Split(data, "\n\n")

	tree := make(Tree, len(sections))
	connections := [][3]int{}
	activity := make([][]int, 0)

	for _, section := range sections {
		var id, weight int

		if !strings.HasPrefix(section, "P") {
			for _, line := range strings.Split(section, "\n") {
				row := make([]int, 0)

				for _, value := range strings.Split(line, " ") {
					n, _ := strconv.Atoi(value)
					row = append(row, n)
				}
				activity = append(activity, row)
			}
			continue
		}

		for i, line := range strings.Split(section, "\n") {
			values := utils.QuickMatch(line, `-*\d+`)

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

	return tree, activity
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
