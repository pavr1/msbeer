package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"msbeer.com/adapter"
	"msbeer.com/application"
	"msbeer.com/handlers"
	"msbeer.com/infra"
	"msbeer.com/models"
)

func main() {
	config := models.NewConfig()
	httpClient := http.Client{}
	adapter := adapter.NewBeerAdapter(fmt.Sprintf(config.CurrencyURL, config.CurrencyToken), &httpClient)
	infra := infra.NewBeerInfraImpl()
	app := application.NewApplication(adapter, infra)
	handler := handlers.NewHandler(app)

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/beers", handler.HandleSearchBeers)
	router.HandleFunc("/beers", handler.HandleAddBeers).Methods("POST")
	router.HandleFunc("/beers/{beerID}", handler.HandleSearchBeerById)
	router.HandleFunc("/beers/{beerID}/boxprice", handler.HandleBoxBeerPriceById)

	if err := http.ListenAndServe(":3000", router); err != nil {
		log.Fatalln("ListenAndServer Error", err)
	}
}
