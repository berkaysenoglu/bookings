package main

import (
	"fmt"
	"net/http"

	"github.com/berkaysenoglu/bookings/pkg/handlers"
)

const portNumber = ":8080"

// main is the main function
func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Staring application on port %s", portNumber)
	_ = http.ListenAndServe(":8080", nil)
}
