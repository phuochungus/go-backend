package initialize

import (
	"go-backend/global"
	"go-backend/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}
