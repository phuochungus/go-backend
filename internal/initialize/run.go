package initialize

import (
	"fmt"
	"go-backend/global"
)

func Run() {
	LoadConfig()

	fmt.Println("Loaded config: ", global.Config)
	InitLogger()
	InitMysql()
	InitRedis()

	r := InitRouter()

	r.Run(":8002")
}
