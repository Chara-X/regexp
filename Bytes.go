package regexp

import (
	"slices"
)

type Bytes struct {
	Ranges []entry
	common
}

func (exp *Bytes) Match(str string, offset int) int {
	if offset == len(str) || !slices.ContainsFunc(exp.Ranges, func(entry entry) bool {
		return str[offset] >= entry.Low && str[offset] <= entry.High
	}) {
		return 0
	} else {
		return exp.next.Match(str, offset+1)
	}
}

type entry struct {
	Low  byte
	High byte
}
