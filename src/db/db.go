package db

import (
	"os"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Connection() *gorm.DB {
	DBMS := "mysql"
	USER := os.Getenv("DB_USER")
	PASS := os.Getenv("DB_PASSWORD")
	PROTOCOL := "tcp(" + os.Getenv("DB_ADDRESS") + ")"
	DBNAME := os.Getenv("DB_NAME")
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=True&loc=Local"
	fmt.Println(CONNECT)

	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database: ")
	}
	db.LogMode(true)
	return db
}
