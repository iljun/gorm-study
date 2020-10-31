package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func main() {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: "gorm:gorm@tcp(127.0.0.1:3306)/gorm?charset=utf8&parseTime=True&loc=Local", // data source name
		DefaultStringSize: 256,
		DisableDatetimePrecision: true, // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex: true, // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn: true, // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
	}), &gorm.Config{})

	if err != nil {
		panic("connection failed")
	}

	sqlDb, err := db.DB()

	sqlDb.SetMaxIdleConns(1)
	sqlDb.SetMaxOpenConns(10)
	sqlDb.SetConnMaxIdleTime(time.Hour)
}