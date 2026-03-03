package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

var payments []map[string]interface{}
var pid = 1

func getPayments(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(payments)
}

func processPayment(w http.ResponseWriter, r *http.Request) {
	var payment map[string]interface{}
	json.NewDecoder(r.Body).Decode(&payment)
	payment["id"] = pid
	payment["status"] = "SUCCESS"
	pid++
	payments = append(payments, payment)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(payment)
}

func getPayment(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/payments/"):]
	id, _ := strconv.Atoi(idStr)

	for _, p := range payments {
		if int(p["id"].(int)) == id {
			json.NewEncoder(w).Encode(p)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

func main() {
	http.HandleFunc("/payments", getPayments)
	http.HandleFunc("/payments/process", processPayment)
	http.HandleFunc("/payments/", getPayment)

	http.ListenAndServe(":8083", nil)
}