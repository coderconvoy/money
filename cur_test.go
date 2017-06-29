package money

import (
	"encoding/json"
	"fmt"
	"testing"
)

func Test_CanString(t *testing.T) {
	s := "100CAD"
	m := s[len(s)-3:]
	c := s[:len(s)-3]
	fmt.Println("TEST CAN STRING")
	fmt.Println(m)
	fmt.Println(c)
}

func Test_MultiMarshal(t *testing.T) {
	s := []*CM{
		{1000, "CAD"},
		{600, "USD"},
		{300, "GBP"},
		{4050, "AND"},
	}
	m, err := json.Marshal(s)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	fmt.Println(string(m))
	var ss []*CM
	err = json.Unmarshal(m, &ss)
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	for k, v := range s {
		if ss[k].mon != v.mon || ss[k].cur != v.cur {
			t.Logf("Expected %s , got %s", v, ss[k])
			t.Fail()
		}
	}

}
