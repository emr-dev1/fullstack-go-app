package main

import (
	"github.com/kingwerd/fullstack-go-app/api"
)

// main function acts as the entry point of the executable programs
// does not take any arguments or return anything
// Go automatically calls the function when you execute to run the program
func main() {
	// HandleFunc tells HTTP package to handle all requests to the webroot with
	// our hello function
	// http.HandleFunc("/", hello)
	// fmt.Println("Server started")
	// log.Fatal(http.ListenAndServe(":8080", nil))
	api.Run()
}

// func hello(w http.ResponseWriter, r *http.Request) {
// 	w.WriteHeader(http.StatusOK)
// 	w.Header().Set("Content-Type", "application/json")
// 	w.Write([]byte(`{"message":"hello world!"}`))
// }
