package regexp

type Eof struct {
	common
}

func (exp *Eof) Match(str string, offset int) int {
	return offset
}
