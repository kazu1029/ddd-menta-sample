package userapp

import (
	"github.com/kazu1029/ddd-menta-sample/src/core/domain/tagdm"
	"github.com/kazu1029/ddd-menta-sample/src/core/domain/userdm"
	"github.com/kazu1029/ddd-menta-sample/src/core/domain/vo"
	"golang.org/x/xerrors"
)

type CreateUserApp struct {
	userRepo userdm.UserRepository
	tagRepo  tagdm.TagRepository
}

func NewCreateUserApp(userRepo userdm.UserRepository, tagRepo tagdm.TagRepository) *CreateUserApp {
	return &CreateUserApp{
		userRepo: userRepo,
		tagRepo:  tagRepo,
	}
}

type CreateUserRequest struct {
	UserName         string
	Email            string
	Password         string
	SelfIntroduction string
	Skills           []CreateUserSkillRequset
	WorkExperiences  []CreateUserWorkExperienceRequest
}

type CreateUserSkillRequset struct {
	ID                string
	YearsOfExperience int
}

type CreateUserWorkExperienceRequest struct {
	Description string
	YearFrom    uint
	YearTo      uint
}

type CreateUserResponse struct {
	ID string
}

func (app *CreateUserApp) Exec(req *CreateUserRequest) (*CreateUserResponse, error) {
	email, err := vo.NewEmail(req.Email)
	if err != nil {
		return nil, err
	}
	password, err := vo.NewPassword(req.Password)
	if err != nil {
		return nil, err
	}
	userID := userdm.NewUserID()

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

	// This might be user_domain_service logic
	var workExperiences []*userdm.UserWorkExperience
	if len(req.WorkExperiences) > 0 {
		for _, we := range req.WorkExperiences {
			workExperienceID := userdm.NewWorkExperienceID()
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

	user, err := userdm.NewUser(
		userID, req.UserName, email, password, req.SelfIntroduction, userSkills, workExperiences,
	)
	if err != nil {
		return nil, err
	}

	insertedUser, err := app.userRepo.Create(user)
	if err != nil {
		return nil, err
	}

	return &CreateUserResponse{ID: insertedUser.ID().Value()}, nil
}
