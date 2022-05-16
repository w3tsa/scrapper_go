package main

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
)

var items []Item

func getArticles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}


func main() {
	r := mux.NewRouter()

	// cors issue
	credentials := handlers.AllowCredentials()
	methods := handlers.AllowedMethods([]string{"GET"})
	// ttl := handlers.MaxAge(3600)
	origins := handlers.AllowedOrigins([]string{"http://127.0.0.1:5500"})

	// Get the data from the scrapper file
	items = scrapper()

	// Route handler/endpoints
	router :=  r.PathPrefix("/api").Subrouter()
	router.HandleFunc("/articles", getArticles).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(credentials, methods, origins)(r)))

}