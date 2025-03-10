package routes

import (
	"backend/auth"
	"backend/controllers"
	middleware "backend/middlewares"

	"github.com/gorilla/mux"
	// replacement for http.ServeMux
)

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()

	// public Routes
	r.HandleFunc("/register", auth.RegisterUser).Methods("POST")
	r.HandleFunc("/login", auth.LoginUser).Methods("POST")

	// Protected Routes
	protectedRoutes := r.PathPrefix("/api").Subrouter() // starting with /api walo ke liye
	protectedRoutes.Use(middleware.ValidateJWT)

	protectedRoutes.HandleFunc("/users", controllers.GetUsers).Methods("GET")
	protectedRoutes.HandleFunc("/users/{id}", controllers.GetUserByID).Methods("GET")
	protectedRoutes.HandleFunc("/users/{id}", controllers.UpdateUser).Methods("PUT")
	protectedRoutes.HandleFunc("/users/{id}", controllers.DeleteUser).Methods("DELETE")

	return r
}
