package utilities

import (
	"github.com/gofrs/uuid"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const TOKEN_EXP = time.Hour * 24
const SECRET_KEY = "supersecretkey"

type (
	TokenService interface {
		GenerateUserToken(id uuid.UUID, email string) (string, error)
	}

	tokenService struct{}
)

func NewTokenService() TokenService {
	return &tokenService{}
}

type Claims struct {
	jwt.RegisteredClaims
	UserID    uuid.UUID
	UserEmail string
}

func (ts *tokenService) GenerateUserToken(id uuid.UUID, email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			// когда создан токен
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TOKEN_EXP)),
		},
		// собственное утверждение
		UserID:    id,
		UserEmail: email,
	})

	return token.SignedString([]byte(SECRET_KEY))
}
