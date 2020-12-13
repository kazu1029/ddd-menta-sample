package plandm

import "golang.org/x/xerrors"

type Price uint

const (
	priceMin = 1000
)

func NewPrice(price uint) (Price, error) {
	if err := priceValidation(price); err != nil {
		return Price(0), err
	}
	return Price(price), nil
}

func priceValidation(price uint) error {
	if price <= 1000 {
		return xerrors.New("price must be over 1000 yen")
	}
	return nil
}

func (p Price) Value() uint {
	return uint(p)
}

func (p Price) Equals(p2 Price) bool {
	return p == p2
}
