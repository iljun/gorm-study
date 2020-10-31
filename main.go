package main

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
)

type Product struct {
	gorm.Model
	Code string
	Price uint
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("filed to connect database")
	}

	// schema migration
	db.AutoMigrate(&Product{})

	// Create
	db.Create(&Product{Code: "D42", Price: 100})

	// Read
	var product Product
	// select by Primary Key
	db.First(&product, 1)
	fmt.Println(product)

	// select by Code
	db.First(&product, "code = ?", "D42")
	fmt.Print(product)

	// Update
	db.Model(&product).Update("Price", 200)
	fmt.Print(product)
	db.Model(&product).Updates(Product{Price: 200, Code: "F42"})
	fmt.Print(product)
	db.Model(&product).Updates(map[string]interface{}{"Price":200, "Code": "F42"})
	fmt.Print(product)

	// Delete
	db.Delete(&product, 1)
}
