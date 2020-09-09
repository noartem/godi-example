package service

import (
	"fmt"
	types "github.com/noartem/godi-example"
	"github.com/noartem/godi-example/pkg/util"
	"github.com/noartem/godi-example/pkg/util/config"
	"github.com/noartem/godi-example/pkg/util/jwt"
	"time"
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
	hasher *util.Hasher
}

func NewService(udb IUserDB, jwt jwt.JWT, config *config.Config, hasher *util.Hasher) IService {
	service := &Service{udb: udb, jwt: jwt, config: config, hasher: hasher}

	return service
}

func (service *Service) Authenticate(email string, password string) (*types.AuthToken, error) {
	user, err := service.udb.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	passwordIsCorrect, err := service.hasher.Compare(password, user.Password)
	if err != nil {
		return nil, err
	}

	if !passwordIsCorrect {
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

func (service *Service) parseToken(token string) (string, time.Time, error) {
	claims, err := service.jwt.ParseToken(token)
	if err != nil {
		return "", time.Time{}, fmt.Errorf("parse token: %v", err)
	}

	emailI := claims["e"]
	if emailI == nil {
		return "", time.Time{}, fmt.Errorf("invalid claims email")
	}

	email, ok := emailI.(string)
	if !ok {
		return "", time.Time{}, fmt.Errorf("invalid claims email")
	}

	expI := claims["exp"]
	if expI == nil {
		return "", time.Time{}, fmt.Errorf("invalid claims expiration")
	}

	exp, ok := expI.(float64)
	if !ok {
		return "", time.Time{}, fmt.Errorf("invalid claims expiration")
	}

	expTime := time.Unix(int64(exp), 0)

	return email, expTime, nil
}

func (service *Service) Refresh(token string) (*types.AuthToken, error) {
	email, exp, err := service.parseToken(token)
	if err != nil {
		return nil, err
	}

	if exp.Add(time.Minute * time.Duration(service.config.JWT.RefreshTTL)).Before(time.Now()) {
		return nil, fmt.Errorf("token is expired")
	}

	user, err := service.udb.FindByEmail(email)
	if err != nil {
		return nil, err
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

func (service *Service) Create(email string, password string, name string) (*types.User, error) {
	hashedPassword, err := service.hasher.Hash(password)
	if err != nil {
		return nil, fmt.Errorf("hash passsword: %v", err)
	}

	user, err := service.udb.CreateUser(types.User{Email: email, Password: hashedPassword, Name: name})
	if err != nil {
		return nil, fmt.Errorf("platfrom: %v", err)
	}

	return user, nil
}
