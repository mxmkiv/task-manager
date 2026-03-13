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

func (u *UserService) UpdateUserData(dto *model.UpdateUserRequest, role string, requestId int) error {

	updatesList := make(map[string]string)

	/*

		new data validation

	*/

	if role != model.AdminType.RoleToString() {
		if dto.Role != nil {
			return errors.New("user can't change role")
		}

		if dto.Login != nil && *dto.Login != "" {
			updatesList["login"] = *dto.Login
		}

		if dto.Password != nil && *dto.Password != "" {
			updatesList["password_hash"] = *dto.Password
		}

	} else {
		if dto.Role != nil && *dto.Role != "" {
			if *dto.Role == model.AdminType.RoleToString() || *dto.Role == model.UserType.RoleToString() {
				updatesList["role"] = *dto.Role
			} else {
				return errors.New("incorrect role")
			}
		}

		if dto.Login != nil && *dto.Login != "" {
			updatesList["login"] = *dto.Login
		}

		if dto.Password != nil && *dto.Password != "" {
			hash, err := u.encoder.Encode(*dto.Password)
			if err != nil {
				return errors.New("hash generate error")
			}
			updatesList["password_hash"] = hash
		}
	}

	if len(updatesList) == 0 {
		return errors.New("no fields to update")
	}

	if err := u.userRepo.UpdateUserData(updatesList, requestId); err != nil {
		return err
	}

	return nil
}
