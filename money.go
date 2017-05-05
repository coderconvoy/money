package money

import (
	"errors"
	"fmt"
)

type M int

func Parse(st string) (M, error) {
	dotted := -1
	res := M(0)
	for _, s := range st {
		if dotted >= 0 {
			dotted++
			if dotted > 3 {
				return 0, errors.New("dp too small")
			}
		}
		switch s {
		case '.':
			if dotted >= 0 {
				return 0, errors.New("Too many dots")
			}
			dotted = 1
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			n := M(s - '0')
			res = res*10 + n
		default:
			return 0, errors.New("Unexpected Char: " + string(s))
		}

	}
	if dotted == -1 {
		return res * 100, nil
	}
	for dotted < 3 {
		res *= 10
		dotted++
	}
	return res, nil

}

func SParse(s string, def int) M {
	r, err := Parse(s)
	if err != nil {
		return M(def)
	}
	return r
}

func (m M) String() string {
	if m < 0 {
		return "-£" + fmt.Sprintf("%.2f", float32(-m)/100)
	}
	return "£" + fmt.Sprintf("%.2f", float32(m)/100)
}
