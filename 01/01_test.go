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
	target int
}

var tests = []Test{
	{part: 1, file: "example_I.txt", target: 5},
	{part: 1, file: "input_I.txt", target: 1306},
	{part: 2, file: "example_II.txt", target: 28},
	{part: 2, file: "input_II.txt", target: 5636},
	{part: 3, file: "example_III.txt", target: 30},
	{part: 3, file: "input_III.txt", target: 27983},
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
