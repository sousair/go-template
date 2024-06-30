package usecase

import (
	"github.com/sousair/go-template/internal/core"
	"github.com/sousair/go-template/internal/infra/cipher"
	"github.com/sousair/go-template/internal/infra/database"
	"github.com/sousair/go-template/internal/infra/token"
)

type Dependencies struct {
	UserRepo *database.Repository[core.User]
	Cipher   cipher.Cipher
	Token    token.Token[TokenPayload]
}

type Usecase struct {
	deps *Dependencies
}

func New(deps *Dependencies) *Usecase {
	return &Usecase{deps}
}
