package routers

import (
	c "go-backend/internal/controller"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.GET("/ping", c.NewUserController().GetUser)
		v1.POST("/ping", c.NewUserController().GetUser)
	}

	return r
}
