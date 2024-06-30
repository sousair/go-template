package main

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	"github.com/sousair/go-template/internal/core"
	"github.com/sousair/go-template/internal/infra/cipher"
	"github.com/sousair/go-template/internal/infra/database"
	"github.com/sousair/go-template/internal/infra/token"
	"github.com/sousair/go-template/internal/modules/user/delivery/http"
	"github.com/sousair/go-template/internal/modules/user/usecase"
	"golang.org/x/crypto/bcrypt"
)

func panicError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	var (
		connectionUrl = os.Getenv("DATABASE_URL")
		jwtSecret     = os.Getenv("JWT_SECRET")
	)

	db, err := database.NewPostgres(connectionUrl)
	panicError(err)

	userRepo, err := database.NewRepository[core.User](db)
	panicError(err)

	cipher := cipher.NewBcryptCipher(bcrypt.DefaultCost)
	token := token.New[usecase.TokenPayload](jwtSecret)

	userUsecase := usecase.New(&usecase.Dependencies{
		UserRepo: userRepo,
		Cipher:   cipher,
		Token:    token,
	})

	echoServer := echo.New()

	http.New(&http.Dependencies{
		Server:      echoServer,
		UserUsecase: userUsecase,
	})

	echoServer.Logger.Fatal(
		echoServer.Start(":" + os.Getenv("PORT")),
	)
}
