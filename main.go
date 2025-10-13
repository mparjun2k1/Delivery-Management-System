package main

import (
    "log"
    "net/http"
    "deliverysystem/internal/order"
    "deliverysystem/internal/user"
)

func main() {
    http.HandleFunc("/register", user.RegisterHandler)
    http.HandleFunc("/login", user.LoginHandler)
    http.HandleFunc("/order", order.OrderHandler)
    http.HandleFunc("/track", order.TrackHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
