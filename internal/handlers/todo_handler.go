package handler

import (
	"arno/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) CreateTodos(c *gin.Context) {
	var req models.TodoRequest
	var resp models.Response

	if err := c.ShouldBindJSON(&req); err != nil {
		sendResponse(c, http.StatusBadRequest, "Плохой запрос")
		return
	}
	userID := c.GetInt("user_id")
	if err := h.rep.CreateTodo(&req, userID); err != nil {
		sendResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	resp.Code = http.StatusOK
	resp.Message = "Задача добавлена"
	sendSuccessResponse(c, resp)
}

func (h *Handler) GetTodos(c *gin.Context) {
	userID := c.GetInt("user_id")

	todos, err := h.rep.GetTodosByUserID(userID)
	if err != nil {
		sendResponse(c, http.StatusInternalServerError, "Ошибка при получении списка задач")
		return
	}

	c.JSON(http.StatusOK, todos)
}

func (h *Handler) UpdateTodos(c *gin.Context) {
	var resp models.Response
	var req models.TodoRequest

	todoID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		sendResponse(c, http.StatusBadRequest, "ID задачи не найден")
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		sendResponse(c, http.StatusBadRequest, "Плохой запрос")
		return
	}
	userID := c.GetInt("user_id")

	task, err := h.rep.GetTaskByID(todoID)
	if err != nil {
		sendResponse(c, http.StatusNotFound, "Задача не найдена")
		return
	}

	if task.UserID != userID {
		sendResponse(c, http.StatusForbidden, "У вас нет прав на редактирование этой задачи")
		return
	}

	if err := h.rep.Update(todoID, &req); err != nil {
		sendResponse(c, http.StatusInternalServerError, "Ошибка обновления задачи")
		return
	}

	resp.Code = http.StatusOK
	resp.Message = "Задача успешно обновлена"
	sendSuccessResponse(c, resp)
}

func (h *Handler) DeleteTodo(c *gin.Context) {
	var resp models.Response
	todoID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		sendResponse(c, http.StatusBadRequest, "ID задачи не найден")
		return
	}

	userID := c.GetInt("user_id")

	task, err := h.rep.GetTaskByID(todoID)
	if err != nil {
		sendResponse(c, http.StatusNotFound, "Задача не найдена")
		return
	}

	if task.UserID != userID {
		sendResponse(c, http.StatusForbidden, "У вас нет прав на удалении этой задачи")
		return
	}

	if err := h.rep.Delete(todoID); err != nil {
		sendResponse(c, http.StatusInternalServerError, "Ошибка удаления задачи")
		return
	}

	resp.Code = http.StatusOK
	resp.Message = "Задача успешно удалена"
	sendSuccessResponse(c, resp)
}
