package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	Stock int    `json:"stock"`
}

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var products = []Product{
	{ID: 1, Name: "Kopi Susu", Price: 5000, Stock: 100},
	{ID: 2, Name: "Es Teh", Price: 3000, Stock: 100},
}

var categories = []Category{
	{ID: 1, Name: "Minuman"},
	{ID: 2, Name: "Makanan"},
}

var newIdIncrement = 3
var newCategoryIdIncrement = 3

func main() {
	//get product detail
	http.HandleFunc("/api/products/", func(res http.ResponseWriter, req *http.Request) {
		idStr := strings.TrimPrefix(req.URL.Path, "/api/products/")
		id, err := strconv.Atoi(idStr)

		if err != nil {
			fmt.Println(err)
			http.Error(res, "Invalid request", http.StatusBadRequest)
			return
		}

		var searchedProduct Product
		var searchedProductIdx = -1

		for i, product := range products {
			if product.ID == id {
				res.Header().Set("Content-Type", "application/json")
				searchedProduct = product
				searchedProductIdx = i
				break
			}
		}

		if searchedProductIdx == -1 {
			http.Error(res, "Product not found", http.StatusNotFound)
		}

		switch req.Method {
		case "GET":
			json.NewEncoder(res).Encode(searchedProduct)
			return

		case "PUT":
			var newProduct Product
			err := json.NewDecoder(req.Body).Decode(&newProduct)
			products[searchedProductIdx] = newProduct
			products[searchedProductIdx].ID = searchedProduct.ID

			json.NewEncoder(res).Encode(products[searchedProductIdx])

			if err != nil {
				fmt.Println(err)
				http.Error(res, "Invalid request", http.StatusBadRequest)
				return
			}
			return
		case "DELETE":
			products = append(products[:searchedProductIdx], products[searchedProductIdx+1:]...)
			json.NewEncoder(res).Encode(map[string]string{
				"status":  "ok",
				"message": "Product deleted",
			})
		}
	})

	//get product list
	//post product
	http.HandleFunc("/api/products", func(res http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case "GET":
			res.Header().Set("Content-Type", "application/json")
			json.NewEncoder(res).Encode(products)
			return

		case "POST":
			var newProduct Product
			err := json.NewDecoder(req.Body).Decode(&newProduct)

			if err != nil {
				fmt.Println(err)
				http.Error(res, "Invalid request", http.StatusBadRequest)
				return
			}

			newProduct.ID = newIdIncrement
			newIdIncrement++
			products = append(products, newProduct)
			json.NewEncoder(res).Encode(newProduct)
			return
		}
	})

	//get category detail
	http.HandleFunc("/api/categories/", func(res http.ResponseWriter, req *http.Request) {
		idStr := strings.TrimPrefix(req.URL.Path, "/api/categories/")
		id, err := strconv.Atoi(idStr)

		if err != nil {
			fmt.Println(err)
			http.Error(res, "Invalid request", http.StatusBadRequest)
			return
		}

		var searchedCategory Category
		var searchedCategoryIdx = -1

		for i, category := range categories {
			if category.ID == id {
				res.Header().Set("Content-Type", "application/json")
				searchedCategory = category
				searchedCategoryIdx = i
				break
			}
		}

		if searchedCategoryIdx == -1 {
			http.Error(res, "Category not found", http.StatusNotFound)
		}

		switch req.Method {
		case "GET":
			json.NewEncoder(res).Encode(searchedCategory)
			return

		case "PUT":
			var newCategory Category
			err := json.NewDecoder(req.Body).Decode(&newCategory)
			categories[searchedCategoryIdx] = newCategory
			categories[searchedCategoryIdx].ID = searchedCategory.ID

			json.NewEncoder(res).Encode(categories[searchedCategoryIdx])

			if err != nil {
				fmt.Println(err)
				http.Error(res, "Invalid request", http.StatusBadRequest)
				return
			}
			return
		case "DELETE":
			categories = append(categories[:searchedCategoryIdx], categories[searchedCategoryIdx+1:]...)
			json.NewEncoder(res).Encode(map[string]string{
				"status":  "ok",
				"message": "Category deleted",
			})
		}
	})

	//get category list
	//post category
	http.HandleFunc("/api/categories", func(res http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case "GET":
			res.Header().Set("Content-Type", "application/json")
			json.NewEncoder(res).Encode(categories)
			return

		case "POST":
			var newCategory Category
			err := json.NewDecoder(req.Body).Decode(&newCategory)

			if err != nil {
				fmt.Println(err)
				http.Error(res, "Invalid request", http.StatusBadRequest)
				return
			}

			newCategory.ID = newCategoryIdIncrement
			newCategoryIdIncrement++
			categories = append(categories, newCategory)
			json.NewEncoder(res).Encode(newCategory)
			return
		}
	})

	//health check
	http.HandleFunc("/health", func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "application/json")
		json.NewEncoder(res).Encode(map[string]string{
			"status":  "ok",
			"message": "Server is running",
		})
	})

	err := http.ListenAndServe(":8080", nil)

	fmt.Println("Server running at port 8080...")

	if err != nil {
		fmt.Println("Failed to run server")
	}
}
