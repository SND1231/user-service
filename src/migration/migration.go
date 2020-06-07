package main

import (
	"github.com/SND1231/user_service/db"
	"github.com/SND1231/user_service/model"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	db := db.Connection()
	defer db.Close()
	fmt.Printf("migration")

	db.AutoMigrate(&model.User{})
}
