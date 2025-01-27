package utils

import (
	config "arno/configs"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"log"
)

type Claim struct {
	Username string `json:"username"`
	ID       int
	Role     string
}

func TokenParse(tokenString string) (jwt.MapClaims, error) {
	log.Println("Парсинг токена:", tokenString)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		log.Println("Алгоритм токена:", token.Header["alg"])
		log.Println("Секретный ключ:", config.DBConfig.Token.Secret)

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			log.Printf("unexpected signing method: %v", token.Header["alg"])
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.DBConfig.Token.Secret), nil
	})
	if err != nil {
		log.Println("Ошибка парсинга токена:", err)
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	mapClaims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		log.Println("Токен недействителен:", mapClaims)
		return nil, fmt.Errorf("Неправильный токен")
	}

	log.Println("Успешный парсинг токена:", mapClaims)
	return mapClaims, nil
}
