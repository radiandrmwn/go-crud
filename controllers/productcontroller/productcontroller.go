package productcontroller

import (
	"go-web-native/entities"
	"go-web-native/models/categorymodel"
	"go-web-native/models/productmodel"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	products := productmodel.Getall()
	data := gin.H{
		"products": products,
	}
	c.HTML(http.StatusOK, "product_index.html", data)
}

func AddGet(c *gin.Context) {
	categories := categorymodel.GetAll()
	data := gin.H{
		"categories": categories,
	}
	c.HTML(http.StatusOK, "product_create.html", data)
}

func Add(c *gin.Context) {
	var product entities.Product

	categoryId, err := strconv.Atoi(c.PostForm("category_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}

	stock, err := strconv.Atoi(c.PostForm("stock"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid stock value"})
		return
	}

	product.Name = c.PostForm("name")
	product.Category.Id = uint(categoryId)
	product.Stock = int64(stock)
	product.Description = c.PostForm("description")
	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()

	if ok := productmodel.Create(product); !ok {
		c.Redirect(http.StatusTemporaryRedirect, c.Request.Referer())
		return
	}

	c.Redirect(http.StatusSeeOther, "/products")
}

func Detail(c *gin.Context) {
	idString := c.Query("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	product := productmodel.Detail(id)
	data := gin.H{
		"product": product,
	}
	c.HTML(http.StatusOK, "product_detail.html", data)
}

func EditGet(c *gin.Context) {
	idString := c.Query("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	product := productmodel.Detail(id)
	categories := categorymodel.GetAll()

	data := gin.H{
		"product":    product,
		"categories": categories,
	}
	c.HTML(http.StatusOK, "product_edit.html", data)
}

func Edit(c *gin.Context) {
	var product entities.Product

	idString := c.PostForm("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	categoryId, err := strconv.Atoi(c.PostForm("category_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}

	stock, err := strconv.Atoi(c.PostForm("stock"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid stock value"})
		return
	}

	product.Name = c.PostForm("name")
	product.Category.Id = uint(categoryId)
	product.Stock = int64(stock)
	product.Description = c.PostForm("description")
	product.UpdatedAt = time.Now()

	if ok := productmodel.Update(id, product); !ok {
		c.Redirect(http.StatusTemporaryRedirect, c.Request.Referer())
		return
	}

	c.Redirect(http.StatusSeeOther, "/products")
}

func Delete(c *gin.Context) {
	idString := c.Query("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := productmodel.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product"})
		return
	}

	c.Redirect(http.StatusSeeOther, "/products")
}
