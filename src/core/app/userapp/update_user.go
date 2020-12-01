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

	userSkillIDs := make([]string, len(req.Skills))
	for i := 0; i < len(req.Skills); i++ {
		userSkillIDs[i] = req.Skills[i].ID
	}

	user, err := app.userRepo.FindByID(userID)
	if err != nil {
		return nil, err
	}

	tagDomainService := tagdm.NewTagDomainService(app.tagRepo)
	currentSkillsMap := make(map[string]*userdm.UserSkill, len(user.Skills()))
	for _, us := range user.Skills() {
		currentSkillsMap[us.ID().Value()] = us
	}

	tagIDs, err := tagdm.NewTagIDs(userSkillIDs)
	if err != nil {
		return nil, err
	}
	if ok := tagDomainService.ExistsWithIDs(tagIDs); !ok {
		return nil, xerrors.Errorf("invalid skill ids, %v", tagIDs)
	}

	userSkills := make([]*userdm.UserSkill, len(req.Skills))
	for _, skill := range req.Skills {
		tagID, err := tagdm.NewTagIDWithStr(skill.ID)
		if err != nil {
			return nil, err
		}
		yoe, err := userdm.NewYearsOfExperience(userdm.YearsOfExperience(skill.YearsOfExperience))
		if err != nil {
			return nil, err
		}
		var us *userdm.UserSkill
		us, ok := currentSkillsMap[tagID.Value()]
		if ok {
			us.ChangeYearsOfExperience(yoe)
		} else {
			us, err = userdm.NewUserSkill(tagID, userID, yoe)
			if err != nil {
				return nil, err
			}
		}
		userSkills = append(userSkills, us)
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
	workExperiences := make([]*userdm.UserWorkExperience, len(req.WorkExperiences))
	var currentWorkExpMaps map[string]*userdm.UserWorkExperience
	for _, we := range user.WorkExperiences() {
		currentWorkExpMaps[we.ID().Value()] = we
	}
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
			experience, ok := currentWorkExpMaps[workExperienceID.Value()]
			if ok {
				err = experience.ChangeDescription(we.Description)
				if err != nil {
					return nil, err
				}
				experience.ChangeYearFrom(yearFrom)
				experience.ChangeYearTo(yearTo)
			} else {
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
