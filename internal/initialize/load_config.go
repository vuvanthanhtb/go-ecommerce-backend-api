package initialize

import (
	"fmt"

	"github.com/spf13/viper"
	"github.com/vuvanthanhtb/go-ecommerce-backend-api/global"
)

func LoadConfig() {
	viper := viper.New()
	viper.AddConfigPath("./config/")
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")

	// read configuration
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Failed to read configuration %w \n", err))
	}

	// read server configuration
	fmt.Println("Server port::", viper.GetInt("server.port"))
	fmt.Println("Security::", viper.GetString("security.jwt.key"))

	// configure structure
	if err = viper.Unmarshal(&global.Config); err != nil {
		fmt.Printf("Unable to decode configuration %v", err)
	}
}
