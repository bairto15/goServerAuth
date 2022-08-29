package handler

import (
	"goServerAuth/structures"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

//Создание нового админа
func (h *Handler) NewAdmin(c *gin.Context) {
	var user structures.User
	if err := c.BindJSON(&user); err != nil {
		NewErrorResponse(c, http.StatusBadGateway, err.Error())
		return
	}
	user.Role = "admin"

	id, err := h.services.Autorization.CreateAdmin(user)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

//Создание нового пользователя
func (h *Handler) NewUser(c *gin.Context) {
	var user structures.User
	if err := c.BindJSON(&user); err != nil {
		NewErrorResponse(c, http.StatusBadGateway, err.Error())
		return
	}

	userId, err := getUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	logrus.Println("userId", userId)
	user.Root = userId
	user.Role = "user"
	logrus.Println("user", user)

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
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"list": listUsers,
	})
}

//Получить данные пользователя
func (h *Handler) GetUser(c *gin.Context) {
	id := c.Param("id")
	userId, err := strconv.Atoi(id)
	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, "user id invalid type")
		return
	}

	user, err := h.services.GetUser(userId)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

//Редактировать данные пользователя
func (h *Handler) EditUser(c *gin.Context) {
	//Получить в хэдарах id
	adminId, err := getUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	//Новые данные с запроса
	var reqUser structures.User
	if err := c.BindJSON(&reqUser); err != nil {
		NewErrorResponse(c, http.StatusBadGateway, err.Error())
		return
	}

	//Проверяем админ ли данного пользователя
	if reqUser.Root != adminId {
		NewErrorResponse(c, http.StatusInternalServerError, "invalid token")
		return
	}

	err = h.services.EditUser(reqUser)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"description": "Данные пользователя изменены",
	})
}

//Редактировать данные Админа
func (h *Handler) EditAdmin(c *gin.Context) {
	//Новые данные с запроса
	var reqUser structures.User
	if err := c.BindJSON(&reqUser); err != nil {
		NewErrorResponse(c, http.StatusBadGateway, err.Error())
		return
	}

	err := h.services.EditAdmin(reqUser)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"description": "Данные администратора изменены",
	})
}

//Удалить пользователя
func (h *Handler) DeleteUser(c *gin.Context) {
	adminId, err := getUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	id := c.Param("id")
	userId, err := strconv.Atoi(id)
	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, "user id invalid type")
		return
	}

	err = h.services.DeleteUser(userId, adminId)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"description": "Пользователь удален",
	})
}
