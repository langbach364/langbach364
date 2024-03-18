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

func Router_sign_up(router *http.ServeMux) {
	router.HandleFunc("/SignUp", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch r.Method {
		case "POST":
			data, err := io.ReadAll(r.Body)
			check_err(err)
			check := sign_Up(data)
			fmt.Fprintln(w, check)
		case "GET":
			fmt.Println("Get method is not used")
		}
	})
}

func Router_sign_in(router *http.ServeMux) {
	router.HandleFunc("/SignIn", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch r.Method {
		case "POST":
			data, err := io.ReadAll(r.Body)
			check_err(err)

			check := sign_In(data)
			fmt.Fprintln(w, check)
		case "GET":
			fmt.Println("Get method is not used")
		}
	})
}

func Router_handle_search(router *http.ServeMux) {
	router.HandleFunc("/HandleSearch", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch r.Method {
		case "POST":
			data, err := io.ReadAll(r.Body)
			check_err(err)

			result := handle_search(data)
			fmt.Fprintln(w, result)
		case "GET":
			fmt.Println("Get method is not used")
		}
	})
}

func Router_view(router *http.ServeMux) {
	router.HandleFunc("/View", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch r.Method {
		case "POST":
			data, err := io.ReadAll(r.Body)
			check_err(err)

			result := view_more(data)
			fmt.Fprintln(w, result)
		case "GET":
			fmt.Println("Get method is not used")
		}
	})
}


func muxtiplexer_router(router *http.ServeMux) {
	Router_sign_up(router)
	Router_sign_in(router)
	Router_handle_search(router)
	Router_view(router)
}

func Create_server() {
	router := http.NewServeMux()
	muxtiplexer_router(router)

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
