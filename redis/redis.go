package redis

import (
	"context"
	"fmt"
	"gin-first/config"
	"gin-first/constant"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"time"
)

var (
	myRedis *redis.Client
	ctx     = context.Background()
)

// 初始化redis
func InitRedis() {
	myRedis = redis.NewClient(&redis.Options{
		Addr: config.Global.Redis.Host + constant.MH + config.Global.Redis.Port,
		Password: config.Global.Redis.Password,
		DB: config.Global.Redis.Db,
	})
	_, err := myRedis.Ping(context.Background()).Result()
	if err != nil {
		zap.L().Error("Redis Init Error", zap.Error(err))
		panic(any(err))
	}
}

// 设置
func Set(key string, value string, t int64) bool {
	expire := time.Duration(t) * time.Second
	if err := myRedis.Set(ctx, key, value, expire).Err(); err != nil {
		return false
	}
	return true
}

// 获取
func Get(key string) string {
	result, err := myRedis.Get(ctx, key).Result()
	if err != nil {
		return ""
	}
	return result
}

// 删除
func Del(key string) bool {
	_, err := myRedis.Del(ctx, key).Result()
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

// 过期时间
func Expire(key string, t int64) bool {
	// 延长过期时间
	expire := time.Duration(t) * time.Second
	if err := myRedis.Expire(ctx, key, expire).Err(); err != nil {
		fmt.Println(err)
		return false
	}
	return true
}


