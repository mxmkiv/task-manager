package service

import (
	"net/http"
	"task-manager/internal/auth"
	"task-manager/internal/model"
	"task-manager/internal/repository"

	"github.com/labstack/echo/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	jwtSecretKey string
	userRepo     *repository.UserRepository
}

func NewAuthService(key string, repo *repository.UserRepository) *AuthService {
	return &AuthService{
		jwtSecretKey: key,
		userRepo:     repo,
	}
}

func (authSvc *AuthService) Register(dto *model.RequestData) (*model.ResponseData, error) {

	// validation
	if len(dto.Login) < 3 || len(dto.Login) > 20 {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "login must be between 3 and 20 characters")
	}
	if len(dto.Password) < 5 {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "password must be at least 5 characters")
	}

	//exists check

	//create user
	user, err := model.NewUser(dto.Login, dto.Password)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "failed to create new user")
	}

	//repository
	if err := authSvc.userRepo.Create(user); err != nil {
		if err.Error() == "user with this login already exists" {
			return nil, echo.NewHTTPError(http.StatusConflict, err.Error())
		}
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "[db] failed to create user")
	}

	//token generation
	token, err := auth.GenerateToken(user.Id, user.Login, authSvc.jwtSecretKey)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "token generate error")
	}

	return &model.ResponseData{User: *user, Token: token}, nil
}

func (authSvc *AuthService) Login(dto *model.RequestData) (*model.ResponseData, error) {

	// validation
	if dto.Login == "" {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "login is empty")
	}
	if dto.Password == "" {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "password is empty")
	}

	// login check
	user, err := authSvc.userRepo.GetByLogin(dto.Login)
	if err != nil {
		if err.Error() == "no user found" {
			return nil, echo.NewHTTPError(http.StatusUnauthorized, err.Error())
		}
	}

	//password check (will be submitted in a separate package)
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(dto.Password)); err != nil {
		return nil, echo.NewHTTPError(http.StatusUnauthorized, "incorrect password")
	}

	//token generation
	token, err := auth.GenerateToken(user.Id, user.Login, authSvc.jwtSecretKey)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "token generate error")
	}

	return &model.ResponseData{User: *user, Token: token}, nil

}
