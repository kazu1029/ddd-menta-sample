package menteedm

import (
	"golang.org/x/xerrors"
)

type MenteeWorkExperience struct {
	id          WorkExperienceID
	ownerID     MenteeID
	description string
	yearFrom    YearFrom
	yearTo      YearTo
}

const (
	descriptionMaxLength = 1000
)

func NewMenteeWorkExperience(id WorkExperienceID, ownerID MenteeID, description string, yearFrom YearFrom, yearTo YearTo) (*MenteeWorkExperience, error) {
	if err := descriptionValidation(description); err != nil {
		return nil, err
	}
	return &MenteeWorkExperience{
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

func (we *MenteeWorkExperience) ID() WorkExperienceID {
	return we.id
}

func (we *MenteeWorkExperience) OwnerID() MenteeID {
	return we.ownerID
}

func (we *MenteeWorkExperience) Description() string {
	return we.description
}
