package goorm

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Price      float64
	gorm.Model // ID, CreatedAt, UpdatedAt, DeletedAt
}

func Main() {
	dsn := "root:root@tcp(localhost:3306)/goormexpert?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// Create tables
	db.AutoMigrate(&Product{})

	// Insert
	db.Create(&Product{Name: "PC", Price: 1999.90})

	// Insert many
	products := []Product{
		{Name: "Cell", Price: 999.90},
		{Name: "Laptop", Price: 2999.90},
		{Name: "Mouse", Price: 199.90},
	}
	db.Create(&products)

	// Find one by primary key
	var product Product
	db.First(&product, 2)
	fmt.Println("First by id", product)

	// Find one by query
	var product2 Product
	db.First(&product2, "name = ?", "Laptop")
	fmt.Println("First by query", product2)

	// Find all
	var productsList []Product
	db.Find(&productsList)
	fmt.Println("Find all", productsList)

	// Find pagination
	db.Limit(2).Offset(2).Find(&productsList)
	fmt.Println("Find with Pagination", productsList)

	// Find by query
	db.Where("price > ?", 1000).Find(&productsList)
	fmt.Println("Find Where", productsList)
	db.Where("name LIKE ?", "%p%").Find(&productsList)
	fmt.Println("Find Where Like", productsList)

	// Update
	db.Model(&Product{}).Where("id = ?", 4).Update("price", 3999.90)

	// Update with struct
	var p Product
	db.First(&p, 1)
	fmt.Println("Update", p)
	p.Name = "PC Gamer"
	db.Save(&p)
	db.First(&p, 1)
	fmt.Println("Update with struct", p)

	// Delete by struct
	db.Delete(&p)

	// Delete by primary key
	db.Delete(&Product{}, 2)

	// Delete by query
	db.Where("price < ?", 2000).Delete(&Product{})
}
