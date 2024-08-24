package initialize

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/vuvanthanhtb/go-ecommerce-backend-api/global"
	"go.uber.org/zap"
)

var ctx = context.Background()

func InitRedis() {
	r := global.Config.Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%v", r.Host, r.Port),
		Password: r.Password,
		DB:       r.Database,
		PoolSize: 10,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		global.Logger.Error("Failed to connect to Redis: ", zap.Error(err))
	}

	fmt.Println("InitRedis running")
	global.Rdb = rdb
	redisExample()
}

func redisExample() {
	err := global.Rdb.Set(ctx, "Score", "100", 0).Err()
	if err != nil {
		fmt.Println("Error Redis setting:", zap.Error(err))
		return
	}

	value, err := global.Rdb.Get(ctx, "Score").Result()
	if err != nil {
		fmt.Println("Error Redis setting:", zap.Error(err))
		return
	}
	global.Logger.Info("Successfully connected to Redis: ", zap.String("score", value))
}
