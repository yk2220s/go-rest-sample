package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	"github.com/yk2220s/go-rest-sample/api/model"
	"github.com/yk2220s/go-rest-sample/env"
)

func main() {
	mysqlSetting := createMySQLSetting()
	db, err := gorm.Open("mysql", mysqlSetting)

	if err != nil {
		fmt.Println(err)
	}

	db.DropTable(&model.User{})
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&model.User{})

	defer db.Close()
}

func createMySQLSetting() string {
	return env.GetValue("DB_USER") +
		":" +
		env.GetValue("DB_PASSWORD") +
		"@tcp(" +
		env.GetValue("DB_HOST") +
		":" +
		env.GetValue("DB_PORT") +
		")/testing?charset=utf8mb4&parseTime=True&loc=Local"
}
