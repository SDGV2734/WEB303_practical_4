package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type FoodItem struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

var foodItems = []FoodItem{
	{ID: "1", Name: "Coffee", Price: 2.50},
	{ID: "2", Name: "Sandwich", Price: 5.00},
	{ID: "3", Name: "Muffin", Price: 3.25},
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	r.Get("/items", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(foodItems)
	})

	log.Println("Food Catalog Service starting on port 8080...")
	http.ListenAndServe(":8080", r)
}