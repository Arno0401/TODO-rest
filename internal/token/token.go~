package token

import (
	config "arno/configs"
	"arno/internal/models"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

func CreateToken(user models.Users) (models.Token, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   user.ID,
		"role": user.Role,
		"type": "access",
		"exp":  time.Now().Add(time.Minute * config.DBConfig.Token.Access).Unix(),
	})

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   user.ID,
		"role": user.Role,
		"type": "refresh",
		"exp":  time.Now().Add(time.Hour * config.DBConfig.Token.Refresh).Unix(),
	})

	signedAccessToken, err := accessToken.SignedString([]byte(config.DBConfig.Token.Secret))
	if err != nil {
		return models.Token{}, err
	}

	signedRefreshToken, err := refreshToken.SignedString([]byte(config.DBConfig.Token.Secret))
	if err != nil {
		return models.Token{}, err
	}

	tokens := models.Token{
		AccessToken:  signedAccessToken,
		RefreshToken: signedRefreshToken,
	}

	return tokens, nil
}
