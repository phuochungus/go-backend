package manage

import "github.com/gin-gonic/gin"

type AdminRouter struct{}

func (ar *AdminRouter) InitAdminRouter(router *gin.RouterGroup) {
	adminRouterPublic := router.Group("/admin")
	{
		adminRouterPublic.POST("/login")
	}

	adminPrivateRouter := router.Group("/admin/user")
	{
		adminPrivateRouter.POST("/active_user")
	}

}
