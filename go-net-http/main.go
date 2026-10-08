package main

import (
	"fmt"
	"net/http"
)

func helloGoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Go TT ....")
}

func hiGoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi Go TT ....")
}

func main() {
	http.HandleFunc("/hellogo", helloGoHandler)
	http.HandleFunc("/higo", hiGoHandler)

	fmt.Printf("Server is running on port 8080...\n")

	//start the server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Error starting server: %s", err)
	}
}
