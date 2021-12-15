package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/render"
	"github.com/gorilla/mux"
	"msbeer.com/application"
)

type HandlerImpl struct {
	App application.BeerApplication
}

func NewHandler(app application.BeerApplication) HandlerImpl {
	return HandlerImpl{
		App: app,
	}
}

func (h HandlerImpl) HandleSearchBeers(w http.ResponseWriter, r *http.Request) {
	//ctx := context.Background()
	w.Header().Set("Content-Type", "application/json")

	beerList, err := h.App.SearchBeers()

	if err != nil {
		render.Status(r, http.StatusInternalServerError)
	} else {
		json.NewEncoder(w).Encode(beerList)

		render.Status(r, http.StatusOK)
	}
}

func (h HandlerImpl) HandleAddBeers(w http.ResponseWriter, r *http.Request) {
	//ctx := context.Background()

	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	idStr := vars["ID"]
	nameStr := vars["Name"]
	breweryStr := vars["Brewery"]
	countryStr := vars["Country"]
	priceStr := vars["Price"]
	currencyStr := vars["Currency"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		json.NewEncoder(w).Encode("Invalid ID value")
		render.Status(r, http.StatusInternalServerError)
		return
	}
	if nameStr == "" {
		json.NewEncoder(w).Encode("Invalid empty name")
		render.Status(r, http.StatusInternalServerError)
		return
	}
	if breweryStr == "" {
		json.NewEncoder(w).Encode("Invalid empty brewery")
		render.Status(r, http.StatusInternalServerError)
		return
	}
	if countryStr == "" {
		json.NewEncoder(w).Encode("Invalid empty country")
		render.Status(r, http.StatusInternalServerError)
		return
	}
	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		json.NewEncoder(w).Encode("Invalid Price value")
		render.Status(r, http.StatusInternalServerError)
		return
	}
	if currencyStr == "" {
		json.NewEncoder(w).Encode("Invalid empty currency")
		render.Status(r, http.StatusInternalServerError)
		return
	}

	err = h.App.AddBeers(id, nameStr, breweryStr, countryStr, price, currencyStr)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		render.Status(r, http.StatusInternalServerError)
	} else {
		render.Status(r, http.StatusOK)
	}
}

func (h HandlerImpl) HandleSearchBeerById(w http.ResponseWriter, r *http.Request) {
	//ctx := context.Background()
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	idStr := vars["ID"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		json.NewEncoder(w).Encode("Invalid ID value")
		render.Status(r, http.StatusInternalServerError)
		return
	}

	beer, err := h.App.SearchBeerById(id)

	if err != nil {
		render.Status(r, http.StatusInternalServerError)
	} else {
		json.NewEncoder(w).Encode(beer)

		render.Status(r, http.StatusOK)
	}
}

func (h HandlerImpl) HandleBoxBeerPriceById(w http.ResponseWriter, r *http.Request) {
	//ctx := context.Background()
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	idStr := vars["ID"]
	quantityStr := vars["Quantity"]
	currencyStr := vars["Currency"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		json.NewEncoder(w).Encode("Invalid ID value")
		render.Status(r, http.StatusInternalServerError)
		return
	}
	quantity, err := strconv.Atoi(quantityStr)
	if err != nil {
		json.NewEncoder(w).Encode("Invalid quantity value")
		render.Status(r, http.StatusInternalServerError)
		return
	}

	beer, err := h.App.BoxBeerPriceById(id, quantity, currencyStr)

	if err != nil {
		render.Status(r, http.StatusInternalServerError)
	} else {
		json.NewEncoder(w).Encode(beer)

		render.Status(r, http.StatusOK)
	}
}
