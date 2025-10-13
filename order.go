package order

import (
    "log"
    "net/http"
)

// OrderHandler handles order creation or viewing.
func OrderHandler(w http.ResponseWriter, r *http.Request) {
    log.Println("OrderHandler called with method:", r.Method)

    switch r.Method {
    case http.MethodPost:
        // Simulate order creation logic
        log.Println("Processing order creation...")
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("Order created successfully!"))
    case http.MethodGet:
        // Simulate order view
        log.Println("Viewing orders...")
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("Order details page"))
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

// TrackHandler handles order tracking requests.
func TrackHandler(w http.ResponseWriter, r *http.Request) {
    log.Println("TrackHandler called with method:", r.Method)

    switch r.Method {
    case http.MethodGet:
        // Simulate order tracking
        log.Println("Processing order tracking...")
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("Order tracking details"))
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}
