package jwt

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type UserClaims struct {
	jwt.RegisteredClaims
	UserId string `json:"userId"`
	PhoneNumber string `json:"phoneNumber"`
}

func SignJWT(userId string, phoneNumber string) (string, error) {
	exp := time.Now().Add(time.Hour * 8)
	claims := UserClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt: jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(exp),
			Issuer: "Eniqilo Store",
		},
		UserId: userId,
		PhoneNumber: phoneNumber,
	}
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)
	signedToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil{
		return "", err
	}
	return signedToken, nil
}

func ParseToken(jwtToken string) (string, error) {
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		if _, OK := token.Method.(*jwt.SigningMethodHMAC); !OK {
			return nil, errors.New("bad signed method received")
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		log.Fatal(err)
		return "", err
	}

	parsedToken, OK := token.Claims.(jwt.MapClaims)
	if !OK {
		return "", errors.New("unable to parse claims")
	}
	id := fmt.Sprint(parsedToken["userId"])
	return id, nil
}