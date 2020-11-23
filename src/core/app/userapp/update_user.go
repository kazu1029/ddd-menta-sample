package userapp

import (
	"github.com/kazu1029/ddd-menta-sample/src/core/domain/tagdm"
	"github.com/kazu1029/ddd-menta-sample/src/core/domain/userdm"
	"github.com/kazu1029/ddd-menta-sample/src/core/domain/vo"
	"golang.org/x/xerrors"
)

type UpdateUserApp struct {
	userRepo userdm.UserRepository
	tagRepo  tagdm.TagRepository
}

func NewUpdateUserApp(userRepo userdm.UserRepository, tagRepo tagdm.TagRepository) *UpdateUserApp {
	return &UpdateUserApp{
		userRepo: userRepo,
		tagRepo:  tagRepo,
	}
}

type UpdateUserRequest struct {
	ID               string
	UserName         string
	Email            string
	Password         string
	SelfIntroduction string
	Skills           []UpdateUserSkillRequest
	WorkExperiences  []UpdateUserWorkExperienceRequest
}

type UpdateUserSkillRequest struct {
	ID                string
	YearsOfExperience int
}

type UpdateUserWorkExperienceRequest struct {
	ID          string
	Description string
	YearFrom    uint
	YearTo      uint
}

type UpdateUserResponse struct {
	ID               string
	UserName         string
	Email            string
	SelfIntroduction string
	Skills           []UpdateUserSkillResponse
	WorkExperiences  []UpdateUserWorkExperienceResponse
}

type UpdateUserSkillResponse struct {
	ID                string
	YearsOfExperience int
}

type UpdateUserWorkExperienceResponse struct {
	ID          string
	Description string
	YearFrom    uint
	YearTo      uint
}

func (app *UpdateUserApp) Exec(req *UpdateUserRequest) (*UpdateUserResponse, error) {
	userID, err := userdm.NewUserIDWithStr(req.ID)
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

	var userSkillIDs []string
	for _, skill := range req.Skills {
		userSkillIDs = append(userSkillIDs, skill.ID)
	}

	tagDomainService := tagdm.NewTagDomainService(app.tagRepo)
	var userSkills []*userdm.UserSkill
	for _, skill := range req.Skills {
		tagID, err := tagdm.NewTagIDWithStr(skill.ID)
		if err != nil {
			return nil, err
		}
		if ok := tagDomainService.Exists(tagID); !ok {
			return nil, xerrors.Errorf("invalid skill id, %d", skill.ID)
		}
		yoe, err := userdm.NewYearsOfExperience(userdm.YearsOfExperience(skill.YearsOfExperience))
		if err != nil {
			return nil, err
		}
		var us *userdm.UserSkill
		us, err = app.userRepo.FindSkillBySkillID(userID, tagID)
		if err != nil {
			return nil, err
		}
		if us == nil {
			us, err = userdm.NewUserSkill(tagID, userID, yoe)
			if err != nil {
				return nil, err
			}
		} else {
			us.ChangeYearsOfExperience(yoe)
		}
		userSkills = append(userSkills, us)
	}

	user, err := app.userRepo.FindByID(userID)
	if err != nil {
		return nil, err
	}

	user.ChangeEmail(email)
	user.ChangePassword(password)
	err = user.ChangeUserName(req.UserName)
	if err != nil {
		return nil, err
	}
	err = user.ChangeSelfIntroduction(req.SelfIntroduction)
	if err != nil {
		return nil, err
	}
	err = user.ChangeSkills(userSkills)
	if err != nil {
		return nil, err
	}
	var workExperiences []*userdm.UserWorkExperience
	if len(req.WorkExperiences) > 0 {
		for _, we := range req.WorkExperiences {
			workExperienceID, err := userdm.NewWorkExperienceIDWithStr(we.ID)
			if err != nil {
				return nil, err
			}
			yearFrom, err := userdm.NewYearFrom(we.YearFrom)
			if err != nil {
				return nil, err
			}
			yearTo, err := userdm.NewYearTo(yearFrom.Value(), we.YearTo)
			if err != nil {
				return nil, err
			}
			var experience *userdm.UserWorkExperience
			experience, err = app.userRepo.FindWorkExperienceByWorkExperienceID(userID, workExperienceID)
			if err != nil {
				return nil, err
			}
			if experience == nil {
				experience, err = userdm.NewUserWorkExperience(
					workExperienceID,
					userID,
					we.Description,
					yearFrom,
					yearTo,
				)
				if err != nil {
					return nil, err
				}
			} else {
				err = experience.ChangeDescription(we.Description)
				if err != nil {
					return nil, err
				}
				experience.ChangeYearFrom(yearFrom)
				experience.ChangeYearTo(yearTo)
			}
			workExperiences = append(workExperiences, experience)
		}
	}

	user.ChangeWorkExperiences(workExperiences)
	if err != nil {
		return nil, err
	}

	updatedUser, err := app.userRepo.Update(user)
	if err != nil {
		return nil, err
	}

	var skillsResponse []UpdateUserSkillResponse
	for _, skill := range updatedUser.Skills() {
		s := UpdateUserSkillResponse{
			ID:                skill.ID().Value(),
			YearsOfExperience: skill.YearsOfExperience().Value(),
		}
		skillsResponse = append(skillsResponse, s)
	}

	var workExperiencesResponse []UpdateUserWorkExperienceResponse
	for _, we := range updatedUser.WorkExperiences() {
		e := UpdateUserWorkExperienceResponse{
			ID:          we.ID().Value(),
			Description: we.Description(),
			YearFrom:    we.YearFrom().Value(),
			YearTo:      we.YearTo().Value(),
		}
		workExperiencesResponse = append(workExperiencesResponse, e)
	}

	return &UpdateUserResponse{
		ID:               updatedUser.ID().Value(),
		Email:            updatedUser.Email().Value(),
		UserName:         updatedUser.UserName(),
		SelfIntroduction: updatedUser.SelfIntroduction(),
		Skills:           skillsResponse,
		WorkExperiences:  workExperiencesResponse,
	}, nil
}
