package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	switch r.Method {
	case "GET":
		for k, v := range r.URL.Query() {
			fmt.Printf("%s:%s\n", k, v)
		}
		w.Write([]byte("Recieved a GET request \n"))

	case "POST":
		res, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)

			fmt.Printf("%s\n", res)
			w.Write([]byte("Recieved a POST request \n"))

		}
	default:
		w.WriteHeader(http.StatusNotImplemented) //501->Server does not know/support HTTP method
		w.Write([]byte(http.StatusText(http.StatusNotImplemented)))
	}

}
func main() {
	http.HandleFunc("/", helloWorld)
	http.ListenAndServe(":8090", nil)

}

// curl.exe -si "http://localhost:8090/?foo=1&bar=2"
//curl.exe -si -X POST -d "some data to send" http://localhost:8090/
