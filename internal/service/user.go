package service

import (
	"task-manager/internal/model"
	"task-manager/internal/repository"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		userRepo: repo,
	}
}

func (u *UserService) GetAllUsers() (*[]*model.UserData, error) {
	rawData, err := u.userRepo.GetAllUsers()
	if err != nil {
		return nil, err
	}

	resData := make([]*model.UserData, len(*rawData))
	for index, user := range *rawData {
		d := &model.UserData{
			Id:        user.Id,
			Login:     user.Login,
			Role:      user.Role,
			CreatedAt: user.CreatedAt,
			UpdateAt:  user.UpdateAt,
		}
		resData[index] = d
	}

	return &resData, nil
}
