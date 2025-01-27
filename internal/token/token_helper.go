package token

import (
	"arno/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func sendResponse(c *gin.Context, code int, message string) {
	resp := models.Response{
		Code:    code,
		Message: message,
	}
	c.JSON(http.StatusOK, resp)
}
