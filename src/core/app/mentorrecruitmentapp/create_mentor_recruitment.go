package mentorrecruitmentapp

import (
	"github.com/kazu1029/ddd-menta-sample/src/core/domain/menteedm"
	"github.com/kazu1029/ddd-menta-sample/src/core/domain/mentorrecruitmentdm"
)

type CreateMentorRecruitmentApp struct {
	mRecruitmentRepo mentorrecruitmentdm.MentorRecruitmentRepository
}

func NewCreateMentorRecruitmentApp(mRecruitmentRepo mentorrecruitmentdm.MentorRecruitmentRepository) *CreateMentorRecruitmentApp {
	return &CreateMentorRecruitmentApp{
		mRecruitmentRepo: mRecruitmentRepo,
	}
}

type CreateMentorRecruitmentRequest struct {
	MenteeID string
	Title    string
	Budget   struct {
		Fee            int
		IsSubscription bool
	}
	Description string
	Status      int
}

type CreateMentorRecruitmentResponse struct {
	ID string
}

func (app *CreateMentorRecruitmentApp) Exec(req *CreateMentorRecruitmentRequest) (*CreateMentorRecruitmentResponse, error) {
	menteeID, err := menteedm.NewMenteeIDWithStr(req.MenteeID)
	if err != nil {
		return nil, err
	}

	budget, err := mentorrecruitmentdm.NewBudget(req.Budget.Fee, req.Budget.IsSubscription)
	if err != nil {
		return nil, err
	}

	status, err := mentorrecruitmentdm.NewStatus(req.Status)
	if err != nil {
		return nil, err
	}

	mentorRecruitmentID := mentorrecruitmentdm.NewMentorRecruitmentID()

	mentorRecruitment, err := mentorrecruitmentdm.NewMentorRecruitment(
		mentorRecruitmentID,
		menteeID,
		req.Title,
		budget,
		req.Description,
		status,
	)
	if err != nil {
		return nil, err
	}

	insertedMentorRecruitment, err := app.mRecruitmentRepo.Create(mentorRecruitment)
	if err != nil {
		return nil, err
	}

	return &CreateMentorRecruitmentResponse{ID: insertedMentorRecruitment.ID().Value()}, nil
}
