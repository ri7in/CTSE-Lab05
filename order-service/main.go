package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

var orders []map[string]interface{}
var idCounter = 1

func getOrders(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(orders)
}

func addOrder(w http.ResponseWriter, r *http.Request) {
	var order map[string]interface{}
	json.NewDecoder(r.Body).Decode(&order)
	order["id"] = idCounter
	order["status"] = "PENDING"
	idCounter++
	orders = append(orders, order)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(order)
}

func getOrder(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/orders/"):]
	id, _ := strconv.Atoi(idStr)

	for _, o := range orders {
		if int(o["id"].(int)) == id {
			json.NewEncoder(w).Encode(o)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

func main() {
	http.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			getOrders(w, r)
		} else if r.Method == "POST" {
			addOrder(w, r)
		}
	})

	http.HandleFunc("/orders/", getOrder)

	http.ListenAndServe(":8082", nil)
}