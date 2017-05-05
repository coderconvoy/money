package money

import "testing"

func Test_Amount(t *testing.T) {
	tr := []struct {
		In  string
		Out M
		Err bool
	}{
		{"helo", 0, true},
		{"0.4", 40, false},
		{"-1.7", 0, true},
		{"1.54-", 0, true},
		{"1.15", 115, false},
		{"1.453", 0, true},
		{"134", 13400, false},
		{"135.3", 13530, false},
	}

	for _, ti := range tr {
		o, err := Parse(ti.In)
		if o != ti.Out || ((err == nil) == ti.Err) {
			t.Logf("%s,expected :%d,%s. Got %d,%s", ti.In, ti.Out, ti.Err, o, err)
			t.Fail()
		}

	}
}
