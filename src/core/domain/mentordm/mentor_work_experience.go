package mentordm

import (
	"golang.org/x/xerrors"
)

type MentorWorkExperience struct {
	id          WorkExperienceID
	description string
	yearFrom    YearFrom
	yearTo      YearTo
}

const (
	descriptionMaxLength = 1000
)

func NewMentorWorkExperience(id WorkExperienceID, description string, yearFrom YearFrom, yearTo YearTo) (*MentorWorkExperience, error) {
	if err := descriptionValidation(description); err != nil {
		return nil, err
	}
	return &MentorWorkExperience{
		id:          id,
		description: description,
		yearFrom:    yearFrom,
		yearTo:      yearTo,
	}, nil
}

func descriptionValidation(description string) error {
	if len(description) > descriptionMaxLength {
		return xerrors.New("description must be less than 1000")
	}
	return nil
}

func (we *MentorWorkExperience) ID() WorkExperienceID {
	return we.id
}

func (we *MentorWorkExperience) Description() string {
	return we.description
}

func (we *MentorWorkExperience) YearFrom() YearFrom {
	return we.yearFrom
}

func (we *MentorWorkExperience) YearTo() YearTo {
	return we.yearTo
}

func (we *MentorWorkExperience) ChangeDescription(description string) error {
	if err := descriptionValidation(description); err != nil {
		return err
	}
	we.description = description
	return nil
}

func (we *MentorWorkExperience) ChangeYearFrom(yearFrom YearFrom) {
	we.yearFrom = yearFrom
}

func (we *MentorWorkExperience) ChangeYearTo(yearTo YearTo) {
	we.yearTo = yearTo
}
