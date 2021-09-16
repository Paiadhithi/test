package main

import "net/http"

func helloWorld(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	//Types of requests
	//curl -si -X GET http://localhost:8080
	switch r.Method {
	case "GET":
		w.Write([]byte("Recieved a GET request\n"))
	case "POST":
		w.Write([]byte("Recieved a POST request\n"))
	default:
		w.WriteHeader(http.StatusNotImplemented) //501->Server does not know/support HTTP method
		w.Write([]byte(http.StatusText(http.StatusNotImplemented)))
	}

}
func main() {
	http.HandleFunc("/", helloWorld)
	http.ListenAndServe(":8080", nil)
}
