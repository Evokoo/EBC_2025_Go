package quest04_test

import (
	"fmt"

	. "github.com/Evokoo/EBC_2025_Go/04"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type Test struct {
	file   string
	part   int
	target int
}

var tests = []Test{
	{part: 1, file: "inputs/example_I.txt", target: 15888},
	{part: 1, file: "inputs/input_I.txt", target: 12898},
	{part: 2, file: "inputs/example_II.txt", target: 1274509803922},
	{part: 2, file: "inputs/input_II.txt", target: 1661237785017},
	{part: 3, file: "inputs/example_III.txt", target: 6818},
	{part: 3, file: "inputs/input_III.txt", target: 972550704505},
}

var _ = Describe("EBC 2025 - Quest 04", func() {
	for _, test := range tests {
		msg := fmt.Sprintf("Testing Part %d with %s", test.part, test.file)
		It(msg, func() {
			result := Solve(test.file, test.part)
			Expect(result).To(Equal((test.target)))
		})
	}
})
