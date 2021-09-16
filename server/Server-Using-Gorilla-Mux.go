package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	//Route
	r.HandleFunc("/", welcomeMessage)
	r.HandleFunc("/book/{title}/page/{page}", book)
	http.ListenAndServe(":8081", r)

}
func welcomeMessage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Welcome")
}
func book(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]
	page := vars["page"]
	fmt.Fprintf(w, "You have requested for %s on %s \n", title, page)
}
