package mentorapp

import (
	"github.com/kazu1029/ddd-menta-sample/src/core/domain/mentordm"
	"github.com/kazu1029/ddd-menta-sample/src/core/domain/tagdm"
	"github.com/kazu1029/ddd-menta-sample/src/core/domain/vo"
	"golang.org/x/xerrors"
)

type UpdateMentorApp struct {
	mentorRepo mentordm.MentorRepository
	tagRepo    tagdm.TagRepository
}

func NewUpdateMentorApp(mentorRepo mentordm.MentorRepository, tagRepo tagdm.TagRepository) *UpdateMentorApp {
	return &UpdateMentorApp{
		mentorRepo: mentorRepo,
		tagRepo:    tagRepo,
	}
}

type UpdateMentorRequest struct {
	ID               string
	MentorName       string
	Email            string
	Password         string
	SelfIntroduction string
	Skills           []UpdateMentorSkillRequest
	WorkExperiences  []UpdateMentorWorkExperienceRequest
}

type UpdateMentorSkillRequest struct {
	ID                string
	YearsOfExperience int
}

type UpdateMentorWorkExperienceRequest struct {
	ID          string
	Description string
	YearFrom    uint
	YearTo      uint
}

type UpdateMentorResponse struct {
	ID               string
	MentorName       string
	Email            string
	SelfIntroduction string
	Skills           []UpdateMentorSkillResponse
	WorkExperiences  []UpdateMentorWorkExperienceResponse
}

type UpdateMentorSkillResponse struct {
	ID                string
	YearsOfExperience int
}

type UpdateMentorWorkExperienceResponse struct {
	ID          string
	Description string
	YearFrom    uint
	YearTo      uint
}

func (app *UpdateMentorApp) Exec(req *UpdateMentorRequest) (*UpdateMentorResponse, error) {
	mentorID, err := mentordm.NewMentorIDWithStr(req.ID)
	if err != nil {
		return nil, err
	}
	email, err := vo.NewEmail(req.Email)
	if err != nil {
		return nil, err
	}
	password, err := vo.NewPassword(req.Password)
	if err != nil {
		return nil, err
	}

	mentorSkillIDs := make([]string, len(req.Skills))
	for i := 0; i < len(req.Skills); i++ {
		mentorSkillIDs[i] = req.Skills[i].ID
	}

	mentor, err := app.mentorRepo.FindByID(mentorID)
	if err != nil {
		return nil, err
	}

	tagDomainService := tagdm.NewTagDomainService(app.tagRepo)
	currentSkillsMap := make(map[string]*mentordm.MentorSkill, len(mentor.Skills()))
	for _, us := range mentor.Skills() {
		currentSkillsMap[us.ID().Value()] = us
	}

	tagIDs, err := tagdm.NewTagIDs(mentorSkillIDs)
	if err != nil {
		return nil, err
	}
	if ok := tagDomainService.ExistsWithIDs(tagIDs); !ok {
		return nil, xerrors.Errorf("invalid skill ids, %v", tagIDs)
	}

	mentorSkills := make([]*mentordm.MentorSkill, len(req.Skills))
	for _, skill := range req.Skills {
		tagID, err := tagdm.NewTagIDWithStr(skill.ID)
		if err != nil {
			return nil, err
		}
		yoe, err := mentordm.NewYearsOfExperience(mentordm.YearsOfExperience(skill.YearsOfExperience))
		if err != nil {
			return nil, err
		}
		var us *mentordm.MentorSkill
		us, ok := currentSkillsMap[tagID.Value()]
		if ok {
			us.ChangeYearsOfExperience(yoe)
		} else {
			us, err = mentordm.NewMentorSkill(tagID, yoe)
			if err != nil {
				return nil, err
			}
		}
		mentorSkills = append(mentorSkills, us)
	}

	mentor.ChangeEmail(email)
	mentor.ChangePassword(password)
	err = mentor.ChangeMentorName(req.MentorName)
	if err != nil {
		return nil, err
	}
	err = mentor.ChangeSelfIntroduction(req.SelfIntroduction)
	if err != nil {
		return nil, err
	}
	err = mentor.ChangeSkills(mentorSkills)
	if err != nil {
		return nil, err
	}
	workExperiences := make([]*mentordm.MentorWorkExperience, len(req.WorkExperiences))
	var currentWorkExpMaps map[string]*mentordm.MentorWorkExperience
	for _, we := range mentor.WorkExperiences() {
		currentWorkExpMaps[we.ID().Value()] = we
	}
	if len(req.WorkExperiences) > 0 {
		for _, we := range req.WorkExperiences {
			workExperienceID, err := mentordm.NewWorkExperienceIDWithStr(we.ID)
			if err != nil {
				return nil, err
			}
			yearFrom, err := mentordm.NewYearFrom(we.YearFrom)
			if err != nil {
				return nil, err
			}
			yearTo, err := mentordm.NewYearTo(yearFrom.Value(), we.YearTo)
			if err != nil {
				return nil, err
			}
			var experience *mentordm.MentorWorkExperience
			experience, ok := currentWorkExpMaps[workExperienceID.Value()]
			if ok {
				err = experience.ChangeDescription(we.Description)
				if err != nil {
					return nil, err
				}
				experience.ChangeYearFrom(yearFrom)
				experience.ChangeYearTo(yearTo)
			} else {
				experience, err = mentordm.NewMentorWorkExperience(
					workExperienceID,
					we.Description,
					yearFrom,
					yearTo,
				)
				if err != nil {
					return nil, err
				}
			}
			workExperiences = append(workExperiences, experience)
		}
	}

	mentor.ChangeWorkExperiences(workExperiences)
	if err != nil {
		return nil, err
	}

	updatedMentor, err := app.mentorRepo.Update(mentor)
	if err != nil {
		return nil, err
	}

	var skillsResponse []UpdateMentorSkillResponse
	for _, skill := range updatedMentor.Skills() {
		s := UpdateMentorSkillResponse{
			ID:                skill.ID().Value(),
			YearsOfExperience: skill.YearsOfExperience().Value(),
		}
		skillsResponse = append(skillsResponse, s)
	}

	var workExperiencesResponse []UpdateMentorWorkExperienceResponse
	for _, we := range updatedMentor.WorkExperiences() {
		e := UpdateMentorWorkExperienceResponse{
			ID:          we.ID().Value(),
			Description: we.Description(),
			YearFrom:    we.YearFrom().Value(),
			YearTo:      we.YearTo().Value(),
		}
		workExperiencesResponse = append(workExperiencesResponse, e)
	}

	return &UpdateMentorResponse{
		ID:               updatedMentor.ID().Value(),
		Email:            updatedMentor.Email().Value(),
		MentorName:       updatedMentor.MentorName(),
		SelfIntroduction: updatedMentor.SelfIntroduction(),
		Skills:           skillsResponse,
		WorkExperiences:  workExperiencesResponse,
	}, nil
}
