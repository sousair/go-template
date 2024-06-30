package usecase

import (
	"context"

	"github.com/sousair/go-template/internal/core"
)

type TokenPayload struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
}

type LoginResponse struct {
	Token   string        `json:"token"`
	Payload *TokenPayload `json:"payload"`
}

func (u Usecase) Login(ctx context.Context, req *core.LoginRequest) (*LoginResponse, error) {
	user, err := u.deps.UserRepo.FindBy(ctx, &core.User{Email: req.Email})
	if err != nil {
		return nil, err
	}

	if ok, err := u.deps.Cipher.Compare(req.Password, user.Password); err != nil || !ok {
		return nil, ErrUserInvalidCredentials
	}

	payload := &TokenPayload{
		UserID: user.ID,
		Email:  user.Email,
	}

	token, err := u.deps.Token.Generate(payload)
	if err != nil {
		return nil, err
	}

	return &LoginResponse{
		Token:   token,
		Payload: payload,
	}, nil
}
