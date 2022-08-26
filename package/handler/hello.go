package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//Открыть начальную страницу и показывать версию
func (h *Handler) Hello(c *gin.Context) {
	c.String(http.StatusOK, "Go Server Auth v0.1")
}