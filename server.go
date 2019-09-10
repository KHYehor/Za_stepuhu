package main;
// Importing libs
import ("net/http"; "fmt"; "time"; "encoding/json");

func handleRoot(res http.ResponseWriter, req *http.Request) {
	// Setting header to root
	res.Header().Set("Content-Type", "text/html");
	// Sending hello world answer
	res.Write([]byte("Hello World!"));
}

func handleTime(res http.ResponseWriter, req *http.Request) {
	// Setting header to time
	res.Header().Set("Content-Type", "application/json");
	// Parsing from Map to JSON, Parsing time from cuurent to RFC3339 format
	jsonTime, err := json.Marshal(map[string]string{ "time": time.Now().Format(time.RFC3339) });
	// Handle error if it exists
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError);
		return;
	}
	// Sending time
	res.Write(jsonTime);
}

func main() {
	// Request handles
	http.HandleFunc("/", handleRoot);
	http.HandleFunc("/time", handleTime);
	// Deffered call
	defer http.ListenAndServe(":8795", nil);
	// Server status message
	fmt.Println("Server is listening on 9000 port...");
}