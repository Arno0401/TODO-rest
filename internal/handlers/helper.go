package handler

import (
	"arno/internal/models"
	"arno/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"slices"
	"unicode"
)

func sendResponse(c *gin.Context, code int, message string) {
	resp := models.Response{
		Code:    code,
		Message: message,
	}
	c.JSON(http.StatusOK, resp)
}

func sendSuccessResponse(c *gin.Context, resp models.Response) {
	c.JSON(http.StatusOK, resp)
}

func validateLogin(login string) bool {
	return len(login) >= 5
}
func validateLatinLogin(login string) bool {

	for _, char := range login {
		if !unicode.Is(unicode.Latin, char) && !unicode.IsDigit(char) {
			return false
		}
	}
	return len(login) >= 5
}

func validatePassword(password string) bool {
	if len(password) < 8 {
		return false
	}

	HasNumber := false
	HasSymbol := false
	Haslett := false

	for _, char := range password {
		if unicode.IsDigit(char) {
			HasNumber = true
		} else if unicode.IsLetter(char) {
			Haslett = true
		} else if isSymbol(char) {
			HasSymbol = true
		}
	}

	return Haslett && HasSymbol && HasNumber
}

func isSymbol(char rune) bool {
	symbols := []rune{'!', '@', '#', '$', '%', '^', '&', '*'}
	return slices.Contains(symbols, char)
}

func parseAndValidateToken(c *gin.Context) (int, error) {
	// Получаем токен из заголовка Authorization
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		return 0, fmt.Errorf("Токен отсутствует")
	}

	// Убираем "Bearer ", если оно есть
	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		tokenString = tokenString[7:]
	}

	// Парсим токен
	mapClaims, err := utils.TokenParse(tokenString)
	if err != nil {
		return 0, fmt.Errorf("Неверный токен")
	}

	// Проверяем, что это access-токен
	if mapClaims["type"] != "access" {
		return 0, fmt.Errorf("Токен не является access-токеном")
	}

	// Получаем ID пользователя из токена
	userID, ok := mapClaims["id"].(float64)
	if !ok {
		return 0, fmt.Errorf("Ошибка получения ID пользователя")
	}

	return int(userID), nil
}
