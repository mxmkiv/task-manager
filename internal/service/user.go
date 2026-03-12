package service

import (
	"net/http"
	"task-manager/internal/model"
	"task-manager/internal/repository"

	"github.com/labstack/echo/v5"
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
			UpdatedAt: user.UpdatedAt,
		}
		resData[index] = d
	}

	return &resData, nil
}

func (u *UserService) GetUserById(id int) (*model.UserData, error) {

	rawData, err := u.userRepo.GetUserById(id)
	if err != nil {
		if err.Error() == "no user founds" {
			return nil, echo.NewHTTPError(http.StatusOK, "user with this id not exists")
		}
		return nil, err
	}

	resData := &model.UserData{
		Id:        rawData.Id,
		Login:     rawData.Login,
		Role:      rawData.Role,
		CreatedAt: rawData.CreatedAt,
		UpdatedAt: rawData.UpdatedAt,
	}

	return resData, nil

}
