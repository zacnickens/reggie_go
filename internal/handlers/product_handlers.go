package handlers

import (
    "net/http"

    "github.com/gorilla/mux"
)

func GetProductsHandler(r *http.Request) (*models.Product, error) {
    // Simulate getting products from database
    return &models.Product{ID: 1, Name: "Apple Watch"}, nil
}