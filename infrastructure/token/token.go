package token

import (
	"fmt"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTToken struct {
	secretKey string
}

func NewJWTToken(secretKey string) (*JWTToken, error) {
	if strings.TrimSpace(secretKey) == "" {
		return nil, fmt.Errorf("SecretKey must not empty")
	}

	return &JWTToken{
		secretKey: secretKey,
	}, nil
}

func (r *JWTToken) CreateToken(content string) (string, error) {
	var err error
	//Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["content"] = content
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(r.secretKey))
	if err != nil {
		return "", err
	}
	return token, nil
}

func (r *JWTToken) VerifyToken(tokenString string) (*jwt.Token, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(r.secretKey), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}
