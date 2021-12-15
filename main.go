package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func searchBeers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func addBeers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func searchBeerById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["beerID"]
	w.Write([]byte("Hello World " + key))
}

func boxBeerPriceById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["beerID"]
	w.Write([]byte("Hello World " + key))
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/beers", searchBeers)
	router.HandleFunc("/beers", addBeers).Methods("POST")
	router.HandleFunc("/beers/{beerID}", searchBeerById)
	router.HandleFunc("/beers/{beerID}/boxprice", boxBeerPriceById)

	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatalln("ListenAndServer Error", err)
	}
}
