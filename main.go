package main
import (
	"github.com/gin-gonic/gin"
	"lotterySite/controller"
	"lotterySite/dao/mysql"
	"lotterySite/dao/redis"
	"lotterySite/logger"
	"fmt"
	"lotterySite/setting"
	"go.uber.org/zap"
)

func Init() {

   
  
		if err := setting.Init("conf/config.yaml"); err!= nil {
			zap.L().Error("config init failed", zap.Error(err))
			return
		}

		if err := logger.Init(setting.Conf.LogConfig, setting.Conf.Mode); err != nil {
			zap.L().Error("logger init failed", zap.Error(err))
			return
		}

		 if err := mysql.InitMySQL(); err != nil { 
        zap.L().Error("mysql init failed", zap.Error(err))
        return 
    }
  
    if err := redis.InitRedis(); err != nil {       
        zap.L().Error("redis init failed", zap.Error(err))
        return
    }



}

func main() {


	Init()
	router := gin.Default()
	fmt.Println("router init success")
	router.GET("/goods", controller.GetAllGoods)
	router.GET("/lottery", controller.Lottery)

	router.Run("localhost:8084")
}