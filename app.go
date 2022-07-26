package main

import (
	"example/helloWorldServer/handler/admin"
	"example/helloWorldServer/handler/users"
	"example/helloWorldServer/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

// Chain applies middlewares to a http.HandlerFunc
func Chain(f http.HandlerFunc, middlewares ...middleware.Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}

func main() {
	r := mux.NewRouter()

	// health check router
	r.HandleFunc("/", admin.HealthCheckHandler).Methods("GET")

	// books router
	bookrouter := r.PathPrefix("/books").Subrouter()
	bookrouter.HandleFunc("", admin.GetBooksHandler).Methods("GET")
	bookrouter.HandleFunc("/patch", admin.PostBooksHandler).Methods("PATCH")
	bookrouter.HandleFunc("/delete", admin.DeleteBooksHandler).Methods("DELETE")

	// user router
	userrouter := r.PathPrefix("/users").Subrouter()
	userrouter.HandleFunc("", Chain(users.GetUserHandler, middleware.Verification())).Methods("GET")
	userrouter.HandleFunc("/newuser", users.PostUserHandler).Methods("GET")

	http.ListenAndServe(":8000", r)
}
