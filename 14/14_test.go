package quest14_test

import (
	"fmt"

	. "github.com/Evokoo/EBC_2025_Go/14"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type Test struct {
	file   string
	part   int
	target int
}

var tests = []Test{
	// {part: 1, file: "inputs/example_I.txt", target: 200},
	// {part: 1, file: "inputs/input_I.txt", target: 496},
	// // {part: 2, file: "inputs/example_II.txt", target: -1},
	// {part: 2, file: "inputs/input_II.txt", target: 1169090},
	// {part: 3, file: "inputs/example_III.txt", target: 278388552},
	{part: 3, file: "inputs/input_III.txt", target: 1263979820},
}

var _ = Describe("EBC 2025 - Quest 14", func() {
	for _, test := range tests {
		msg := fmt.Sprintf("Testing Part %d with %s", test.part, test.file)
		It(msg, func() {
			result := Solve(test.file, test.part)
			Expect(result).To(Equal((test.target)))
		})
	}
})
