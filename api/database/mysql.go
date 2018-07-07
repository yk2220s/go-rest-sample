package database

import (
	"log"

	// mysql
	_ "github.com/go-sql-driver/mysql"

	"github.com/jinzhu/gorm"
	"github.com/yk2220s/go-rest-sample/env"
)

// Open open mysql connection with default dbname
func Open() *gorm.DB {
	dbName := env.GetValue("DB_NAME")
	return openMySQL(dbName)
}

// OpenWithDBName open mysql connection with dbname
func OpenWithDBName(dbName string) *gorm.DB {
	return openMySQL(dbName)
}

func openMySQL(dbName string) *gorm.DB {
	mysqlSetting := createMySQLSetting(dbName)
	db, err := gorm.Open("mysql", mysqlSetting)

	if err != nil {
		log.Fatal(err)
	}

	return db
}

func createMySQLSetting(dbName string) string {
	return env.GetValue("DB_USER") +
		":" +
		env.GetValue("DB_PASSWORD") +
		"@tcp(" +
		env.GetValue("DB_HOST") +
		":" +
		env.GetValue("DB_PORT") +
		")/" +
		dbName +
		"?charset=utf8mb4&parseTime=True&loc=Local"
}
