package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type TokenUtil struct {
	Token string
}

func SecretKey(Secret string) TokenUtil {
	return TokenUtil{
		Token: Secret,
	}
}

func (t TokenUtil) NewJWTWtoken(UserID string) (string, error) {

	claims := jwt.MapClaims{
		"id":  UserID,
		"exp": time.Now().Add(24 * 30 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(&jwt.SigningMethodHMAC{}, claims)

	JWT, err := token.SignedString([]byte(t.Token))

	if err != nil {
		return " ", err
	}

	return JWT, nil
}

func (t TokenUtil) ParseJWT(TokenSTring string) (string, error) {
	tokenparse, errs := jwt.Parse(TokenSTring, func(TokenParam *jwt.Token) (interface{}, error) {
		if _, ok := TokenParam.Method.(*jwt.SigningMethodHMAC); !ok {
			return " ", fmt.Errorf("method signing salah %v", TokenParam.Header["alg"])
		}
		return []byte(t.Token), nil
	})

	if errs != nil {
		return " ", errs
	}

	claims, ok := tokenparse.Claims.(jwt.MapClaims)

	if ok && tokenparse.Valid {
		id, ok := claims["id"].(string)

		if !ok {
			return " ", errors.New("gagal mendapatkan ID dari claims %")
		}

		return id, nil
	}
	return " ", errs
}
