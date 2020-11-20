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
	// Skills         []string
	// WorkExperiences
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
		us, err := userdm.NewUserSkill(tagID, userID, yoe)
		if err != nil {
			return nil, err
		}
		userSkills = append(userSkills, us)
	}

	user, err := app.userRepo.FindByID(userID)
	if err != nil {
		return nil, err
	}

	err = user.ChangeEmail(email)
	if err != nil {
		return nil, err
	}
	err = user.ChangePassword(password)
	if err != nil {
		return nil, err
	}
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
			experience, err := userdm.NewUserWorkExperience(
				workExperienceID,
				userID,
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

	err = user.ChangeWorkExperiences(workExperiences)
	if err != nil {
		return nil, err
	}

	updatedUser, err := app.userRepo.Update(user)
	if err != nil {
		return nil, err
	}

	return &UpdateUserResponse{
		ID:               updatedUser.ID().Value(),
		Email:            updatedUser.Email().Value(),
		UserName:         updatedUser.UserName(),
		SelfIntroduction: updatedUser.SelfIntroduction(),
	}, nil
}
