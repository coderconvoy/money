package money

import (
	"errors"
	"strings"
)

// CM is Currency Money
type CM struct {
	mon M
	cur string
}

func ParseCM(s string) (CM, error) {
	s = strings.TrimSpace(s)
	if len(s) < 4 {
		return CM{}, errors.New("Could not load parse CM, string too short ")
	}
	m := s[:len(s)-3]
	c := s[len(s)-3:]
	mon, err := Parse(m)

	return CM{mon, c}, err
}

func (cm CM) String() string {
	return cm.mon.String() + cm.cur
}

func (cm CM) MarshalJSON() ([]byte, error) {
	return []byte(`"` + cm.String() + `"`), nil
}
func (cm *CM) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)
	r, err := ParseCM(s)
	if err != nil {
		return err
	}
	*cm = r
	return nil
}
