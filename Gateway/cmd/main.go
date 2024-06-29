package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func createProxy(target string) *httputil.ReverseProxy {
	url, err := url.Parse(target)
	if err != nil {
		log.Fatalf("Failed to parse target URL: %v", err)
	}
	return httputil.NewSingleHostReverseProxy(url)
}

func handleRequest(userServiceURL, taskServiceURL string) http.HandlerFunc {
	userProxy := createProxy(userServiceURL)
	taskProxy := createProxy(taskServiceURL)

	return func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.URL.Path == "/users" || r.URL.Path == "/users/":
			userProxy.ServeHTTP(w, r)
		case r.URL.Path == "/tasks" || r.URL.Path == "/tasks/":
			taskProxy.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	}
}

func main() {
	userServiceURL := "http://localhost:7777"
	taskServiceURL := "http://localhost:8888"

	http.HandleFunc("/", handleRequest(userServiceURL, taskServiceURL))

	log.Println("API Gateway listening on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start API Gateway: %v", err)
	}
}
