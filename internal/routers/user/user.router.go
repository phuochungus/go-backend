package user

import "github.com/gin-gonic/gin"

type UserRouter struct{}

func (ur *UserRouter) InitUserRouter(router *gin.RouterGroup) {
	userPublicRouter := router.Group("/user")
	{
		userPublicRouter.POST("/register")
		userPublicRouter.POST("/otp")
	}

	userPrivateRouter := router.Group("/user")
	{
		userPrivateRouter.GET("/info")
	}

}
