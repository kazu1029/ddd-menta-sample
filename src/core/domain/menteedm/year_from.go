package menteedm

import "golang.org/x/xerrors"

type YearFrom uint

const (
	yearFromMin = 1970
)

func NewYearFrom(year uint) (YearFrom, error) {
	if year < yearFromMin {
		return YearFrom(0), xerrors.Errorf("year from must be over %d", yearFromMin)
	}
	return YearFrom(year), nil
}

func (y YearFrom) Value() uint {
	return uint(y)
}

func (y1 YearFrom) Equals(y2 YearFrom) bool {
	return y1 == y2
}
