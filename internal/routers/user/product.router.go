package user

import "github.com/gin-gonic/gin"

type ProductRouter struct{}

func (pr *ProductRouter) InitProductRouter(router *gin.RouterGroup) {
	productPublicRouter := router.Group("/product")
	{
		productPublicRouter.GET("/search")
		productPublicRouter.GET("/:id")
	}
}
