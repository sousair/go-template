package usecase

import (
	"context"
	"errors"

	"github.com/sousair/go-template/internal/core"
	"github.com/sousair/go-template/internal/infra/database"
)

func (u Usecase) CreateUser(ctx context.Context, req *core.CreateUserRequest) (*core.User, error) {
	alreadyExists, err := u.deps.UserRepo.FindBy(ctx, &core.User{Email: req.Email})
	if err != nil {
		if !errors.Is(err, database.ErrNotFound) {
			return nil, err
		}
	}

	if alreadyExists != nil {
		return nil, ErrUserAlreadyRegistered
	}

	userPassword, err := u.deps.Cipher.Encrypt(req.Password)
	if err != nil {
		return nil, err
	}

	newUser := &core.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: userPassword,
	}

	newUser, err = u.deps.UserRepo.Create(ctx, newUser)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}
