package money

import "fmt"

type Exchanger interface {
	Exchange(string, string) (float32, error)
}

func ExchangeTo(e Exchanger, n M, funit, tunit string) (M, error) {
	rate, err := e.Exchange(funit, tunit)
	return M(float32(n) / rate), err
}

func ExchangeFrom(e Exchanger, n M, funit, tunit string) (M, error) {
	rate, err := e.Exchange(funit, tunit)
	return M(float32(n) / rate), err
}

type FlatExchange map[string]float32

func (de FlatExchange) Exchange(funit, tunit string) (float32, error) {

	frate, ok := de[funit]
	if !ok {
		return 1, fmt.Errorf("Could not understand unit %s", funit)
	}
	trate, ok := de[tunit]

	if !ok {
		return 1, fmt.Errorf("Could not understand unit %s", tunit)
	}

	return frate / trate, nil
}

type RateExchange struct {
	Buy, Sell map[string]float32
}

func (re RateExchange) Exchange(funit, tunit string) (float32, error) {

	frate, ok := re.Sell[funit]
	if !ok {
		return 1, fmt.Errorf("Could not understand unit %s", funit)
	}
	trate, ok := re.Buy[tunit]

	if !ok {
		return 1, fmt.Errorf("Could not understand unit %s", tunit)
	}

	return frate / trate, nil
}
