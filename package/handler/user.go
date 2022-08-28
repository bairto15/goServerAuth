package handler

import (
	"goServerAuth/structures"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

//Создание нового пользователя
func (h *Handler) NewUser(c *gin.Context) {
	var user structures.User
	if err := c.BindJSON(&user); err != nil {
		NewErrorResponse(c, http.StatusBadGateway, err.Error())
		return
	}

	id, err := h.services.Autorization.CreateUser(user)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

//Список пользователей
func (h *Handler) GetUsers(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	listUsers, err := h.services.GetUsers(userId)
	
	
	//---------------------------
	logrus.Println(listUsers)
}

//Получить данные пользователя
func (h *Handler) GetUser(c *gin.Context) {	
	c.JSON(http.StatusOK, gin.H{
		"description": "Пользователь администратор",
	})
}

//Удалить пользователя
func (h *Handler) DeleteUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"description": "Пользователь удален",
	})
}

//Редактировать данные пользователя
func (h *Handler) EditUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"description": "Данные пользователя изменены",
	})
}

