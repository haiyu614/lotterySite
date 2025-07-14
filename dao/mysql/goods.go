package mysql

import (
	"lotterySite/model"
	"fmt"
	"go.uber.org/zap"
)

func GetAllGoods(ch chan<- *model.Good) error {
	

	// 分页查询适合于大数据量的查询，一次性查询所有数据会导致内存占用过高，导致系统崩溃
	offset := 0
	page_size := 500
	query := "SELECT id, name, price, number, img_url FROM goods LIMIT ? OFFSET ?"
	for {
		var goods []*model.Good
		err := Db.Select(&goods, query, page_size, offset)
		if err!= nil {
			zap.L().Error("GetAllGoods db.select error", zap.Error(err))
			fmt.Println("GetAllGoods db.select error", err)
			return err
		}
		if len(goods) == 0 {
			break
		}
		for _, good := range goods {
			ch <- good
			fmt.Printf("Sending id:%d, name:%s\n", good.ID, good.Name)
		}
		offset += page_size
	}
	fmt.Println("ch Closed")
	close(ch)
	return nil
}


func GetGoodDetailById(id int) (*model.Good, error) {
	var good model.Good
	err := Db.Get(&good, "SELECT id, name, price, number, img_url FROM goods WHERE id = ?", id)
	if err != nil {
		zap.L().Error("GetGoodsById db.get error", zap.Error(err))
		return nil, err
	}
	return &good, nil
}

func GetGoodDetailByPage(page int, pageSize int) ([]*model.Good, error) {
	var goods []*model.Good
	err := Db.Select(&goods, "SELECT id, name, price, number, img_url FROM goods LIMIT ? OFFSET ?", pageSize, (page-1)*pageSize)
	if err != nil {
		zap.L().Error("GetGoodsById db.get error", zap.Error(err))
		return nil, err
	}
	return goods, nil
}