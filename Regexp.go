package regexp

import (
	"fmt"
	"strings"
)

type Regexp interface {
	SetNext(next Regexp)
	Match(str string, offset int) int
}

func New(pattern string) Regexp {
	return new(strings.NewReader(pattern))
}
func new(reader *strings.Reader) Regexp {
	var next, err = reader.ReadByte()
	fmt.Printf("next: %c\n", next)
	if err == nil {
		switch next {
		case '{':
			var branch = new(reader)
			reader.ReadByte()
			var exp = &Quantifier{Branch: branch, common: common{next: new(reader)}}
			branch.SetNext(exp)
			return exp
		case '(':
			var next = new(reader)
			reader.ReadByte()
			var exp = &Quantifier{Branch: new(reader), common: common{next: next}}
			next.SetNext(exp)
			return exp
		case '[':
			var ranges = []entry{}
			for {
				var next, _ = reader.ReadByte()
				if next != ']' {
					reader.UnreadByte()
					var low, _ = reader.ReadByte()
					var high, _ = reader.ReadByte()
					ranges = append(ranges, entry{Low: low, High: high})
				} else {
					break
				}
			}
			return &Bytes{Ranges: ranges, common: common{next: new(reader)}}
		case '.':
			return &Any{common: common{next: new(reader)}}
		case '}':
			return &Eof{}
		case ')':
			return &Eof{}
		default:
			if next == '\\' {
				next, _ = reader.ReadByte()
			}
			return &Byte{Byte: next, common: common{next: new(reader)}}
		}
	} else {
		return &Eof{}
	}
}

type common struct {
	next Regexp
}

func (c *common) SetNext(next Regexp) {
	c.next = next
}
