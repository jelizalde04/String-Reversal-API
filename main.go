package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Reversed string `json:"reversed"`
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Get the 'string' parameter from the URL query
	query := r.URL.Query()
	str := query.Get("string")

	if str != "" {
		reversed := reverseString(str)
		response := Response{Reversed: reversed}
		json.NewEncoder(w).Encode(response)
	} else {
		http.Error(w, "Missing 'string' parameter", http.StatusBadRequest)
	}
}

func main() {
	http.HandleFunc("/reverse", handler)
	fmt.Println("Server listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
