package main

import (
	"go-web-native/config"
	"go-web-native/controllers/categorycontroller"
	"go-web-native/controllers/homecontroller"
	"go-web-native/controllers/productcontroller"
	"go-web-native/controllers/aboutcontroller"
	"log"
	"net/http"
)

func main() {
	// 	Database connection
	config.ConnectDB()

	// Routes
	// 1.Homepage
	http.HandleFunc("/", homecontroller.Index)

	// 2. Category
	http.HandleFunc("/categories", categorycontroller.Index)
	http.HandleFunc("/categories/add", categorycontroller.Add)
	http.HandleFunc("/categories/edit", categorycontroller.Edit)
	http.HandleFunc("/categories/delete", categorycontroller.Delete)

	// 3. Products
	http.HandleFunc("/products", productcontroller.Index)
	http.HandleFunc("/products/add", productcontroller.Add)
	http.HandleFunc("/products/detail", productcontroller.Detail)
	http.HandleFunc("/products/edit", productcontroller.Edit)
	http.HandleFunc("/products/delete", productcontroller.Delete)

	// 4. About
	http.HandleFunc("/about", aboutcontroller.Index)

	// Run server
	log.Println("Server running on port: 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
