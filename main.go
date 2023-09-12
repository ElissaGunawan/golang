package main

import (
	"go-web/config"
	"go-web/controllers/categorycontroller"
	"go-web/controllers/homecontroller"
	"go-web/controllers/productcontroller"
	"log"
	"net/http"
)

func main() {
	config.ConnectDB()
	http.HandleFunc("/", homecontroller.Welcome)
	http.HandleFunc("/categories", categorycontroller.Index)
	http.HandleFunc("/products", productcontroller.Index)
	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
