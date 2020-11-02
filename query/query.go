package query

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Name string
	Age uint
}

type Result struct {
	gorm.Model
	Name  string
	Email string
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
		panic("mysql connection fail")
	}

	var user User

	db.First(&user)
	// SELECT * FROM users ORDER BY id LIMIT 1;

	db.Take(&user)
	// SELECT * FROM users LIMIT 1;

	db.Last(&user)
	// SELECT * FROM users ORDER BY id DESC LIMIT 1;

	db.First(&user, 10)
	// SELECT * FROM users WHERE id = 10;

	db.First(&user, "10")
	// SELECT * FROM users WHERE id = 10;

	var users []User
	db.Find(&users, []int{1,2,3})
	// SELECT * FROM users WHERE id IN (1, 2, 3);

	result := db.Find(&users)
	// SELECT * FROM users
	fmt.Println(result)

	db.Where("name = ?", "Jinzhu").First(&user)
	// SELECT * FROM users WHERE name = 'Jinzhu' ORDER BY id LIMIT 1;

	db.Where("name <> ?", "Jinzhu").Find(&users)
	// SELECT * FROM users WHERE name <> 'Jinzhu';

	db.Where("name IN ?", []string{"Jinzhu", "Jinzhu 2"}).Find(&users)
	// SELECT * FROM users WHERE name IN ('Jinzhu', 'Jinzhu 2');

	db.Where("name LIKE ?", "%jin%").Find(&users)
	// SELECT * FROM users WHERE name LIKE %jin%;

	db.Where("name = ? AND age >= ?", "Jinzhu", "22").Find(&users)
	// SELECT * FROM users WHERE name = 'Jinzhu' AND age >= 22;

	var lastweek time.Time
	db.Where("update_at > ?", lastweek).Find(&users)
	// SELECT * FROM users WHERE updated_at > '2000-01-01 00:00:00';

	var today time.Time
	db.Where("created_at BETWEEN ? AND ?", lastweek, today).Find(&users)
	// SELECT * FROM users WHERE created_at BETWEEN '2000-01-01 00:00:00' AND '2000-01-08 00:00:00';

	// Inline Condition
	db.First(&user, "id = ?", "primary_key")
	// SELECT * FROM users WHERE id = 'primary_key'

	db.Find(&user, "name = ?", "jinzhu")
	// SELECT * FROM users WHERE name = 'jinzhu'

	db.Find(&user, "name <> ? AND age > ?", "junzhu", 20)
	// SELECT * FROM users WHERE name <> 'junzhu' AND age > 20

	db.Find(&users, User{Age: 20})
	// SELECT * FROM users WHERE age = 20

	db.Find(&users, map[string]interface{}{"age": 20})
	// SELECT * FROM users WHERE age = 20

	// NOT conditions

	db.Not("name = ?", "junzhu").First(&user)
	// SELECT * FROM users WHERE NOT name = 'junzhu' ORDER BY id LIMIT 1;

	db.Not(User{Name: "jinzhu", Age: 18}).First(&user)
	// SELECT * FROM users WHERE name <> "jinzhu" AND age <> 18 ORDER BY id LIMIT 1;

	db.Not([]int64{1,2,3}).First(&user)
	// SELECT * FROM users WHERE id NOT IN (1,2,3) ORDER BY id LIMIT 1;

	// Or Conditions

	db.Where("role = ?", "admin").Or("role = ?", "super_admin").Find(&users)
	// SELECT * FROM users WHERE role = 'admin' OR role = 'super_admin';

	// Struct
	db.Where("name = 'jinzhu'").Or(User{Name: "jinzhu 2", Age: 18}).Find(&users)
	// SELECT * FROM users WHERE name = 'jinzhu' OR (name = 'jinzhu 2' AND age = 18);

	// Map
	db.Where("name = 'jinzhu'").Or(map[string]interface{}{"name": "jinzhu 2", "age": 18}).Find(&users)
	// SELECT * FROM users WHERE name = 'jinzhu' OR (name = 'jinzhu 2' AND age = 18);


	// Order
	db.Order("age desc, name").Find(&users)
	// SELECT * FROM users ORDER BY age desc, name;

	db.Order("age desc").Order("name").Find(&users)
	// SELECT * FROM users ORDER BY age desc, name;

	// Limit & Offset

	db.Limit(3).Find(&users)
	// SELECT * FROM users LIMIT 3;

	// Cancel limit condition with -1
	db.Limit(10).Find(&users).Limit(-1).Find(&user)
	// SELECT * FROM users LIMIT 10; (users1)
	// SELECT * FROM users; (users2)

	db.Offset(3).Find(&users)
	// SELECT * FROM users OFFSET 3;

	db.Limit(10).Offset(5).Find(&users)
	// SELECT * FROM users OFFSET 5 LIMIT 10;

	// Cancel offset condition with -1
	db.Offset(10).Find(&users).Offset(-1).Find(&users)
	// SELECT * FROM users OFFSET 10; (users1)
	// SELECT * FROM users; (users2)


	// Distinct
	db.Distinct("name", "age").Order("name, age desc").Find(&users)


	var results Result

	// Join
	db.Model(&User{}).Select("users.name, emails.email").Joins("left join emails on emails.user_id = users.id").Scan(&results)
	// SELECT users.name, emails.email FROM `users` left join emails on emails.user_id = users.id

	db.Table("users").Select("users.name, emails.email").Joins("left join emails on emails.user_id = users.id").Scan(&results)

	// multiple joins with parameter
	db.Joins("JOIN emails ON emails.user_id = users.id AND emails.email = ?", "jinzhu@example.org").Joins("JOIN credit_cards ON credit_cards.user_id = users.id").Where("credit_cards.number = ?", "411111111111").Find(&user)


	// Joins Preloading
	db.Joins("Company").Find(&users)
	// SELECT * FROM `users` LEFT JOIN `companies` AS `Company` ON `users`.`company_id` = `Company`.`id`;
}
