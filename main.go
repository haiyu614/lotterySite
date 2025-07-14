package main
import (
	"github.com/gin-gonic/gin"
	"lotterySite/controller"
	"lotterySite/dao/mysql"
	"lotterySite/dao/redis"
	"fmt"
	"go.uber.org/zap"
)

func Init() error {
    fmt.Println("开始初始化 MySQL")
    if err := mysql.InitMySQL(); err != nil {
        fmt.Println("MySQL 初始化失败:", err)
        zap.L().Error("mysql init failed", zap.Error(err))
        return err
    }
    fmt.Println("MySQL 初始化成功")

    fmt.Println("开始初始化 Redis")
    if err := redis.InitRedis(); err != nil {
        fmt.Println("Redis 初始化失败:", err)
        zap.L().Error("redis init failed", zap.Error(err))
        return err
    }
    fmt.Println("Redis 初始化成功")
    return nil
}

func main() {
	fmt.Println("server start")

	if err := Init(); err != nil {
		zap.L().Error("init failed", zap.Error(err))
		fmt.Println("init failed")
		return
	}
	router := gin.Default()
	fmt.Println("router init success")
	router.GET("/goods", controller.GetAllGoods)
	router.GET("/lottery", controller.Lottery)

	router.Run("localhost:8084")
}