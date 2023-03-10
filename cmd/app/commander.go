package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// The Commander is responsible for web communication/app level logic. E.g. parsing requests and responses
// to/from JSON, handling HTTP errors, etc.
type Commander struct {
	controller Controller
}

func NewCommander(c Controller) Commander { return Commander{controller: c} }

func (c Commander) GetTopAuctions(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ranking, err := strconv.Atoi(chi.URLParam(r, "ranking"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Get data to return
	auctions, err := c.controller.ListTopAuctions(ctx, ranking)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Serialize response to JSON
	serializeJSONResponse(w, auctions)
}

func serializeJSONResponse(w http.ResponseWriter, value any) {
	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
