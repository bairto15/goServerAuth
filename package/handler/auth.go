package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type AuthStruct struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

//Аутентификация
func (h *Handler) Auth(c *gin.Context) {
	var req AuthStruct

	if err := c.BindJSON(&req); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Autorization.GenerateToken(req.Login, req.Password)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	userId, err := h.services.Autorization.ParseToken(token)
	logrus.Println(userId)
	listUsers, err := h.services.GetUsers(userId)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"listUsers": listUsers,
	})
}
