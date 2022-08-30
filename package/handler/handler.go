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
	
	admin := router.Group("/admin")
	{
		admin.POST("/new", h.NewAdmin)
		admin.PATCH("/edit", h.EditAdmin)
	}
	
	user := router.Group("/user", h.UserIdentify)
	{
		user.POST("/new", h.NewUser)
		user.PATCH("/edit", h.EditUser)
		user.DELETE("/:id", h.DeleteUser)
		user.GET("/:id", h.GetUser)
		user.GET("/users", h.GetUsers)
	}
	return router
}
