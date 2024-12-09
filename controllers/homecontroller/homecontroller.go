package homecontroller

import (
	// "fmt"
	"go-web-native/models/categorymodel"
	"go-web-native/models/productmodel"
	"net/http"
	"strconv"
	// "text/template"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	products := productmodel.Getall()
	categories := categorymodel.GetAll()

data := gin.H{
		"products":   products,
		"categories": categories,
	}

	c.HTML(http.StatusOK, "home_index.html", data)
}

func DetailGet(c *gin.Context) {
	idString := c.Query("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid product ID")
		return
	}

	product := productmodel.Detail(id)
	data := gin.H{
		"product": product,
	}

	c.HTML(http.StatusOK, "home_detail.html", data)
}

func Detail(c *gin.Context) {
	idString := c.PostForm("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid product ID")
		return
	}

	product := productmodel.Detail(id)
	quantityString := c.PostForm("quantity")
	quantity, err := strconv.Atoi(quantityString)
	if err != nil || quantity <= 0 {
		c.String(http.StatusBadRequest, "Invalid quantity")
		return
	}

	if product.Stock < int64(quantity) {
		c.String(http.StatusBadRequest, "Not enough stock available")
		return
	}

	product.Stock -= int64(quantity)
	if ok := productmodel.UpdateStock(id, product); !ok {
		c.String(http.StatusInternalServerError, "Failed to update stock")
		return
	}

	c.Redirect(http.StatusSeeOther, "/")
}