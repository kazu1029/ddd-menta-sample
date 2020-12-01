package mentorapp

import (
	"github.com/kazu1029/ddd-menta-sample/src/core/domain/mentordm"
	"github.com/kazu1029/ddd-menta-sample/src/core/domain/tagdm"
	"github.com/kazu1029/ddd-menta-sample/src/core/domain/vo"
	"golang.org/x/xerrors"
)

type CreateMentorApp struct {
	mentorRepo mentordm.MentorRepository
	tagRepo    tagdm.TagRepository
}

func NewCreateMentorApp(mentorRepo mentordm.MentorRepository, tagRepo tagdm.TagRepository) *CreateMentorApp {
	return &CreateMentorApp{
		mentorRepo: mentorRepo,
		tagRepo:    tagRepo,
	}
}

type CreateMentorRequest struct {
	MentorName       string
	Email            string
	Password         string
	SelfIntroduction string
	Skills           []CreateMentorSkillRequset
	WorkExperiences  []CreateMentorWorkExperienceRequest
}

type CreateMentorSkillRequset struct {
	ID                string
	YearsOfExperience int
}

type CreateMentorWorkExperienceRequest struct {
	Description string
	YearFrom    uint
	YearTo      uint
}

type CreateMentorResponse struct {
	ID string
}

func (app *CreateMentorApp) Exec(req *CreateMentorRequest) (*CreateMentorResponse, error) {
	email, err := vo.NewEmail(req.Email)
	if err != nil {
		return nil, err
	}
	password, err := vo.NewPassword(req.Password)
	if err != nil {
		return nil, err
	}
	mentorID := mentordm.NewMentorID()

	var mentorSkillIDs []string
	for _, skill := range req.Skills {
		mentorSkillIDs = append(mentorSkillIDs, skill.ID)
	}

	tagDomainService := tagdm.NewTagDomainService(app.tagRepo)
	var mentorSkills []*mentordm.MentorSkill
	for _, skill := range req.Skills {
		tagID, err := tagdm.NewTagIDWithStr(skill.ID)
		if err != nil {
			return nil, err
		}
		if ok := tagDomainService.Exists(tagID); !ok {
			return nil, xerrors.Errorf("invalid skill id, %d", skill.ID)
		}
		yoe, err := mentordm.NewYearsOfExperience(mentordm.YearsOfExperience(skill.YearsOfExperience))
		if err != nil {
			return nil, err
		}
		us, err := mentordm.NewMentorSkill(tagID, yoe)
		if err != nil {
			return nil, err
		}
		mentorSkills = append(mentorSkills, us)
	}

	// This might be mentor_domain_service logic
	var workExperiences []*mentordm.MentorWorkExperience
	if len(req.WorkExperiences) > 0 {
		for _, we := range req.WorkExperiences {
			workExperienceID := mentordm.NewWorkExperienceID()
			yearFrom, err := mentordm.NewYearFrom(we.YearFrom)
			if err != nil {
				return nil, err
			}
			yearTo, err := mentordm.NewYearTo(yearFrom.Value(), we.YearTo)
			if err != nil {
				return nil, err
			}
			experience, err := mentordm.NewMentorWorkExperience(
				workExperienceID,
				we.Description,
				yearFrom,
				yearTo,
			)
			if err != nil {
				return nil, err
			}
			workExperiences = append(workExperiences, experience)
		}
	}

	mentor, err := mentordm.NewMentor(
		mentorID, req.MentorName, email, password, req.SelfIntroduction, mentorSkills, workExperiences,
	)
	if err != nil {
		return nil, err
	}

	insertedMentor, err := app.mentorRepo.Create(mentor)
	if err != nil {
		return nil, err
	}

	return &CreateMentorResponse{ID: insertedMentor.ID().Value()}, nil
}
