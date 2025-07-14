package mysql

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/go-sql-driver/mysql"
	
)

var Db *sqlx.DB

func InitMySQL() (err error) {
	dsn := "root:jh790613@tcp(127.0.0.1:3306)/lottery?charset=utf8mb4&parseTime=True&loc=Local"
	Db, err = sqlx.Connect("mysql", dsn)
	if err!= nil {
		return err
	}
	Db.SetMaxOpenConns(10)
	Db.SetMaxIdleConns(5)
	return nil
}

func CloseDB() error {
	return Db.Close()
}