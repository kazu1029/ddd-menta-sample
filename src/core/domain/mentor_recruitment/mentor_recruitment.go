package mentor_recruitmentdm

import (
	"github.com/kazu1029/ddd-menta-sample/src/core/domain/menteedm"
	"golang.org/x/xerrors"
)

type MentorRecruitment struct {
	id          MentorRecruitmentID
	menteeID    menteedm.MenteeID
	title       string
	budget      Budget
	description string
	// TOOD: add isClosed or status for managing closed or not
	status Status
}

const (
	titleMaxLength       = 255
	descriptionMinLength = 3000
)

func NewMentorRecruitment(id MentorRecruitmentID, menteeID menteedm.MenteeID, title string, budget Budget, description string) (*MentorRecruitment, error) {
	if err := titleValidation(title); err != nil {
		return nil, err
	}

	if err := descriptionValidation(description); err != nil {
		return nil, err
	}

	return &MentorRecruitment{
		id:          id,
		menteeID:    menteeID,
		title:       title,
		budget:      budget,
		description: description,
	}, nil
}

func titleValidation(title string) error {
	if len(title) > titleMaxLength {
		return xerrors.Errorf("title lenght must be less than %d", titleMaxLength)
	}
	return nil
}

func descriptionValidation(description string) error {
	if len(description) < descriptionMinLength {
		return xerrors.Errorf("description length must be over %d", descriptionMinLength)
	}
	return nil
}

func (mr *MentorRecruitment) ID() MentorRecruitmentID {
	return mr.id
}

func (mr *MentorRecruitment) MenteeID() menteedm.MenteeID {
	return mr.menteeID
}

func (mr *MentorRecruitment) Title() string {
	return mr.title
}

func (mr *MentorRecruitment) Budget() Budget {
	return mr.budget
}

func (mr *MentorRecruitment) Description() string {
	return mr.description
}
