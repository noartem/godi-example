package service

import (
	"fmt"
	types "github.com/noartem/godi-example"
	"github.com/noartem/godi-example/pkg/util"
	"github.com/noartem/godi-example/pkg/util/config"
	"github.com/noartem/godi-example/pkg/util/jwt"
	"log"
)

type IService interface {
	Create(email string, password string, name string) (*types.User, error)
	Authenticate(email string, password string) (*types.AuthToken, error)
	Refresh(token string) (*types.AuthToken, error)
}

type IUserDB interface {
	CreateUser(types.User) (*types.User, error)
	FindByEmail(string) (*types.User, error)
}

type Service struct {
	udb    IUserDB
	config *config.Config
	jwt    jwt.JWT
}

func NewService(udb IUserDB, jwt jwt.JWT) IService {
	service := &Service{udb: udb, jwt: jwt}

	return service
}

func (service *Service) Authenticate(email string, password string) (*types.AuthToken, error) {
	user, err := service.udb.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	passwordIsCorrect, err := util.CompareHash(password, user.Password)
	if err != nil {
		return nil, err
	}

	if passwordIsCorrect {
		return nil, fmt.Errorf("invalid password")
	}

	authToken, err := service.jwt.GenerateToken(user)
	if err != nil {
		return nil, fmt.Errorf("generate token: %v", err)
	}

	refreshToken, err := service.jwt.GenerateRefreshToken(user.ID)
	if err != nil {
		return nil, fmt.Errorf("generate refresh token: %v", err)
	}

	return &types.AuthToken{
		Token:        authToken,
		RefreshToken: refreshToken,
	}, nil
}

func (service *Service) Refresh(tokenString string) (*types.AuthToken, error) {
	token, err := service.jwt.ParseToken(tokenString)
	if err != nil {
		return nil, fmt.Errorf("parse token: %v", err)
	}

	log.Println(token.Header)

	return &types.AuthToken{}, nil
}

func (service *Service) Create(email string, password string, name string) (*types.User, error) {
	hashedPassword, err := util.HashPassword(password)
	if err != nil {
		return nil, fmt.Errorf("hash passsword: %v", err)
	}

	user, err := service.udb.CreateUser(types.User{Email: email, Password: hashedPassword, Name: name})
	if err != nil {
		return nil, fmt.Errorf("platfrom: %v", err)
	}

	return user, nil
}
