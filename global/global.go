package global

import (
	"go-backend/pkg/logger"
	"go-backend/pkg/setting"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
)
