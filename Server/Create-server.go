package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/rs/cors"
)

func Enable_cors(handler http.Handler) http.Handler {
	return cors.Default().Handler(handler)
}

func enable_middleware_cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Cors := cors.New(cors.Options{
			AllowedHeaders:   []string{"Accept", "Accept-Language", "Content-Language", "Content-Type"},
			AllowedMethods:   []string{"POST", "GET"},
			AllowedOrigins:   []string{"http://127.0.0.1:5500"},
			AllowCredentials: true,
			Debug:            true,
		})
		Cors.ServeHTTP(w, r, next.ServeHTTP)
	})
}

func Create_server() {
	router := http.NewServeMux()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "welcome to server my server")
	})

	server := http.Server{
		Addr:    ":8080",
		Handler: enable_middleware_cors(router),
	}
	server.ListenAndServe()
}
