package quest16_test

import (
	"fmt"

	. "github.com/Evokoo/EBC_2025_Go/16"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type Test struct {
	file   string
	part   int
	target int
}

var tests = []Test{
	{part: 1, file: "inputs/example_I.txt", target: 193},
	{part: 1, file: "inputs/input_I.txt", target: 226},
	{part: 2, file: "inputs/example_II.txt", target: 270},
	{part: 2, file: "inputs/input_II.txt", target: 118764112896},
	{part: 3, file: "inputs/example_III.txt", target: 94439495762954},
	{part: 3, file: "inputs/input_III.txt", target: 95628513597567},
}

var _ = Describe("EBC 2025 - Quest 16", func() {
	for _, test := range tests {
		msg := fmt.Sprintf("Testing Part %d with %s", test.part, test.file)
		It(msg, func() {
			result := Solve(test.file, test.part)
			Expect(result).To(Equal((test.target)))
		})
	}
})
