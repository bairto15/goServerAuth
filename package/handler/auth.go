package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthStruct struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

var token = "der9834ig830tjd94jr32kihgu48ther"

//Аутентификация
func (h *Handler) Auth(c *gin.Context) {
	var req AuthStruct

	if err := c.BindJSON(&req); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if req.Login == "admin" && req.Password == "123789" {
		c.JSON(http.StatusOK, gin.H{
			"status": "Администратор",
			"token":  token,
		})
		return
	} else if req.Login == "user" && req.Password == "789789" {
		c.JSON(http.StatusOK, gin.H{
			"status": "Пользователь",
			"token":  token,
		})
		return
	}

	c.JSON(203, gin.H{
		"description": "Не правильный логин или пароль",
	})
}
