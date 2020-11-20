package userdm

import (
	"golang.org/x/xerrors"
)

type UserWorkExperience struct {
	id          WorkExperienceID
	ownerID     UserID
	description string
	yearFrom    YearFrom
	yearTo      YearTo
}

const (
	descriptionMaxLength = 1000
)

func NewUserWorkExperience(id WorkExperienceID, ownerID UserID, description string, yearFrom YearFrom, yearTo YearTo) (*UserWorkExperience, error) {
	if err := descriptionValidation(description); err != nil {
		return nil, err
	}
	return &UserWorkExperience{
		id:          id,
		ownerID:     ownerID,
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

func (we *UserWorkExperience) ID() WorkExperienceID {
	return we.id
}

func (we *UserWorkExperience) OwnerID() UserID {
	return we.ownerID
}

func (we *UserWorkExperience) Description() string {
	return we.description
}
