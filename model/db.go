package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

// Opening a database and save it to `DB`.
func InitDb(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// some settings
	db.LogMode(true)
	db.DB().SetMaxIdleConns(10)

	return db, nil
}

func GetDb() *gorm.DB {
	return db
}
