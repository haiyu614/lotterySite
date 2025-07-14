package redis

import (
	"lotterySite/model"
	"go.uber.org/zap"
	"strconv"
	"fmt"
)

const prefix = "goods:"

func GetAllGoods() ([]*model.Good, error) {
	keys, err := Client.Keys(prefix + "*").Result()
	if err != nil {
		zap.L().Error("get all goods from redis failed", zap.Error(err))
		return nil, err;
	}
	availableGoods := make([]*model.Good, 0, len(keys))
	
	for _, key := range keys {
		var id int
		if id, err = strconv.Atoi(key[len(prefix):]); err != nil {
			zap.L().Error("Get good ID failed", zap.Error(err))
			return nil, err;
		}	
		number, err := Client.Get(key).Int()
		if err != nil {
			zap.L().Error("Get good number failed", zap.Error(err))
			return nil, err;
		}
		good := &model.Good{
			ID:		 		id,
			Number:   number,
			
		}
		availableGoods = append(availableGoods, good)
	}
	return availableGoods, nil
}

func DecrGoodNumber(id int) error {
	num, err := Client.Decr(prefix + strconv.Itoa(id)).Result()
	if err != nil {
		zap.L().Error("Decr good number failed", zap.Error(err))
		return err;
	} else {
		if num < 0 {
			zap.L().Error("Good number is not enough, Decr failed", zap.Int("id", id))
			return fmt.Errorf("商品%d已无库存，下单失败", id)
		}
	}
	return nil
}
