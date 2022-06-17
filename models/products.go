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
