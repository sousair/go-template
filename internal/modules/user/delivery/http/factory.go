package http

import (
	"github.com/labstack/echo/v4"
	"github.com/sousair/go-template/internal/core"
	"github.com/sousair/go-template/internal/infra/httpx"
	"github.com/sousair/go-template/internal/modules/user/usecase"
)

type Dependencies struct {
	Server      *echo.Echo
	UserUsecase *usecase.Usecase
}

type Handler struct {
	deps *Dependencies
}

func New(deps *Dependencies) {
	h := &Handler{deps}

	deps.Server.POST("/users", h.createUser)
	deps.Server.POST("/users/login", h.userLogin)
}

func (h Handler) createUser(e echo.Context) error {
	req, err := httpx.NewRequest[core.CreateUserRequest](e)
	if err != nil {
		return httpx.NewBadRequestResponse(e, err)
	}

	user, err := h.deps.UserUsecase.CreateUser(e.Request().Context(), req)
	if err != nil {
		return httpx.NewInternalServerErrorResponse(e, err)
	}

	return httpx.NewCreatedResponse(e, user)
}

func (h Handler) userLogin(e echo.Context) error {
	req, err := httpx.NewRequest[core.LoginRequest](e)
	if err != nil {
		return httpx.NewBadRequestResponse(e, err)
	}

	res, err := h.deps.UserUsecase.Login(e.Request().Context(), req)
	if err != nil {
		return httpx.NewInternalServerErrorResponse(e, err)
	}

	return httpx.NewCreatedResponse(e, res)
}
