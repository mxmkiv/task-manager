package handler

import (
	"net/http"
	"task-manager/internal/auth"
	"task-manager/internal/model"

	"github.com/labstack/echo/v5"
)

type RequestData struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type ResponseData struct {
	User  model.User `json:"user"`
	Token string     `json:"token"`
}

func RegisterHandler(ctx *echo.Context) error {

	var getUserData RequestData
	err := ctx.Bind(&getUserData)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid json format")
	}

	// validation
	if len(getUserData.Login) < 3 || len(getUserData.Login) > 20 {
		return echo.NewHTTPError(http.StatusBadRequest, "login must be between 3 and 20 characters")
	}
	if len(getUserData.Password) < 5 {
		return echo.NewHTTPError(http.StatusBadRequest, "password must be at least 5 characters")
	}

	user, err := model.NewUser(getUserData.Login, getUserData.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to create new user")
	}

	token, err := auth.GenerateToken(user.Id, user.Login)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "token generate error")
	}
	return ctx.JSON(http.StatusOK, ResponseData{User: *user, Token: token})
}

func TestHandler(ctx *echo.Context) error {
	return ctx.String(http.StatusOK, "hello world")
}
