//direciona os methods e paths e serveHTTP

package main

import (
	"net/http"
	"log"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func Loading() {
	r := mux.NewRouter()
	
	r.HandleFunc("/", Home)
	r.HandleFunc("/produto", InsertProductHandler).Methods("POST")
	r.HandleFunc("/produtos", ListProductsHandler).Methods("GET")
	r.HandleFunc("/produto", GetProductByNameHandler).Methods("GET")
	r.HandleFunc("/produto", DeleteProductHandler).Methods("DELETE")
	r.HandleFunc("/produto/alterar", UpdateProductHandler).Methods("PUT")
	
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})


	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(headers, methods, origins)(r)))
}
