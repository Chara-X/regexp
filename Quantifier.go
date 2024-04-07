package regexp

type Quantifier struct {
	Branch Regexp
	common
}

func (exp *Quantifier) Match(str string, offset int) int {
	var length = exp.Branch.Match(str, offset)
	if length != 0 {
		return length
	} else {
		return exp.next.Match(str, offset)
	}
}
