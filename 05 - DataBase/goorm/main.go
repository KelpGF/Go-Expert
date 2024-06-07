package goorm

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func Main() {
	// basic()
	// belongsTo()
	// hasOne()
	// hasMany()
	// manyToMany()
	locks()
}

func basic() {
	type Product struct {
		ID         int `gorm:"primaryKey"`
		Name       string
		Price      float64
		gorm.Model // ID, CreatedAt, UpdatedAt, DeletedAt
	}

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

func belongsTo() {
	type Category struct {
		ID   int `gorm:"primaryKey"`
		Name string
	}

	type Product struct {
		ID         int `gorm:"primaryKey"`
		Name       string
		Price      float64
		CategoryID int
		Category   Category
		gorm.Model // ID, CreatedAt, UpdatedAt, DeletedAt
	}

	dsn := "root:root@tcp(localhost:3306)/goormexpert?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{}, &Category{})

	// create category
	category := Category{Name: "Electronics"}
	db.Create(&category)

	// create product
	product := Product{Name: "PC", Price: 1999.90, CategoryID: category.ID}
	db.Create(&product)

	// find all products
	var products []Product
	db.Find(&products)
	fmt.Println("Find all products", products) // without category data
	db.Preload("Category").Find(&products)
	fmt.Println("Find all products", products) // with category data
}

func hasOne() {
	type Category struct {
		ID   int `gorm:"primaryKey"`
		Name string
	}

	type SerialNumber struct {
		ID        int `gorm:"primaryKey"`
		Number    int
		ProductID int
	}

	type Product struct {
		ID           int `gorm:"primaryKey"`
		Name         string
		Price        float64
		CategoryID   int
		Category     Category
		SerialNumber SerialNumber
		gorm.Model   // ID, CreatedAt, UpdatedAt, DeletedAt
	}

	dsn := "root:root@tcp(localhost:3306)/goormexpert?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})

	// create category
	category := Category{Name: "Electronics"}
	db.Create(&category)

	// create product
	product := Product{Name: "PC", Price: 1999.90, CategoryID: category.ID}
	db.Create(&product)

	// create serial number
	serial := SerialNumber{ProductID: 1, Number: 123456}
	db.Create(&serial)

	// find all products
	var products []Product
	db.Preload("Category").Preload("SerialNumber").Find(&products)
	for _, p := range products {
		fmt.Printf("Product: %v, Category: %v, Serial: %v\n", p.Name, p.Category.Name, p.SerialNumber.Number)
	}
}

type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product
}

type Product struct {
	ID           int `gorm:"primaryKey"`
	Name         string
	Price        float64
	CategoryID   int
	Category     Category
	SerialNumber SerialNumber
	gorm.Model   // ID, CreatedAt, UpdatedAt, DeletedAt
}

type SerialNumber struct {
	ID        int `gorm:"primaryKey"`
	Number    int
	ProductID int
}

func hasMany() {
	dsn := "root:root@tcp(localhost:3306)/goormexpert?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})

	// create category
	category := Category{Name: "Electronics"}
	db.Create(&category)

	// create product
	product := Product{Name: "PC", Price: 1999.90, CategoryID: category.ID}
	db.Create(&product)

	// create serial number
	serial := SerialNumber{ProductID: 1, Number: 123456}
	db.Create(&serial)

	// find all categories
	var categories []Category
	_ = db.Model(&Category{}).Preload("Products").Find(&categories).Error

	for _, c := range categories {
		fmt.Printf("Category: %v\n", c.Name)
		for _, p := range c.Products {
			fmt.Printf(">> Product: %v\n", p.Name)
		}
	}

	// find all categories, products and serial numbers
	var categories2 []Category
	_ = db.Model(&Category{}).Preload("Products").Preload("Products.SerialNumber").Find(&categories2).Error

	for _, c := range categories2 {
		fmt.Printf("Category: %v\n", c.Name)
		for _, p := range c.Products {
			fmt.Printf(">> Product: %v | Serial Number: %v", p.Name, p.SerialNumber.Number)
		}
	}
}

type CategoryMM struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []ProductMM `gorm:"many2many:product_categories"`
}

type ProductMM struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Price      float64
	Categories []CategoryMM `gorm:"many2many:product_categories"`
	gorm.Model              // ID, CreatedAt, UpdatedAt, DeletedAt
}

func manyToMany() {
	dsn := "root:root@tcp(localhost:3306)/goormexpert?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&ProductMM{}, &CategoryMM{})

	// create category
	category := CategoryMM{Name: "Kitchen"}
	db.Create(&category)
	category2 := CategoryMM{Name: "Electronics"}
	db.Create(&category2)

	// create product
	product := ProductMM{
		Name:       "PC",
		Price:      1999.90,
		Categories: []CategoryMM{category, category2},
	}
	db.Create(&product)

	// find all categories
	var categories []CategoryMM
	_ = db.Model(&CategoryMM{}).Preload("Products").Find(&categories).Error

	for _, c := range categories {
		fmt.Printf("Category: %v\n", c.Name)
		for _, p := range c.Products {
			fmt.Printf(">> Product: %v\n", p.Name)
		}
	}
}

func locks() {
	type Product struct {
		ID    int `gorm:"primaryKey"`
		Name  string
		Price float64
	}

	dsn := "root:root@tcp(localhost:3306)/goormexpert?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{})

	// create product
	p := Product{Name: "PC", Price: 1999.90}
	db.Create(&p)

	tx := db.Begin()
	var product Product
	err = tx.Debug().Clauses(clause.Locking{Strength: "UPDATE"}).First(&product, 1).Error
	if err != nil {
		tx.Rollback()
		panic(err)
	}

	product.Price = 2999.90
	err = tx.Debug().Save(&product).Error
	if err != nil {
		tx.Rollback()
		panic(err)
	}

	tx.Commit()
}
