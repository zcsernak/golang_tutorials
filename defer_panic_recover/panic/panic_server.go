package main

import "net/http"

func main() {
	http.HandleFunc("/" /*myHandler*/, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Go!"))
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err.Error()) // example: panic: listen tcp :8080: bind: address already in use
	}
}

/*
func myHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Go!"))
}
*/
