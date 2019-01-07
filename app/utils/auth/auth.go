package auth

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// TokenClaim This is the cliam object which gets parsed from the authorization header
type TokenClaim struct {
	*jwt.StandardClaims
	BaseUser
}

// BaseUser ...
type BaseUser struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

// CreateToken ...
func CreateToken(id uint, username string, name string) (string, int64) {
	expiresAt := time.Now().Add(time.Hour * 24 * 30).Unix()

	token := jwt.New(jwt.SigningMethodHS256)

	token.Claims = &TokenClaim{
		&jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
		BaseUser{id, username, name},
	}

	tokenString, error := token.SignedString([]byte("secret"))
	if error != nil {
		fmt.Println(error)
	}
	return tokenString, expiresAt
}
