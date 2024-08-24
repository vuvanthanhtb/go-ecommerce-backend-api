package initialize

import (
	"github.com/vuvanthanhtb/go-ecommerce-backend-api/global"
	"github.com/vuvanthanhtb/go-ecommerce-backend-api/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}
