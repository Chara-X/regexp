package regexp

type Byte struct {
	Byte byte
	common
}

func (exp *Byte) Match(str string, offset int) int {
	if offset == len(str) || str[offset] != exp.Byte {
		return 0
	} else {
		return exp.next.Match(str, offset+1)
	}
}
