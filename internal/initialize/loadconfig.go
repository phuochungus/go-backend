package initialize

import (
	"fmt"
	"go-backend/global"

	"github.com/spf13/viper"
)

func LoadConfig() {
	viper := viper.New()
	viper.AddConfigPath("./config/")
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Error reading config: %s", err.Error()))
	}

	if err := viper.Unmarshal(&global.Config); err != nil {
		fmt.Println("Error unmarshalling config:", err.Error())
	}
}
