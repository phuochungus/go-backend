package manage

import "github.com/gin-gonic/gin"

type UserRouter struct{}

func (ur *UserRouter) InitUserRouter(router *gin.RouterGroup) {
	// userPrivateRouter := router.Group("/admin/user")
	// {
	// 	userPrivateRouter.POST("/active_user")
	// }

}
