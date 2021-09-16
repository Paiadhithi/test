package main

import (
	"fmt"
	"net/http"
)

func main() {
	/*Within the main function, a route / is created using the HandleFunc method.
	This takes a pattern describing a path, followed by a function that defines how to respond to a request to that path.*/

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		/*pointer to the request.
		Inside the function, the request can be examined or manipulated before returning a response to the client*/

		fmt.Fprintf(w, "Hello World!")
	})
	http.ListenAndServe(":80", nil)
}
