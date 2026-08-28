package controller

import (
	"go-backend/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		service: service.NewUserService(),
	}
}

func (uc *UserController) GetUser(c *gin.Context) {
	// Get user ID from the request parameters
	c.JSON(http.StatusOK,
		gin.H{
			"user": uc.service.GetUser(),
		})
}
