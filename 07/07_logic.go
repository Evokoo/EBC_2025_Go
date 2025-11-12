package quest07

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/Evokoo/EBC_2025_Go/utils"
)

// ========================
// PART I & II
// ========================

func I(names []string, rules []*regexp.Regexp, part int) any {
	var score int
	var validNames []string

list:
	for i, name := range names {
	pairs:
		for i := 2; i <= len(name); i++ {
			pair := name[i-2 : i]

			for _, rule := range rules {
				if isMatch := rule.MatchString(pair); isMatch {
					continue pairs
				}
			}

			continue list
		}
		if part == 1 {
			return name
		}
		if part == 2 {
			score += (i + 1)
		}
		if part == 3 {
			validNames = append(validNames, name)
		}
	}

	if part == 3 {
		return validNames
	}

	return score
}

// ========================
// PART III
// ========================

func III(names []string, rules []*regexp.Regexp) {
	suffixes := make(map[string][]string)
	for _, rule := range rules {
		chars := utils.QuickMatch(rule.String(), `[a-z|A-Z]`)

		fmt.Println(rule, chars)

		key := chars[0]
		suffixes[key] = chars[1:]
	}

	count := 0

	for _, name := range I(names, rules, 3).([]string)[:1] {
		validNames := make(Set)
		DFS(name, suffixes, &validNames)
		count += len(validNames)

		fmt.Println(name)
		for n := range validNames {
			fmt.Println(n)
		}
	}

	fmt.Println(count)
}

func DFS(name string, suffixes map[string][]string, valid *Set) {
	if len(name) >= 7 && len(name) <= 11 {
		valid.Add(name)
	}
	if len(name) >= 11 {
		return
	}

	key := string(name[len(name)-1])
	options := suffixes[key]
	for _, option := range options {
		DFS(name+option, suffixes, valid)
	}
}

// ========================
// SET
// ========================
type Set map[string]struct{}

func (s Set) Add(name string) {
	s[name] = struct{}{}
}
func (s Set) Has(name string) bool {
	_, found := s[name]
	return found
}

// ========================
// PARSER
// ========================

func ParseInput(file string) ([]string, []*regexp.Regexp) {
	data := utils.ReadFile(file)

	var names []string
	var rules []*regexp.Regexp

	for i, line := range strings.Split(data, "\n") {
		if i == 0 {
			names = strings.Split(line, ",")
		}

		if i > 1 {
			characters := utils.QuickMatch(line, `\w`)

			var selection strings.Builder
			for _, char := range characters[1:] {
				selection.WriteString(char)
			}

			pattern := fmt.Sprintf(`%v[%v]`, characters[0], selection.String())
			re := regexp.MustCompile(pattern)

			rules = append(rules, re)
		}
	}

	return names, rules
}
