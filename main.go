package main

import (
    "net/http"

    "test-signer/internal/api"
    "test-signer/internal/storage"
)

func main() {
    // Initialize storage
    storage := storage.NewInMemoryStorage() // Replace with other storage if needed

    // Create API handlers
    handlers := api.NewHandlers(storage)

    // Create router and handle API requests
    router := api.NewRouter(handlers)
    http.Handle("/", router)

    // Start the server
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        panic(err)
    }
}
