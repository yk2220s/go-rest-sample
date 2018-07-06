package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	"github.com/yk2220s/go-rest-sample/api/model"
)

func main() {
	db, err := gorm.Open("mysql", "root:@tcp(database:3306)/testing?charset=utf8mb4&parseTime=True&loc=Local")
	fmt.Println(err)
	db.CreateTable(&model.User{})
}
