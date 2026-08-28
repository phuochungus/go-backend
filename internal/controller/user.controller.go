package controller

import (
	"go-backend/internal/service"
	"go-backend/pkg/response"

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
	response.SuccessResponse(c, 20001, []string{"John Doe", "Jane Smith"})
}

func (uc *UserController) TestError(c *gin.Context) {
	// Simulate an error response
	response.ErrorResponse(c, 20003, "Invalid parameters")
}
