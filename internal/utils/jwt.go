package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/mbilarusdev/fool_base/v2/log"
)

type CustomClaim struct {
	jwt.RegisteredClaims
	Name string `json:"username"`
}

func CreateJWT(playerID string, name string) (string, error) {
	claims := CustomClaim{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "fool_auth_service",
			Subject:   playerID,
			ExpiresAt: &jwt.NumericDate{Time: time.Now().UTC().Add(time.Hour * 24 * 30 * 6)},
			IssuedAt:  &jwt.NumericDate{Time: time.Now().UTC()},
		},
		Name: name,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	if tokenString, err := token.SignedString([]byte(Conf.Secret)); err != nil {
		return "", log.Err(err, "Failed to create jwt")
	} else {
		return tokenString, nil
	}
}

func CheckJwt(tokenString string) bool {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, log.Err(
				fmt.Errorf("Unexpected signing method: %v", token.Header["alg"]),
				"",
			)
		}
		return []byte(Conf.Secret), nil
	})

	if err != nil {
		log.Err(err, "Error parsing token")
		return false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		exp, _ := claims["exp"].(float64)
		currentTime := float64(time.Now().Unix())

		if currentTime > exp {
			log.Warn("Token expired")
			return false
		}
		return true
	} else {
		log.Err(errors.New("Invalid token or claims"), "")
		return false
	}
}
