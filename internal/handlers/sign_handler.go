package handler

import (
	"arno/internal/models"
	"arno/internal/token"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (h *Handler) SignUpHandler(c *gin.Context) {
	var req models.SignUpRequest
	var resp models.Response

	if err := c.ShouldBindJSON(&req); err != nil {
		sendResponse(c, http.StatusBadRequest, "Плохой запрос")
		return
	}
	if !validateLogin(req.Login) {
		sendResponse(c, http.StatusBadRequest, "Логин пользователя должно состоять от 5 символов")
		return
	}

	if !validateLatinLogin(req.Login) {
		sendResponse(c, http.StatusBadRequest, "Логин должен состоять из латинских букв и цифр")
		return
	}

	if !validatePassword(req.Password) {
		sendResponse(c, http.StatusBadRequest, "Пароль должен состоять от 8 символов, из латинских букв, цифр и специальных символов")
		return
	}

	err := h.rep.SignUpUser(req.UserName, req.Login, req.Password)

	if err != nil {
		sendResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	resp.Code = http.StatusOK
	resp.Message = "Пользователь успешно зарегистрирован"
	sendSuccessResponse(c, resp)
}

func (h *Handler) SignInHandler(c *gin.Context) {
	var req models.LoginRequest
	var resp models.Response

	if err := c.ShouldBindJSON(&req); err != nil {
		sendResponse(c, http.StatusBadRequest, "Плохой запрос")
		return
	}

	exists, err := h.rep.IsExistingUser(req.Login, req.Password)
	log.Printf("exists: %v, err: %v", exists, err)
	log.Printf("req.Login: %v, req.Password: %v", req.Login, req.Password)
	if err != nil {
		sendResponse(c, http.StatusInternalServerError, "Внутренняя ошибка сервера")
		return
	}

	if !exists {
		sendResponse(c, http.StatusBadRequest, "Неверный логин или пароль")
		return
	}
	user, err := h.rep.GetUser(req.Login)
	if err != nil {
		sendResponse(c, http.StatusInternalServerError, "Внутренняя ошибка сервера")
		return
	}

	tokens, err := token.CreateToken(*user)
	if err != nil {
		sendResponse(c, http.StatusInternalServerError, "Внутренняя ошибка сервера")
		return
	}

	resp.Code = http.StatusOK
	resp.Message = "Вход выполнен успешно"
	resp.Token = &tokens
	sendSuccessResponse(c, resp)
}

func (h *Handler) Profile(c *gin.Context) {

	userID, err := parseAndValidateToken(c)
	if err != nil {
		sendResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	user, err := h.rep.GetUserByID(userID)
	if err != nil {
		sendResponse(c, http.StatusInternalServerError, "Ошибка при получении данных пользователя")
		return
	}

	resp := models.Response{
		Code:    http.StatusOK,
		Message: "Профиль успешно получен",
		Profile: &models.ProfileResponse{
			ID:    user.ID,
			Name:  user.UserName,
			Login: user.Login,
			Role:  user.Role,
		},
	}
	sendSuccessResponse(c, resp)
}

func (h *Handler) ChangePassword(c *gin.Context) {
	var req models.ChangePassRequest
	var resp models.Response

	if err := c.ShouldBindJSON(&req); err != nil {
		sendResponse(c, http.StatusBadRequest, "Плохой запрос")
		return
	}

	userID, err := parseAndValidateToken(c)
	if err != nil {
		sendResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	err = h.rep.ChangePassword(userID, req.OldPass, req.NewPass)
	if err != nil {
		if err.Error() == "incorrect password" {
			sendResponse(c, http.StatusUnauthorized, "Старый пароль неверен")
			return
		}
		sendResponse(c, http.StatusInternalServerError, "Ошибка при обновлении пароля")
		return
	}

	resp = models.Response{
		Code:    http.StatusOK,
		Message: "Пароль успешно обновлен",
	}
	sendSuccessResponse(c, resp)
}
