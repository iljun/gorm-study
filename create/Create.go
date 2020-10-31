package create

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Name 		string
	Age 		uint
	BirthDay 	time.Time
}

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
		panic("connection fail")
	}

	create(db)
	omit(db)

}

func create(db *gorm.DB) {
	user := User{Name: "Jinzhu", Age: 18, BirthDay: time.Now()}

	result := db.Create(&user)

	//return insert data primary key
	fmt.Println(user.ID)
	// return error
	fmt.Println(result.Error)
	// return insert record count
	fmt.Println(result.RowsAffected)

	db.Select("Name", "Age", "CreatedAt").Create(&user)
	// INSERT INTO `users` (`name`,`age`,`created_at`) VALUES ("jinzhu", 18, "2020-07-04 11:05:21.775")
}

func omit(db *gorm.DB) {
	user := User{BirthDay: time.Now()}
	db.Omit("Name", "Age", "CreatedAt").Create(&user)
	// INSERT INTO `users` (`birthday`,`updated_at`) VALUES ("2020-01-01 00:00:00.000", "2020-07-04 11:05:21.775")

}
