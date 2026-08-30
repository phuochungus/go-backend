package routers

import (
	c "go-backend/internal/controller"
	"go-backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.GET("/ping", middleware.AuthenMiddleware, c.NewUserController().GetUser)
		v1.GET("/error", c.NewUserController().TestError)
	}

	return r
}
