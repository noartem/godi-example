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

func (controller *authController) Middlewares() []echo.MiddlewareFunc {
	return []echo.MiddlewareFunc{}
}

func (controller *authController) Register(e *echo.Echo) error {
	// swagger:route POST /auth/login auth login
	// Logs in user by username and Password.
	// responses:
	//  200: loginResp
	//  400: errMsg
	//  401: errMsg
	// 	403: err
	//  404: errMsg
	//  500: err
	e.POST("/auth/login", controller.login)

	// swagger:route POST /auth/register auth register
	// Register new user.
	// responses:
	//  200: registerResp
	//  400: errMsg
	//  401: errMsg
	// 	403: err
	//  404: errMsg
	//  500: err
	e.POST("/auth/register", controller.register)

	// swagger:operation POST /refresh/{token} auth refresh
	// ---
	// summary: Refreshes jwt token.
	// parameters:
	// - Name: token
	//   in: path
	//   description: refresh token
	//   type: string
	//   required: true
	// responses:
	//   "200":
	//     "$ref": "#/responses/refreshResp"
	//   "400":
	//     "$ref": "#/responses/errMsg"
	//   "401":
	//     "$ref": "#/responses/err"
	//   "500":
	//     "$ref": "#/responses/err"
	e.POST("/refresh", controller.refresh)

	return nil
}

type credentials struct {
	Email    string `json:"Email" validate:"required,email"`
	Password string `json:"Password" validate:"required,password"`
}

func (controller *authController) login(ctx echo.Context) error {
	cred := new(credentials)
	if err := ctx.Bind(cred); err != nil {
		return err
	}

	if err := ctx.Validate(cred); err != nil {
		return err
	}

	token, err := controller.service.Authenticate(cred.Email, cred.Password)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, token)
}

type newUser struct {
	Email    string `json:"Email" validate:"required,email"`
	Password string `json:"Password" validate:"required,password"`
	Name     string `json:"Name" validate:"required"`
}

func (controller *authController) register(ctx echo.Context) error {
	newUser := new(newUser)
	if err := ctx.Bind(newUser); err != nil {
		return err
	}

	if err := ctx.Validate(newUser); err != nil {
		return err
	}

	user, err := controller.service.Create(newUser.Email, newUser.Password, newUser.Name)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, user)
}

func (controller *authController) refresh(ctx echo.Context) error {
	token, err := controller.service.Refresh(ctx.Param("token"))
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, token)
}
