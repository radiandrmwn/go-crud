package aboutcontroller

import (
	"net/http"
	// "text/template"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "about_index.html", nil)
}