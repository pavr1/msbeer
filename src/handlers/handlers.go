package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/go-chi/render"
	"github.com/gorilla/mux"
	"msbeer.com/src/application"
	"msbeer.com/src/models"
)

type CustomRenderer struct {
	StatusCode int
	Status     string
	Data       interface{}
}

func (c CustomRenderer) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, c.StatusCode)
	render.Respond(w, r, c)

	return nil
}

func NewRenderer(statusCode int, err error, data interface{}) render.Renderer {
	status := ""
	if err != nil {
		status = err.Error()
	} else {
		status = "Ok"
	}

	return CustomRenderer{
		StatusCode: statusCode,
		Status:     status,
		Data:       data,
	}
}

type HandlerImpl struct {
	App application.BeerApplication
}

func NewHandler(app application.BeerApplication) HandlerImpl {
	return HandlerImpl{
		App: app,
	}
}

func (h HandlerImpl) HandleBeers(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "GET":
		list, err := h.App.SearchBeers(ctx)

		if err != nil {
			render := NewRenderer(http.StatusInternalServerError, err, nil)
			render.Render(w, r)
		} else {
			render := NewRenderer(http.StatusOK, nil, list)
			render.Render(w, r)
		}
	case "POST":
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			render := NewRenderer(http.StatusInternalServerError, err, nil)
			render.Render(w, r)
			return
		}

		var beer models.BeerItem
		err = json.Unmarshal(b, &beer)
		if err != nil {
			render := NewRenderer(http.StatusInternalServerError, err, nil)
			render.Render(w, r)
			return
		}

		result, err := h.App.SearchBeerById(ctx, beer.ID)
		if err != nil {
			render := NewRenderer(http.StatusInternalServerError, errors.New("error while validating beer existance"), nil)
			render.Render(w, r)
			return
		}
		if result != nil {
			render := NewRenderer(http.StatusInternalServerError, errors.New("id value already existent, please change id"), nil)
			render.Render(w, r)
			return
		}

		if beer.ID <= 0 {
			render := NewRenderer(http.StatusInternalServerError, errors.New("invalid id value"), nil)
			render.Render(w, r)
			return
		}
		if beer.Name == "" {
			render := NewRenderer(http.StatusInternalServerError, errors.New("invalid empty name value"), nil)
			render.Render(w, r)
			return
		}
		if beer.Brewery == "" {
			render := NewRenderer(http.StatusInternalServerError, errors.New("invalid empty brewery value"), nil)
			render.Render(w, r)
			return
		}
		if beer.Country == "" {
			render := NewRenderer(http.StatusInternalServerError, errors.New("invalid empty country value"), nil)
			render.Render(w, r)
			return
		}
		if beer.Price == 0 {
			render := NewRenderer(http.StatusInternalServerError, errors.New("invalid price value"), nil)
			render.Render(w, r)
			return
		}
		if beer.Currency == "" {
			render := NewRenderer(http.StatusInternalServerError, errors.New("invalid empty currency value"), nil)
			render.Render(w, r)
			return
		}

		err = h.App.AddBeers(ctx, beer)
		if err != nil {
			json.NewEncoder(w).Encode(err.Error())
			render.Status(r, http.StatusInternalServerError)
		} else {
			render.Status(r, http.StatusOK)
		}
	}
}

func (h HandlerImpl) HandleSearchBeerById(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	idStr := vars["ID"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		json.NewEncoder(w).Encode("Invalid ID value")
		render.Status(r, http.StatusInternalServerError)
		return
	}

	beer, err := h.App.SearchBeerById(ctx, id)

	if err != nil {
		render.Status(r, http.StatusInternalServerError)
	} else {
		json.NewEncoder(w).Encode(beer)

		render.Status(r, http.StatusOK)
	}
}

func (h HandlerImpl) HandleBoxBeerPriceById(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
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

	beer, err := h.App.BoxBeerPriceById(ctx, id, quantity, currencyStr)

	if err != nil {
		render.Status(r, http.StatusInternalServerError)
	} else {
		json.NewEncoder(w).Encode(beer)

		render.Status(r, http.StatusOK)
	}
}
