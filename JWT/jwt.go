package JWT

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type MyClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

var mySigningKey = []byte("secret")

func GenToken(username string) (string, error) {
	claims := MyClaims{
		username,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			Issuer:    "blog-back",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(mySigningKey)
}
func ParseToken(tokenString string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return mySigningKey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
