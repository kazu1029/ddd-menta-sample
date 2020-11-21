package menteedm

import "golang.org/x/xerrors"

type YearsOfExperience int

const (
	LessThanOneYear YearsOfExperience = iota + 1
	LessThanThreeYears
	LessThanFiveYears
	OverFiveYears
)

const (
	LessThanOneYearStr    string = "less than 1 year"
	LessThanThreeYearsStr string = "less than three years"
	LessThanFiveYearsStr  string = "less than five years"
	OverFiveYearsStr      string = "over five years"
)

func NewYearsOfExperience(yearsOfExperience YearsOfExperience) (YearsOfExperience, error) {
	if yearsOfExperience < LessThanOneYear {
		return YearsOfExperience(0), xerrors.Errorf("years_of_experience must be over than %d", LessThanOneYear)
	}

	if yearsOfExperience > OverFiveYears {
		return YearsOfExperience(0), xerrors.Errorf("years_of_experience must be less than %d", OverFiveYears)
	}

	return YearsOfExperience(yearsOfExperience), nil
}

func (years YearsOfExperience) Value() int {
	return int(years)
}

func (years YearsOfExperience) String() string {
	switch years {
	case LessThanOneYear:
		return LessThanOneYearStr
	case LessThanThreeYears:
		return LessThanThreeYearsStr
	case LessThanFiveYears:
		return LessThanFiveYearsStr
	case OverFiveYears:
		return OverFiveYearsStr
	default:
		return "Unknown"
	}
}

func (years YearsOfExperience) Equals(years2 YearsOfExperience) bool {
	return years == years2
}
