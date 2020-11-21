package menteedm

import "golang.org/x/xerrors"

type YearTo uint

const (
	yearToMin = 1970
)

func NewYearTo(yearFrom, yearTo uint) (YearTo, error) {
	if yearTo == 0 {
		return YearTo(0), xerrors.New("year to must be not be zero")
	}
	if yearTo < yearToMin {
		return YearTo(0), xerrors.Errorf("year to must be over %d", yearToMin)
	}
	if yearTo < yearFrom {
		return YearTo(0), xerrors.New("year to must be more than year from")
	}
	return YearTo(yearTo), nil
}

func (y YearTo) Value() uint {
	return uint(y)
}

func (y1 YearTo) Equals(y2 YearTo) bool {
	return y1 == y2
}
