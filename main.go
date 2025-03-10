package main

import (
	"fmt"
	"net/http"

	"backend/config"
	"backend/routes"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to Go Backend") // sends res on client side
}

func ProfilePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to Profile Page")
}

func main() {
	config.InitDB()

	defer config.DB.Close() // closing connection at last

	r := routes.SetupRoutes()

	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", r)
}
