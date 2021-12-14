package main

import (
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func main() {
	http.HandleFunc("/hola", helloHandler)

	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatalln("ListenAndServer Error", err)
	}
}
