package utils

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

var secretKey = "123456789"

func GenerateToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		username: username,
	})
	fmt.Println("secret key", secretKey)
	tokenString, err := token.SignedString("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJuYmYiOjE0NDQ0Nzg0MDB9.u1riaD1rW97opCoAuRCTy4w58Br-Zk-bh7vLiRIsrpU")
	fmt.Println("token ", tokenString)
	if err != nil {
		return "", nil
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok || token.Valid {
		return "", err
	}

	username := claims["username"].(string)
	return username, nil

}
