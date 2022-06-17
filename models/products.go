package models

import (
	"github.com/LeonardoCampello-dev/go-web/db"
)

type Product struct {
	Id int
	Name,
	Description string
	Price    float64
	Quantity int
}

func SelectAllProducts() []Product {
	db := db.ConnectDB()

	selectAllProducts, err := db.Query("select * from products")

	if err != nil {
		panic(err.Error())
	}

	currentProduct := Product{}

	products := []Product{}

	for selectAllProducts.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = selectAllProducts.Scan(&id, &name, &description, &price, &quantity)

		if err != nil {
			panic(err.Error())
		}

		currentProduct.Name = name
		currentProduct.Description = description
		currentProduct.Price = price
		currentProduct.Quantity = quantity

		products = append(products, currentProduct)
	}

	defer db.Close()

	return products
}

func InsertNewProduct(name, description string, price float64, quantity int) {
	db := db.ConnectDB()

	insertProductQuery, err := db.Prepare("insert into products (name, description, price, quantity) values ($1, $2, $3, $4)")

	if err != nil {
		panic(err.Error())
	}

	insertProductQuery.Exec(name, description, price, quantity)

	defer db.Close()
}
