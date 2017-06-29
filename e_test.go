package money

import (
	"fmt"
	"testing"
)

func Test_Exchange(t *testing.T) {
	a := FlatExchange{
		"CAD": 1.1,
		"GBP": 0.9,
	}

	n, err := ExchangeTo(a, SParse("10", 0), "CAD", "GBP")
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	fmt.Println(n)
}
