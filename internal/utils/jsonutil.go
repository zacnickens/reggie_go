package utils

import (
    "encoding/json"
)

func NewEncoder(w http.ResponseWriter) *json.Encoder {
    return json.NewEncoder(w)
}

func Encode(w http.ResponseWriter, v interface{}) error {
    return json.NewEncoder(w).Encode(v)
}