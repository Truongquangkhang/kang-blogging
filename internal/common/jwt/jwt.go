package jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
	"kang-blogging/internal/common/errors"
	"os"
	"strings"
	"time"
)

func CreateAccessToken(userId string, role string, secretKey string, expireHours int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"userId": userId,
			"role":   role,
			"exp":    time.Now().Add(time.Duration(expireHours) * time.Hour).Unix(),
		})

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func CreateRefreshToken(userId string, role string, secretKey string, expireHours int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"userId": userId,
			"role":   role,
			"exp":    time.Now().Add(time.Duration(expireHours) * time.Hour).Unix(),
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

func GetIDAndRoleFromRequest(ctx context.Context) (*string, *string, error) {
	secretKey := os.Getenv("JWT_SECRET_KEY")
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, nil, errors.NewAuthorizationError("get an error when get data from header")
	}
	authorizations := md.Get("authorization")
	if len(authorizations) < 1 {
		return nil, nil, errors.NewAuthorizationError("bearer token not found")
	}
	bearerToken := strings.TrimPrefix(authorizations[0], "Bearer ")

	if err := VerifyToken(bearerToken, secretKey); err != nil {
		return nil, nil, errors.NewAuthorizationDefaultError()
	}
	id, role, err := GetIDAndRoleFromJwtToken(bearerToken, secretKey)
	if err != nil {
		return nil, nil, errors.NewAuthorizationDefaultError()
	}
	return &id, &role, err
}

func GetIDAndRoleFromJwtToken(jwtToken string, secretKey string) (string, string, error) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(jwtToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return "", "", err
	}
	return claims["userId"].(string), claims["role"].(string), nil
}
