package routers

import (
	c "go-backend/internal/controller"
	"go-backend/internal/middleware"
	"go-backend/internal/repo"
	"go-backend/internal/service"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	repo := repo.NewUserRepository()
	service := service.NewUserService(repo)
	controller := c.NewUserController(service)

	v1 := r.Group("/v1")
	{
		v1.GET("/ping", middleware.AuthenMiddleware, controller.GetUser)
		v1.GET("/error", controller.TestError)
	}

	return r
}
