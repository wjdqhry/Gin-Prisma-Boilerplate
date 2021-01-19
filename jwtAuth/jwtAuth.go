package jwtAuth

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

func Login(userid string) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["id"] = userid
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)

	token, err := at.SignedString([]byte("your-secret-code"))
	if err != nil {
		return "", err
	} else {
		return token, nil
	}
}
func ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("Invalid token", token.Header["alg"])

		}

		return []byte("your-secret-code"), nil
	})

}
