package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const userCtx = "userId"

//Валидатся токена и устанвка в хедарах id user для аутентификации
func (h *Handler) UserIdentify(c *gin.Context) {
	header := c.GetHeader("Token")
	if header == "" {
		NewErrorResponse(c, http.StatusUnauthorized, "empty token header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		NewErrorResponse(c, http.StatusUnauthorized, "invalid token header")
		return
	}

	userId, err := h.services.Autorization.ParseToken(headerParts[1])
	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, "invalid token header")
		return
	}

	c.Set(userCtx, userId)
}

func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		NewErrorResponse(c, http.StatusUnauthorized, "user id not found")
		return 0, errors.New("user id not found")
	}

	idInt, ok := id.(int)
	if !ok {
		NewErrorResponse(c, http.StatusUnauthorized, "user id invalid type")
		return 0, errors.New("user id not found")
	}

	return idInt, nil
}