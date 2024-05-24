package shared

import (
	"awesomeProject/utilities"
	"github.com/gofrs/uuid"
	"github.com/golang-jwt/jwt/v5"
	"strings"
)

type Claims struct {
	jwt.RegisteredClaims
	UserID    uuid.UUID
	UserEmail string
}

func GetUserIdFromAuthHeader(authHeader string) (uuid.UUID, bool) {
	if authHeader == "" {
		return uuid.Nil, false
	}

	_, err := ParseTokenFromHeader(authHeader)
	if err != nil {
		return uuid.Nil, false
	}

	claims := &Claims{}

	tokenString := strings.Split(authHeader, " ")

	jwt.ParseWithClaims(tokenString[1], claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(utilities.SECRET_KEY), nil
	})

	return claims.UserID, true

}

func ParseTokenFromHeader(authTokenHeader string) (*jwt.Token, error) {
	tokenParts := strings.Split(authTokenHeader, " ")

	if len(tokenParts) != 2 {
		return nil, ErrInvalidTokenHeader
	}

	return jwt.Parse(tokenParts[1], func(token *jwt.Token) (interface{}, error) {
		if _, tokenIsValid := token.Method.(*jwt.SigningMethodHMAC); !tokenIsValid {
			return nil, ErrInvalidSigningMethod
		}
		return []byte(utilities.SECRET_KEY), nil
	})
}
