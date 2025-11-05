package quest02_test

import (
	"fmt"

	. "github.com/Evokoo/EBC_2025_Go/02"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type Test struct {
	file   string
	part   int
	target Pair
}

var tests = []Test{
	{part: 1, file: "inputs/example_I.txt", target: Pair{357, 862}},
	{part: 1, file: "inputs/input_I.txt", target: Pair{198405, 991528}},
	{part: 2, file: "inputs/example_II.txt", target: Pair{4076, 0}},
	{part: 2, file: "inputs/input_II.txt", target: Pair{583, 0}},
	{part: 3, file: "inputs/example_III.txt", target: Pair{406954, 0}},
	{part: 3, file: "inputs/input_III.txt", target: Pair{54687, 0}},
}

var _ = Describe("EBC 2025 - Quest 02", func() {
	for _, test := range tests {
		msg := fmt.Sprintf("Testing Part %d with %s", test.part, test.file)
		It(msg, func() {
			result := Solve(test.file, test.part)
			Expect(result).To(Equal((test.target)))
		})
	}
})
