package main

import (
	_ "github.com/go-sql-driver/mysql"

	"github.com/yk2220s/go-rest-sample/api/database"
	"github.com/yk2220s/go-rest-sample/api/model"
	"github.com/yk2220s/go-rest-sample/env"
)

func main() {
	db := database.OpenWithDBName("")
	defer db.Close()

	db.Exec("CREATE DATABASE IF NOT EXISTS " + env.GetValue("DB_NAME"))
	db.Exec("USE " + env.GetValue("DB_NAME"))

	db.DropTableIfExists(&model.User{})
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&model.User{})

	cu := model.User{Name: "Ponta"}
	db.Create(&cu)
}
