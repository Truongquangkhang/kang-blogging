package jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func CreateAccessToken(userId string, secretKey string, expireHours int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"userId": userId,
			"exp":    time.Now().Add(time.Duration(expireHours) * time.Hour).Unix(),
		})

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func CreateRefreshToken(secretKey string, expireHours int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"exp": time.Now().Add(time.Duration(expireHours) * time.Hour).Unix(),
		})

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func VerifyToken(tokenString string, secretKey string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}

//
//func (j *jwtService) GenerateToken(UserID string) string {
//	claims := &jwtCustomClaim{
//		UserID,
//		jwt.StandardClaims{
//			ExpiresAt: time.Now().AddDate(1, 0, 0).Unix(),
//			Issuer:    j.issuer,
//			IssuedAt:  time.Now().Unix(),
//		},
//	}
//	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
//	t, err := token.SignedString([]byte(j.secretKey))
//	if err != nil {
//		panic(err)
//	}
//	return t
//}
