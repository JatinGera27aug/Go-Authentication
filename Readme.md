# Go Backend API Project

## Project Overview
A robust Go backend API demonstrating user authentication, CRUD operations, and JWT-based security.

## API Endpoints

-> Authentication: 
POST /register: Register a new user
POST /login: User login
POST /logout: User logout

-> Protected User Routes
GET /api/users: Get all users
GET /api/users/{id}: Get user by ID
PUT /api/users/{id}: Update user
DELETE /api/users/{id}: Delete user


## Features
- User Registration with Email Validation
- Secure Password Hashing
- JWT Authentication
- User and Product CRUD Operations
- Protected Routes
- Error Handling and Validation

## Technology Stack
- Language: Go (Golang)
- Router: Gorilla Mux
- Database: SQLite
- Authentication: JWT
- Password Hashing: bcrypt


Clone this app
cd /root-folder
Run `go mod tidy` for installing dependencies

Run `go run main.go` to start the application


