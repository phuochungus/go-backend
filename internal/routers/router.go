package routers

import (
	"go-backend/internal/controller"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.GET("/ping", controller.NewUserController().GetUser)
		v1.POST("/ping", controller.NewUserController().GetUser)
	}

	return r
}
