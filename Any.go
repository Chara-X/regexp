package regexp

type Any struct {
	common
}

func (exp *Any) Match(str string, offset int) int {
	if offset == len(str) {
		return 0
	} else {
		return exp.next.Match(str, offset+1)
	}
}
