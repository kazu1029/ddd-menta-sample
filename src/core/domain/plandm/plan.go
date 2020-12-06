package plandm

import (
	"github.com/kazu1029/ddd-menta-sample/src/core/domain/mentordm"
	"github.com/kazu1029/ddd-menta-sample/src/core/domain/tagdm"
)

type Plan struct {
	ID       PlanID
	MentorID mentordm.MentorID
	Title    string
	// TODO: create categorydm
	// CategoryIDs []Category
	TagIDs []tagdm.TagID
}
