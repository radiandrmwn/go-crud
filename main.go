package main

import (
	"go-web-native/config"
	"go-web-native/controllers/aboutcontroller"
	"go-web-native/controllers/categorycontroller"
	"go-web-native/controllers/homecontroller"
	"go-web-native/controllers/productcontroller"
	"log"

	"github.com/gin-gonic/gin"
	"net/http"
	// "net/http"
)

func main() {
	// Database connection
	config.ConnectDB()

	router := gin.Default()
	router.Use(CORSMiddleware())

	// Load all the templates
	router.LoadHTMLGlob("views/**/*")

	// Routes
	// 1. Homepage
	router.GET("/", homecontroller.Index)
	router.GET("/home/detail", homecontroller.DetailGet)
	router.POST("/home/detail", homecontroller.Detail)

	// 2. Categories
	router.GET("/categories", categorycontroller.Index)
	router.GET("/categories/add", categorycontroller.AddGet)
	router.POST("/categories/add", categorycontroller.Add)
	router.GET("/categories/edit", categorycontroller.EditGet)
	router.POST("/categories/edit", categorycontroller.Edit)
	router.GET("/categories/delete", categorycontroller.Delete) // Have to use a get request because delete in JavaScript and can not use DELETE route

	// 3. Products
	router.GET("/products", productcontroller.Index)
	router.GET("/products/add", productcontroller.AddGet)
	router.POST("/products/add", productcontroller.Add)
	router.GET("/products/detail", productcontroller.Detail)
	router.GET("/products/edit", productcontroller.EditGet)
	router.POST("/products/edit", productcontroller.Edit)
	router.GET("/products/delete", productcontroller.Delete)

	// 4. About
	router.GET("/about", aboutcontroller.Index)

	// Run server
	log.Println("Server running on port: 8080")
	log.Fatal(router.Run(":8080"))
}

// To create own connection to have access with html
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
}
