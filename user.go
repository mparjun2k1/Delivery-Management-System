package user

import (
    "log"
    "net/http"
)

// RegisterHandler handles user registration requests.
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
    log.Println("RegisterHandler called with method:", r.Method)

    switch r.Method {
    case http.MethodPost:
        // Simulate processing registration logic
        log.Println("Processing user registration...")
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("User registered successfully!"))
    case http.MethodGet:
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("User registration page"))
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

// LoginHandler handles user login requests.
func LoginHandler(w http.ResponseWriter, r *http.Request) {
    log.Println("LoginHandler called with method:", r.Method)

    switch r.Method {
    case http.MethodPost:
        log.Println("Processing user login...")
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("User logged in successfully!"))
    case http.MethodGet:
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("User login page"))
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}
