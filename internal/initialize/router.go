package initialize

import (
	"go-backend/internal/routers"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := routers.NewRouter()
	return r
}
