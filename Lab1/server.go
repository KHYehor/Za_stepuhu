package main

// Importing libs
import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"log"
)

func handleRoot(res http.ResponseWriter, req *http.Request) {
	// Setting header to root
	res.Header().Set("Content-Type", "text/html")
	// Sending hello world answer
	if _, err := res.Wrtite([]byte("Hello World!")); err != nil {
		log.Printf("Error writing response to the client: %s", err)
	}
}

func handleTime(res http.ResponseWriter, req *http.Request) {
	// Setting header to time
	res.Header().Set("Content-Type", "application/json")
	// Parsing from Map to JSON, Parsing time from cuurent to RFC3339 format
	jsonTime, err := json.Marshal(map[string]string{"time": time.Now().Format(time.RFC3339)})
	// Handle error if it exists
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	// Sending time
	if _, err := res.Wrtite(jsonTime); err != nil {
		log.Printf("Error writing response to the client: %s", err)
	}
}

func main() {
	// Request handles
	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/time", handleTime)
	// Deffered call
	defer http.ListenAndServe(":8795", nil)
	// Server status message
	fmt.Println("Server is listening on 8795 port...")
}
