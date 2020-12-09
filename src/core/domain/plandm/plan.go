package plandm

import (
	"github.com/kazu1029/ddd-menta-sample/src/core/domain/categorydm"
	"github.com/kazu1029/ddd-menta-sample/src/core/domain/mentordm"
	"github.com/kazu1029/ddd-menta-sample/src/core/domain/tagdm"
	"golang.org/x/xerrors"
)

type Plan struct {
	id             PlanID
	mentorID       mentordm.MentorID
	title          string
	description    string
	status         Status
	isSubscription bool
	price          Price
	categoryIDs    []categorydm.CategoryID
	tagIDs         []tagdm.TagID
}

const (
	titleMaxLength       = 255
	descriptionMinLength = 3000
)

func NewPlan(id PlanID, mentorID mentordm.MentorID, title string, description string, status Status, isSbuscription bool, price Price, categoryIDs []categorydm.CategoryID, tagIDs []tagdm.TagID) (*Plan, error) {
	if err := titleValidation(title); err != nil {
		return nil, err
	}

	if err := descriptionValidation(description); err != nil {
		return nil, err
	}

	return &Plan{
		id:             id,
		mentorID:       mentorID,
		title:          title,
		description:    description,
		status:         status,
		isSubscription: isSbuscription,
		price:          price,
		categoryIDs:    categoryIDs,
		tagIDs:         tagIDs,
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

func (p *Plan) ID() PlanID {
	return p.id
}

func (p *Plan) MenteeID() mentordm.MentorID {
	return p.mentorID
}

func (p *Plan) Title() string {
	return p.title
}

func (p *Plan) Description() string {
	return p.description
}

func (p *Plan) Status() Status {
	return p.status
}

func (p *Plan) IsSubScription() bool {
	return p.isSubscription
}

func (p *Plan) Price() Price {
	return p.price
}

func (p *Plan) CategoryIDs() []categorydm.CategoryID {
	return p.categoryIDs
}

func (p *Plan) TagIDs() []tagdm.TagID {
	return p.tagIDs
}
