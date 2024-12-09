package categorycontroller

import (
	"go-web-native/entities"
	"go-web-native/models/categorymodel"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	categories := categorymodel.GetAll()
	data := gin.H{
		"categories": categories,
	}
	c.HTML(http.StatusOK, "category_index.html", data)
}

func AddGet(c *gin.Context) {
	c.HTML(http.StatusOK, "category_create.html", nil)
}

// Traite l'ajout d'une catégorie (POST)
func Add(c *gin.Context) {
	var category entities.Category
	category.Name = c.PostForm("name")
	category.CreatedAt = time.Now()
	category.UpdatedAt = time.Now()

	if ok := categorymodel.Create(category); !ok {
		c.HTML(http.StatusInternalServerError, "category_create.html", gin.H{"error": "Failed to create category"})
		return
	}

	c.Redirect(http.StatusSeeOther, "/categories")
}

func EditGet(c *gin.Context) {
	idString := c.Query("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid category ID")
		return
	}

	category := categorymodel.Detail(id)
	data := gin.H{
		"category": category,
	}

	c.HTML(http.StatusOK, "category_edit.html", data)
}

func Edit(c *gin.Context) {
	var category entities.Category
	idString := c.PostForm("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid category ID")
		return
	}

	category.Name = c.PostForm("name")
	category.UpdatedAt = time.Now()

	if ok := categorymodel.Update(id, category); !ok {
		c.Redirect(http.StatusTemporaryRedirect, c.Request.Referer())
		return
	}

	c.Redirect(http.StatusSeeOther, "/categories")
}

func Delete(c *gin.Context) {
	idString := c.Query("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid category ID")
		return
	}

	if err := categorymodel.Delete(id); err != nil {
		c.String(http.StatusInternalServerError, "Error deleting category")
		return
	}

	c.Redirect(http.StatusSeeOther, "/categories")
}