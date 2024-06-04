package main

import (
	"database/sql"
	"fmt"
	"gofc-database/model"
	"gofc-database/service"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	productService := service.NewProductService(db)
	product := model.NewProduct("PC", 1999.90)
	product2 := model.NewProduct("Cell", 999.90)

	err = productService.Insert(product)
	if err != nil {
		panic(err)
	}

	product, err = productService.FindOne(product.ID)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Product %v, Name: %v, Price: %v\n", product.ID, product.Name, product.Price)

	product.Name = "Laptop"
	product.Price = 2999.90
	err = productService.Update(product)
	if err != nil {
		panic(err)
	}

	product, err = productService.FindOne(product.ID)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Product %v, Name: %v, Price: %v\n", product.ID, product.Name, product.Price)

	err = productService.Insert(product2)
	if err != nil {
		panic(err)
	}

	products, err := productService.FindAll()
	if err != nil {
		panic(err)
	}
	for _, product := range products {
		fmt.Printf("Product %v, Name: %v, Price: %v\n", product.ID, product.Name, product.Price)
	}

	err = productService.Delete(product2.ID)
	if err != nil {
		panic(err)
	}

	_, err = productService.FindOne(product2.ID)
	if err != nil {
		fmt.Println("Product not found")
	}

	err = productService.Delete(product.ID)
	if err != nil {
		panic(err)
	}
}
