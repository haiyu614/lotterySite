package redis

import (
	"lotterySite/dao/mysql"
	"lotterySite/model"
	"strconv"
	"go.uber.org/zap"
	"fmt"
	"github.com/go-redis/redis"
	"lotterySite/setting"

)

var Client *redis.Client

func InitRedis() error {
	Client = redis.NewClient(&redis.Options{
		Addr: 	    		fmt.Sprintf("%s:%d", setting.Conf.RedisConfig.Host, setting.Conf.RedisConfig.Port),
		Password:		 		setting.Conf.RedisConfig.Password, 
		DB:       			setting.Conf.RedisConfig.DB,
		PoolSize: 			setting.Conf.RedisConfig.PoolSize, // 连接池大小
		MinIdleConns: 	setting.Conf.RedisConfig.MinIdleConns, // 最小空闲连接数
	})
	fmt.Println(Client)
	if _, err := Client.Ping().Result(); err != nil {
		zap.L().Error("redis init error", zap.Error(err))
		return err
	}
	ch := make(chan *model.Good, 100)
	go func() {
		
		if err := mysql.GetAllGoods(ch); err!= nil {
			zap.L().Error("mysql get all goods error", zap.Error(err))
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