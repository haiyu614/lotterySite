package redis

import (
	"lotterySite/dao/mysql"
	"lotterySite/model"
	"strconv"
	"go.uber.org/zap"
	"fmt"
	"github.com/go-redis/redis"
)

var Client *redis.Client

func InitRedis() error {
	Client = redis.NewClient(&redis.Options{
		Addr: 	    		"localhost:6379",
		Password:		 		"jh790613", // no password set
		DB:       			0,  // use default DB
		PoolSize: 			10, // 连接池大小
		MinIdleConns: 	5, // 最小空闲连接数
	})
	ch := make(chan *model.Good, 100)
	go func() {
		fmt.Println("start get all goods from mysql")
		if err := mysql.GetAllGoods(ch); err!= nil {
			zap.L().Error("mysql get all goods error", zap.Error(err))
			fmt.Println(err)
			return
		}
	}()
	
	const prefix = "goods:"
	for good := range ch {
		if good.Number <= 0 {
			continue
		}
		key := prefix + strconv.Itoa(good.ID)
		//* 最后一个参数0表示永不过期
		err := Client.Set(key, good.Number, 0).Err()
		if err != nil {
			zap.L().Error("redis set error", zap.Error(err))
			return err
		}
	}
	return nil
}