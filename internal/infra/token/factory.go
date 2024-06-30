package token

import (
	jwt "github.com/dgrijalva/jwt-go"
)

type Token[T any] interface {
	Generate(payload *T) (string, error)
	Validate(token string) (*T, error)
}

type TokenPayload[T any] struct {
	jwt.StandardClaims
	Payload *T
}

type JWT[T any] struct {
	Secret string
}

var _ Token[any] = (*JWT[any])(nil)

func New[T any](secret string) Token[T] {
	return &JWT[T]{Secret: secret}
}

func (j JWT[T]) Generate(payload *T) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &TokenPayload[T]{Payload: payload})
	return token.SignedString([]byte(j.Secret))
}

func (j JWT[T]) Validate(token string) (*T, error) {
	claims := &TokenPayload[T]{}

	t, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.Secret), nil
	})
	if err != nil {
		return nil, err
	}

	if !t.Valid {
		return nil, ErrInvalidToken
	}

	return claims.Payload, nil
}
