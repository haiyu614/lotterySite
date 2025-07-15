package controller


import (
	"github.com/gin-gonic/gin"
	"lotterySite/model"
	"lotterySite/utils"
	"net/http"
	"strconv"
	"lotterySite/dao/redis"
	"lotterySite/dao/mysql"
	"go.uber.org/zap"
)


func filterGoods(goods []*model.Good) []*model.Good {
	availableGoods := make([]*model.Good, 0, len(goods))	
	for _, good := range goods {
		if(good.Number > 0) {
				availableGoods = append(availableGoods, good)
		}
	}
	return availableGoods
}
func Lottery(ctx *gin.Context) {
	for try := 0; try < 10; try++ { 
		// 获取所有奖品的库存
		goods, err := redis.GetAllGoods()
		if err!= nil {
			ctx.String(http.StatusOK, "获取奖品库存失败")
			return
		}
		// 过滤掉库存为0的奖品
		goods = filterGoods(goods)
		if len(goods) == 0 {
			ctx.String(http.StatusOK, strconv.Itoa(0))
			return
		}
		// 执行抽奖算法，计算抽中奖品的id
		luckyId := utils.GetLotteryId(goods)

		// 更新缓存和数据库
		err = redis.DecrGoodNumber(luckyId)
		if err!= nil {
			zap.L().Error("更新奖品库存失败", zap.Error(err))
			ctx.String(http.StatusOK, "抽奖失败")
			return
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"lucky_id": luckyId,
			})
			break
		}
	}

}

func GetGoodById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err!= nil {
		zap.L().Error("GetGoodById 参数错误", zap.Error(err))
		ctx.String(http.StatusOK, "参数错误")
		return
	}
	goods, err := mysql.GetGoodDetailById(id)
	if err!= nil {
		ctx.String(http.StatusOK, "获取奖品库存失败")
		return
	}
	ctx.JSON(http.StatusOK, goods)
}

func GetGoodDetailByPage(ctx *gin.Context) {
	page, err := strconv.Atoi(ctx.Query("page"))
	if err != nil {
		page = 1
	}
	pageSize, err := strconv.Atoi(ctx.Query("pageSize"))
	if err != nil {
		pageSize = 10
	}
	goods, err := mysql.GetGoodDetailByPage(page, pageSize)
	if err != nil {
		ctx.String(http.StatusOK, "GetGoodDetailByPage 获取奖品库存失败")
		return
	}
	ctx.JSON(http.StatusOK, goods)
}

func GetGoodsForLottery(ctx *gin.Context) {
	goods, err := redis.GetAllGoods()
	if err != nil {
		ctx.String(http.StatusOK, "GetGoodsForLottery 获取奖品库存失败")
		return
	}
	goods = filterGoods(goods)
	ctx.JSON(http.StatusOK, goods)
}