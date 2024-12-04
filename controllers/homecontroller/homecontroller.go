package homecontroller

import (
	"fmt"
	"go-web-native/models/categorymodel"
	"go-web-native/models/productmodel"
	"net/http"
	"strconv"
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

func Detail(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		idString := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			http.Error(w, "Invalid product ID", http.StatusBadRequest)
			return
		}

		product := productmodel.Detail(id)
		data := map[string]any{
			"product": product,
		}

		temp, err := template.ParseFiles("views/home/detail.html")
		if err != nil {
			panic(err)
		}

		temp.Execute(w, data)
	}

	if r.Method == "POST" {
		idString := r.FormValue("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			http.Error(w, "Invalid product ID", http.StatusBadRequest)
			return
		}

		product := productmodel.Detail(id)
		
		quantityString := r.FormValue("quantity")
		quantity, err := strconv.Atoi(quantityString)
		if err != nil || quantity <= 0 {
			http.Error(w, "Invalid quantity", http.StatusBadRequest)
			return
		}

		if product.Stock < int64(quantity) {
			http.Error(w, "Not enough stock available", http.StatusBadRequest)
			return
		}

		product.Stock -= int64(quantity)

		fmt.Printf("Updating product with ID %d, new stock: %d, %s\n", id, product.Stock, product.Name)
		
		if ok := productmodel.UpdateStock(id, product); !ok {
			http.Error(w, "Failed to update stock", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

}
