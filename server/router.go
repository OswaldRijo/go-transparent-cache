package server

import (
	"Golang-challenge/view"
	"github.com/gorilla/mux"
	"net/http"
)

type api struct {
	router http.Handler
}

type Server interface {
	Router() http.Handler
}

func New() Server {
	a := &api{}
	r := mux.NewRouter()
	r.HandleFunc("/v1/price/item/{item_code:[a-zA-Z0-9_]+}", view.Get).Methods(http.MethodGet)
	r.HandleFunc("/v1/price/items/search", view.GetMany).Methods(http.MethodPost)

	a.router = r
	return a
}

func (a *api) Router() http.Handler {
	return a.router
}
