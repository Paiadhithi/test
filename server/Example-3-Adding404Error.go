package main

import "net/http"

func helloWorld(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	//Setting the header field
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte("Hello world \n"))

}
func main() {
	http.HandleFunc("/", helloWorld)
	http.ListenAndServe(":8080", nil)
}
