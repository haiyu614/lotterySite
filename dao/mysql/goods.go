package mysql

import (
	"lotterySite/model"
	"fmt"
	"go.uber.org/zap"
)

func GetAllGoods(ch chan<- *model.Good) error {
	
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