package quest07

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/Evokoo/EBC_2025_Go/utils"
)

// ========================
// PART I
// ========================
func I(names []string, tests Tests) string {
	for _, name := range names {
		if IsValidName(name, tests) {
			return name
		}
	}
	panic("Name not found")
}

// ========================
// PART II
// ========================
func II(names []string, tests Tests) (score int) {
	for i, name := range names {
		if IsValidName(name, tests) {
			score += (i + 1)
		}
	}
	return
}

// ========================
// PART III
// ========================
func III(names []string, rules Rules, tests Tests) int {
	list := make(Set)
	for _, name := range names {
		if IsValidName(name, tests) {
			DFS(name, rules, &list)
		}
	}
	return len(list)
}

func DFS(name string, rules Rules, set *Set) {
	if len(name) >= 7 && len(name) <= 11 {
		set.Add(name)
	}
	if len(name) >= 11 {
		return
	}

	key := string(name[len(name)-1])
	for _, char := range rules[key] {
		DFS(name+char, rules, set)
	}
}

// ========================
// VALIDATOR
// ========================
func IsValidName(name string, tests Tests) bool {
pairs:
	for i := 2; i <= len(name); i++ {
		pair := name[i-2 : i]
		for _, test := range tests {
			if isMatch := test.MatchString(pair); isMatch {
				continue pairs
			}
		}

		return false
	}
	return true
}

// ========================
// PARSER
// ========================

type Rules map[string][]string
type Tests []*regexp.Regexp

func ParseInput(file string) ([]string, Rules, Tests) {
	data := utils.ReadFile(file)
	names := make([]string, 0)
	rules := make(Rules)

	for i, line := range strings.Split(data, "\n") {
		if i == 0 {
			names = strings.Split(line, ",")
		}
		if i > 1 {
			sections := strings.Split(line, " > ")
			characters := strings.Split(sections[1], ",")
			rules[sections[0]] = characters
		}
	}
	return names, rules, ConvertRulesToTests(rules)
}
func ConvertRulesToTests(rules Rules) Tests {
	tests := make(Tests, 0, len(rules))

	for key, chars := range rules {
		var selection strings.Builder
		for _, char := range chars {
			selection.WriteString(char)
		}

		pattern := fmt.Sprintf(`%v[%v]`, key, selection.String())
		test := regexp.MustCompile(pattern)

		tests = append(tests, test)
	}

	return tests
}

// ========================
// SET
// ========================
type Set map[string]struct{}

func (s *Set) Add(name string) {
	(*s)[name] = struct{}{}
}
