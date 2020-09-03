package platform

import (
	types "github.com/noartem/godi-example"
	"github.com/noartem/godi-example/ent/user"
	"github.com/noartem/godi-example/pkg/api/auth/service"
	"github.com/noartem/godi-example/pkg/util/ent"
)

type userDB struct {
	ent *ent.ClientWithCtx
}

func NewUserDB(ent *ent.ClientWithCtx) service.IUserDB {
	return &userDB{ ent: ent }
}

func (udb *userDB) CreateUser(user types.User) (*types.User, error) {
	u, err := udb.ent.DB.User.
		Create().
		SetEmail(user.Email).
		SetPassword(user.Password).
		SetName(user.Name).
		Save(udb.ent.Ctx)
	if err != nil {
		return nil, err
	}

	return &types.User{
		Name:  u.Name,
		Email: u.Email,
		Password: u.Password,
	}, nil
}

func (udb *userDB) FindByEmail(email string) (*types.User, error) {
	u, err := udb.ent.DB.User.
		Query().
		Where(
			user.EmailEQ(email),
		).
		Only(udb.ent.Ctx)
	if err != nil {
		return nil, err
	}

	return &types.User{
		Name:  u.Name,
		Email: u.Email,
		Password: u.Password,
	}, nil
}
