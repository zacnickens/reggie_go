// Package main provides a simple server implementation.
package main

import (
    "encoding/json"
    "log"
    "net/http"

   
)

type Response struct {
    Message string `json:"message"`
}

func Serve(w http.ResponseWriter, r *http.Request) {
    resp := &Response{Message: "Hello from Reggie!"}
    json.NewEncoder(w).Encode(resp)
}

func main() {
    http.HandleFunc("/", Serve)
    log.Fatal(http.ListenAndServe(":8880", nil))
}