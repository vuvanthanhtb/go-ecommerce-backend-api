package initialize

import (
	"fmt"

	"github.com/vuvanthanhtb/go-ecommerce-backend-api/global"
	"go.uber.org/zap"
)

func Run() {
	LoadConfig()
	c := global.Config.Mysql
	fmt.Println("Loading configuration mysql: ", c.Username, c.Password)
	InitLogger()
	global.Logger.Info("Logger initialized", zap.String("OK", "SUCCESS"))
	InitMysql()
	InitRedis()

	r := InitRouter()
	r.Run(fmt.Sprintf(":%d", global.Config.Server.Port))
}
