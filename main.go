package main

import (
	"lotterySite/controller"
	"lotterySite/dao/mysql"
	"lotterySite/dao/redis"
	"lotterySite/logger"
	"lotterySite/middleware"
	"lotterySite/setting"

	"github.com/gin-gonic/gin"
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
	
	router.Use(middleware.CORSMiddleware())
	router.Use(logger.GinLogger(), logger.GinRecovery(true))

	router.GET("/goods/:id", controller.GetGoodById)
	router.GET("/goods", controller.GetGoodDetailByPage)
	router.GET("/lottery", controller.Lottery)
	router.GET("/goodsforlottery", controller.GetGoodsForLottery)

	router.Run("localhost:8084")
}