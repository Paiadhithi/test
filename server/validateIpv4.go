package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func validateip(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	fmt.Fprintf(w, "IP address is: %v\n", vars["ipaddress"])
	ip := vars["ipaddress"]
	fmt.Println(net.ParseIP(ip))
	if net.ParseIP(ip) == nil {
		fmt.Fprintf(w, "IP address : %v is not valid\n", vars["ipaddress"])

	} else {
		fmt.Fprintf(w, "IP address : %v is valid\n", vars["ipaddress"])
		w.WriteHeader(http.StatusOK)
	}

}
func main() {
	// Creating a new router
	r := mux.NewRouter()
	// Attach a path with handler
	r.HandleFunc("/ipvalidate/{ipaddress}", validateip)
	//r.HandleFunc("/ipv4validate/{ipaddress}", ipv4validate)

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}

