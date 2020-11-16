package userdm

import "golang.org/x/xerrors"

type YearTo int

const (
	yearToMin = 1970
)

func NewYearTo(yearFrom, yearTo int) (YearTo, error) {
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

func (y YearTo) Value() int {
	return int(y)
}

func (y1 YearTo) Equals(y2 YearTo) bool {
	return y1 == y2
}
