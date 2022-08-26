package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) NewUser(c *gin.Context) {
	getToken := c.Request.Header["Token"]
	if len(token) > 0 && getToken[0] == token {
		c.JSON(http.StatusOK, gin.H{
			"description": "Новый пользователь создан",
		})
		return
	}
	NewErrorResponse(c, http.StatusBadRequest, "Error validate token")
}

func (h *Handler) GetUser(c *gin.Context) {
	getToken := c.Request.Header["Token"]
	if len(token) > 0 && getToken[0] == token {
		c.JSON(http.StatusOK, gin.H{
			"description": "Пользователь администратор",
		})
		return
	}
	NewErrorResponse(c, http.StatusBadRequest, "Error validate token")
}

func (h *Handler) DeleteUser(c *gin.Context) {
	getToken := c.Request.Header["Token"]
	if len(token) > 0 && getToken[0] == token {
		c.JSON(http.StatusOK, gin.H{
			"description": "Пользователь удален",
		})
		return
	}
	NewErrorResponse(c, http.StatusBadRequest, "Error validate token")
}

func (h *Handler) EditUser(c *gin.Context) {
	getToken := c.Request.Header["Token"]
	if len(token) > 0 && getToken[0] == token {
		c.JSON(http.StatusOK, gin.H{
			"description": "Данные пользователя изменены",
		})
		return
	}
	NewErrorResponse(c, http.StatusBadRequest, "Error validate token")
}

func (h *Handler) GetUsers(c *gin.Context) {
	getToken := c.Request.Header["Token"]
	if len(token) > 0 && getToken[0] == token {
		c.JSON(http.StatusOK, gin.H{
			"description": "Список пользователей",
		})
		return
	}
	NewErrorResponse(c, http.StatusBadRequest, "Error validate token")
}
