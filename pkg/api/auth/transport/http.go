package transport

import (
	"github.com/labstack/echo"
	types "github.com/noartem/godi-example"
	auth "github.com/noartem/godi-example/pkg/api/auth/service"
	"net/http"
)

type authController struct {
	service auth.IService
}

func NewAuth(service auth.IService) types.IController {
	return &authController{
		service: service,
	}
}

func (controller *authController) Name() string {
	return "/auth"
}

func (controller *authController) Middlewares() []echo.MiddlewareFunc {
	return []echo.MiddlewareFunc{}
}

func (controller *authController) Register(g *echo.Group) error {
	g.GET("/login", controller.login)
	g.GET("/refresh", controller.refresh)

	return nil
}

type credentials struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (controller *authController) login(ctx echo.Context) error {
	cred := new(credentials)
	if err := ctx.Bind(cred); err != nil {
		return err
	}

	token, err := controller.service.Authenticate(cred.Email, cred.Password)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, token)
}

func (controller *authController) refresh(ctx echo.Context) error {
	token, err := controller.service.Refresh(ctx.Param("token"))
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, token)
}
