package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"go_hexagonal/application"
	db2 "go_hexagonal/adapters/db"
)

func main() {
	db, _ := sql.Open("sqlite3", "db.sqlite")
	productAdapter := db2.NewProductDb(db)
	productService := application.NewProductService(productAdapter)
	
	product, err := productService.Create("Product Test 2", 25)
	if err != nil {
		panic(err)
	}
	
	productService.Enable(product)
	
}