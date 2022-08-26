package handler

import (
	"goServerAuth/package/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.GET("/", h.Hello)
	router.POST("/auth", h.Auth)
	router.POST("/user", h.NewUser)
	router.PATCH("/user", h.EditUser)
	router.DELETE("/user", h.DeleteUser)
	router.GET("/user", h.GetUser)
	router.GET("/users", h.GetUsers)
	return router
}
