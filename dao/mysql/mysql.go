package mysql

import (
	"fmt"
	"lotterySite/setting"
	"go.uber.org/zap"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB

func InitMySQL() (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", 
	setting.Conf.MySQLConfig.User, 
	setting.Conf.MySQLConfig.Password,
	setting.Conf.MySQLConfig.Host,
	setting.Conf.MySQLConfig.Port,
	setting.Conf.MySQLConfig.DBName,
)
	Db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		zap.L().Error("mysql connect error", zap.Error(err))
		return err
	}
	Db.SetMaxOpenConns(setting.Conf.MySQLConfig.MaxOpenConns)
	Db.SetMaxIdleConns(setting.Conf.MySQLConfig.MaxIdleConns)
	return nil
}

func CloseDB() error {
	return Db.Close()
}