package helpers

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JwtCustomClaimsAccessToken struct {
	User interface{} `json:"user"`
	Name string
	jwt.RegisteredClaims
}

type JwtCustomClaimsRefreshToken struct {
	User interface{} `json:"user"`
	Name string
	jwt.RegisteredClaims
}

type GetJwtToken struct {
	User         interface{} `json:"user"`
	AccessToken  string      `json:"access_token"`
	RefreshToken string      `json:"refresh_token"`
}

func GetToken(User interface{}) (GetJwtToken, error) {
	var token GetJwtToken
	secret := os.Getenv("jwt_secret")
	AccessTokenExpiredTime := 1
	RefreshTokenExpiredTime := 168

	ClaimAccesToken := &JwtCustomClaimsAccessToken{
		User,
		"access_token",
		jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(AccessTokenExpiredTime))),
		},
	}

	ClaimsRefreshToken := &JwtCustomClaimsRefreshToken{
		User,
		"refresh_token",
		jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(RefreshTokenExpiredTime))),
		},
	}

	AccessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, ClaimAccesToken)
	RefreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, ClaimsRefreshToken)

	SignedStringAccessToken, err := AccessToken.SignedString([]byte(secret))
	if err != nil {
		return GetJwtToken{}, err
	}
	SignedStringRefreshToken, err := RefreshToken.SignedString([]byte(secret))
	if err != nil {
		return GetJwtToken{}, err
	}
	token.User = &User
	token.AccessToken = SignedStringAccessToken
	token.RefreshToken = SignedStringRefreshToken
	return token, nil
}
