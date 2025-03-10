package controllers

import (
	"encoding/json"
	"net/http"

	"backend/config"
	"backend/models"

	"database/sql"

	"log"
	"strconv"

	"github.com/gorilla/mux"
)

// func CreateUser(w http.ResponseWriter, r *http.Request) {
// 	var user models.User
// 	json.NewDecoder(r.Body).Decode(&user)

// 	_, err := config.DB.Exec("INSERT INTO users (email, password) VALUES (?, ?)", user.Email, user.Password)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	json.NewEncoder(w).Encode(user)
// }

// getting all users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	rows, _ := config.DB.Query("SELECT id, email, password FROM users")
	var users []models.User

	for rows.Next() {
		var user models.User
		rows.Scan(&user.ID, &user.Email, &user.Password)
		users = append(users, user)
	}

	json.NewEncoder(w).Encode(users)
}

// gettimng user by id
func GetUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var user models.User
	err = config.DB.QueryRow("SELECT id, email FROM users WHERE id = ?", id).Scan(&user.ID, &user.Email)
	if err == sql.ErrNoRows {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	} else if err != nil {
		log.Printf("Database error: %v", err)
		http.Error(w, "Error retrieving user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// updating user
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if user.Email == "" {
		http.Error(w, "Email is required", http.StatusBadRequest)
		return
	}

	// Update user
	_, err = config.DB.Exec("UPDATE users SET email = ? WHERE id = ?", user.Email, id)
	if err != nil {
		log.Printf("Database update error: %v", err)
		http.Error(w, "Error updating user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "User updated successfully"})
}

// removing a user
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	result, err := config.DB.Exec("DELETE FROM users WHERE id = ?", id)
	if err != nil {
		log.Printf("Database deletion error: %v", err)
		http.Error(w, "Error deleting user", http.StatusInternalServerError)
		return
	}

	// extra validation to check if user was found and deleted
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Error checking rows affected: %v", err)
		http.Error(w, "Error confirming deletion", http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "User deleted successfully"})
}
