package quest01_test

import (
	"fmt"

	. "github.com/Evokoo/EBC_2025_Go/01"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type Test struct {
	file   string
	part   int
	target string
}

var tests = []Test{
	{part: 1, file: "inputs/example_I.txt", target: "Fyrryn"},
	{part: 1, file: "inputs/input_I.txt", target: "Draithacris"},
	{part: 2, file: "inputs/example_II.txt", target: "Elarzris"},
	{part: 2, file: "inputs/input_II.txt", target: "Orysaral"},
	{part: 3, file: "inputs/example_III.txt", target: "Drakzyph"},
	{part: 3, file: "inputs/input_III.txt", target: "Bryntal"},
}

var _ = Describe("EBC 2025 - Quest 01", func() {
	for _, test := range tests {
		msg := fmt.Sprintf("Testing Part %d with %s", test.part, test.file)
		It(msg, func() {
			result := Solve(test.file, test.part)
			Expect(result).To(Equal((test.target)))
		})
	}
})
