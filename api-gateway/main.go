package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

func proxy(target string) http.Handler {
	url, _ := url.Parse(target)
	return httputil.NewSingleHostReverseProxy(url)
}

func main() {
	http.Handle("/items/", proxy("http://item-service:8081"))
	http.Handle("/items", proxy("http://item-service:8081"))

	http.Handle("/orders/", proxy("http://order-service:8082"))
	http.Handle("/orders", proxy("http://order-service:8082"))

	http.Handle("/payments/", proxy("http://payment-service:8083"))
	http.Handle("/payments", proxy("http://payment-service:8083"))

	http.ListenAndServe(":8080", nil)
}