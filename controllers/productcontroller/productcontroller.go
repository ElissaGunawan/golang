package productcontroller

import (
	"encoding/json"
	"go-web/models/productmodel"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {

	products := productmodel.Getall()

	jsonData, err := json.Marshal(products)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.Write(jsonData)
}

// func Add(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == "GET" {
// 		temp, err := template.ParseFiles("views/product/create.html")
// 		if err != nil {
// 			panic(err)
// 		}

// 		categories := categorymodel.GetAll()
// 		data := map[string]any{
// 			"categories": categories,
// 		}

// 		temp.Execute(w, data)
// 	}

// 	if r.Method == "POST" {
// 		var product entities.Product

// 		categoryId, err := strconv.Atoi(r.FormValue("category_id"))
// 		if err != nil {
// 			panic(err)
// 		}

// 		price, err := strconv.Atoi(r.FormValue("price"))
// 		if err != nil {
// 			panic(err)
// 		}

// 		product.Name = r.FormValue("name")
// 		product.Category.Id = uint(categoryId)
// 		product.Brand = r.FormValue("brand")
// 		product.Price = uint(price)
// 		product.Status = r.FormValue("status")

// 		if ok := productmodel.Create(product); !ok {
// 			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusTemporaryRedirect)
// 			return
// 		}

// 		http.Redirect(w, r, "/products", http.StatusSeeOther)
// 	}
// }

// func Detail(w http.ResponseWriter, r *http.Request) {
// 	idString := r.URL.Query().Get("id")

// 	id, err := strconv.Atoi(idString)
// 	if err != nil {
// 		panic(err)
// 	}

// 	product := productmodel.Detail(id)
// 	data := map[string]any{
// 		"product": product,
// 	}

// 	temp, err := template.ParseFiles("views/product/detail.html")
// 	if err != nil {
// 		panic(err)
// 	}

// 	temp.Execute(w, data)
// }

// func Edit(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == "GET" {
// 		temp, err := template.ParseFiles("views/product/edit.html")
// 		if err != nil {
// 			panic(err)
// 		}

// 		idString := r.URL.Query().Get("id")
// 		id, err := strconv.Atoi(idString)
// 		if err != nil {
// 			panic(err)
// 		}

// 		product := productmodel.Detail(id)
// 		categories := categorymodel.GetAll()

// 		data := map[string]any{
// 			"product":    product,
// 			"categories": categories,
// 		}

// 		temp.Execute(w, data)
// 	}

// 	if r.Method == "POST" {
// 		var product entities.Product

// 		idString := r.FormValue("id")
// 		id, err := strconv.Atoi(idString)
// 		if err != nil {
// 			panic(err)
// 		}

// 		categoryId, err := strconv.Atoi(r.FormValue("category_id"))
// 		if err != nil {
// 			panic(err)
// 		}

// 		price, err := strconv.Atoi(r.FormValue("price"))
// 		if err != nil {
// 			panic(err)
// 		}

// 		product.Name = r.FormValue("name")
// 		product.Category.Id = uint(categoryId)
// 		product.Brand = r.FormValue("brand")
// 		product.Price = uint(price)
// 		product.Status = r.FormValue("status")

// 		if ok := productmodel.Update(id, product); !ok {
// 			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusTemporaryRedirect)
// 			return
// 		}

// 		http.Redirect(w, r, "/products", http.StatusSeeOther)
// 	}
// }

// func Delete(w http.ResponseWriter, r *http.Request) {
// 	idString := r.URL.Query().Get("id")

// 	id, err := strconv.Atoi(idString)
// 	if err != nil {
// 		panic(err)
// 	}

// 	if err := productmodel.Delete(id); err != nil {
// 		panic(err)
// 	}

// 	http.Redirect(w, r, "/products", http.StatusSeeOther)
// }
