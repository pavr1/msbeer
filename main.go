package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"msbeer.com/src/adapter"
	"msbeer.com/src/application"
	"msbeer.com/src/handlers"
	"msbeer.com/src/infra"
	"msbeer.com/src/models"
)

func main() {
	config := models.NewConfig()
	httpClient := http.Client{}
	adapter := adapter.NewBeerAdapter(fmt.Sprintf(config.CurrencyURL, config.CurrencyToken), &httpClient)

	connectionString := config.DbConnectionString

	sqlObj, err := sql.Open(config.DbProvider, connectionString)
	if err != nil {
		panic(err)
	}
	conn := infra.NewDbConnectorImpl()

	infra, err := infra.NewBeerInfraImpl(sqlObj, conn)
	if err != nil {
		panic(err)
	}
	app := application.NewApplication(adapter, infra)
	handler := handlers.NewHandler(app)

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/beers", handler.HandleBeers).Methods("GET", "POST")
	router.HandleFunc("/beers/{beerID}", handler.HandleSearchBeerById)
	router.HandleFunc("/beers/{beerID}/boxprice", handler.HandleBoxBeerPriceById)

	if err := http.ListenAndServe(":3000", router); err != nil {
		log.Fatalln("ListenAndServer Error", err)
	}
}
