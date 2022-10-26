package helper

import (
	"log"
	"os"
	"time"

	"sleekflow/models"
	"sleekflow/utils/errors"

	"github.com/golang-jwt/jwt/v4"
)

type authCustomClaims struct {
	UserID uint64 `json:"user_id"`

	jwt.RegisteredClaims
}

func GenerateToken(user *models.User, expTime, issuedAt time.Time) (string, error) {
	claims := &authCustomClaims{
		user.ID,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expTime),
			IssuedAt:  jwt.NewNumericDate(issuedAt),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return "", err
	}
	return t, nil
}

func ParseJwt(tokenStr string) (*authCustomClaims, error) {
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
		if t.Method.Alg() != jwt.SigningMethodHS512.Alg() {
			customError := errors.ErrUnauthorized
			customError.Message = "Invalid Token Signing Method"
			return nil, customError
		}
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		customError := errors.ErrUnauthorized
		customError.Message = "Invalid Token"
		return nil, customError
	}

	return &authCustomClaims{
		UserID: uint64(claims["user_id"].(float64)),
	}, nil
}

func ValidateToken(token string) bool {
	tokenJwt, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})
	if err != nil {
		log.Println("err-invalid-token: ", err)
		return false
	}
	return tokenJwt.Valid
}
