package database

import (
	"fmt"

	"github.com/jinzhu/gorm"

	//specify mysql driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	// DB : global variable DB
	Client *gorm.DB
)

// InitDBClient : connects to mysql database
func InitDBClient(dbUser string, dbPassword string, dbName string) (*gorm.DB, error) {
	var connectionString = fmt.Sprintf(
		"%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		dbUser, dbPassword, dbName,
	)

	db, err := gorm.Open("mysql", connectionString)
	db.LogMode(true)
	Client = db
	return db, err
}
