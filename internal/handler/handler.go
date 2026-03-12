package handler

import (
	"net/http"
	"strconv"
	"task-manager/internal/model"
	"task-manager/internal/service"

	"github.com/labstack/echo/v5"
)

type Handler struct {
	authSrv *service.AuthService
	userSrv *service.UserService
}

func NewHandler(authSvc *service.AuthService, userSvc *service.UserService) *Handler {
	return &Handler{
		authSrv: authSvc,
		userSrv: userSvc,
	}
}

func (h *Handler) RegisterHandler(ctx *echo.Context) error {

	var getUserData model.RequestData
	err := ctx.Bind(&getUserData)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid json format")
	}

	response, err := h.authSrv.Register(&getUserData)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, response)
}

func (h *Handler) LoginHabdler(ctx *echo.Context) error {

	var getUserData model.RequestData
	err := ctx.Bind(&getUserData)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid json format")
	}

	response, err := h.authSrv.Login(&getUserData)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, response)

}

func (h *Handler) TestHandler(ctx *echo.Context) error {
	return ctx.String(http.StatusOK, "hello world")
}

func (h *Handler) AllUsersHandler(ctx *echo.Context) error {

	response, err := h.userSrv.GetAllUsers()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, response)

}

func (h *Handler) UserById(ctx *echo.Context) error {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "incorrect id")
	}

	response, err := h.userSrv.GetUserById(id)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, response)
}

func (h *Handler) CreateAdmin(ctx *echo.Context) error {

	var getUserData model.RequestData
	err := ctx.Bind(&getUserData)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid json format")
	}

	response, err := h.authSrv.AdminRegister(&getUserData)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, response)

}
