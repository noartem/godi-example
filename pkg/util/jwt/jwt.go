package jwt

import (
	"fmt"
	types "github.com/noartem/godi-example"
	"github.com/noartem/godi-example/pkg/util/config"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWT interface {
	ParseToken(authHeader string) (jwt.MapClaims, error)
	GenerateToken(user *types.User) (string, error)
	GenerateRefreshToken(userId uint) (string, error)
}

// New generates new JWT service necessary for auth middleware
func NewJWT(config *config.Config) (JWT, error) {
	signingMethod := jwt.GetSigningMethod(config.JWT.Algorithm)
	if signingMethod == nil {
		return nil, fmt.Errorf("invalid jwt signing method: %s", config.JWT.Algorithm)
	}

	return &Service{
		key:        []byte(config.JWT.Secret),
		algo:       signingMethod,
		ttl:        time.Duration(config.JWT.TTL) * time.Minute,
		refreshTtl: time.Duration(config.JWT.RefreshTTL) * time.Minute,
	}, nil
}

// Service provides a Json-Web-Token authentication implementation
type Service struct {
	// Secret key used for signing.
	key []byte

	// Duration for which the jwt token is valid.
	ttl time.Duration

	refreshTtl time.Duration

	// JWT signing algorithm
	algo jwt.SigningMethod
}

// ParseToken parses token from Authorization header
func (s *Service) ParseToken(tokenString string) (jwt.MapClaims, error) {
	claims := make(jwt.MapClaims)
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if s.algo != token.Method {
			return nil, fmt.Errorf("invalid token")
		}
		return s.key, nil
	})
	if err != nil {
		return nil, err
	}

	return claims, nil
}

// GenerateToken generates new JWT token and populates it with user data
func (s *Service) GenerateToken(u *types.User) (string, error) {
	return jwt.NewWithClaims(s.algo, jwt.MapClaims{
		"id":  u.ID,
		"n":   u.Name,
		"e":   u.Email,
		"exp": time.Now().Add(s.ttl).Unix(),
	}).SignedString(s.key)
}

// GenerateRefreshToken generates new refresh JWT token
func (s *Service) GenerateRefreshToken(id uint) (string, error) {
	return jwt.NewWithClaims(s.algo, jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(s.ttl).Unix(),
	}).SignedString(s.key)
}
