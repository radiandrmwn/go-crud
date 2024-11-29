package homecontroller

import (
	"go-web-native/models/productmodel"
	"go-web-native/models/categorymodel"
	"net/http"
	"text/template"
)

func Index(w http.ResponseWriter, r *http.Request) {
	products := productmodel.Getall()
	categories := categorymodel.GetAll()

	data := map[string]any{
		"products":   products,
		"categories": categories,
	}

	temp, err := template.ParseFiles("views/home/index.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)
}

