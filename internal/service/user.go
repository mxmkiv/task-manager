package service

import (
	"errors"
	"net/http"
	"task-manager/internal/encoder"
	"task-manager/internal/model"
	"task-manager/internal/repository"

	"github.com/labstack/echo/v5"
)

type UserService struct {
	userRepo *repository.UserRepository
	encoder  encoder.HashEncoder
}

func NewUserService(repo *repository.UserRepository, encoder encoder.HashEncoder) *UserService {
	return &UserService{
		userRepo: repo,
		encoder:  encoder,
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

func (u *UserService) UpdateUserData(dto *model.UpdateUserRequest, role string, requestId int) (*model.UserData, error) {

	// role check
	if role == model.UserType.RoleToString() {
		if dto.Role != nil {
			return nil, errors.New("user cant change role")
		}
	}

	// validation
	if dto.Login != nil && (len(*dto.Login) < 3 || len(*dto.Login) > 20) {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "login must be between 3 and 20 characters")
	}
	if dto.Password != nil && len(*dto.Password) < 5 {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "password must be at least 5 characters")
	}

	//data transfrom
	var updateList []model.UpdateFields

	if dto.Login != nil {
		updateList = append(updateList, model.UpdateFields{FieldName: "login", Data: dto.Login})
	}
	if dto.Password != nil {
		hash, err := u.encoder.Encode(*dto.Password)
		if err != nil {
			return nil, errors.New("hash generate error")
		}
		updateList = append(updateList, model.UpdateFields{FieldName: "password_hash", Data: hash})
	}
	if dto.Role != nil {
		updateList = append(updateList, model.UpdateFields{FieldName: "role", Data: dto.Role})
	}

	response, err := u.userRepo.UpdateUserData(&updateList, requestId)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (u *UserService) DeleteUser(requestId int, role string) error {

	err := u.userRepo.DeleteUser(requestId)
	if err != nil {
		return err
	}

	return nil
}
