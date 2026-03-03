package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

var items = []string{"Book", "Laptop", "Phone"}

func getItems(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(items)
}

func addItem(w http.ResponseWriter, r *http.Request) {
	var item string
	json.NewDecoder(r.Body).Decode(&item)
	items = append(items, item)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Item added"))
}

func getItem(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/items/"):]
	id, _ := strconv.Atoi(idStr)

	if id < 0 || id >= len(items) {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Write([]byte(items[id]))
}

func main() {
	http.HandleFunc("/items", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			getItems(w, r)
		} else if r.Method == "POST" {
			addItem(w, r)
		}
	})

	http.HandleFunc("/items/", getItem)

	http.ListenAndServe(":8081", nil)
}