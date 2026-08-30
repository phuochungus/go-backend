package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	viper := viper.New()
	viper.AddConfigPath("./config")
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Error reading config: %s", err.Error())
	}

	secretJwt := viper.GetString("secret.jwt.key")
	fmt.Println("Secret JWT Key:", secretJwt)
	fmt.Println("All Settings: ", viper.AllSettings())
	fmt.Println("Test: ", viper.GetString("test"))
}
