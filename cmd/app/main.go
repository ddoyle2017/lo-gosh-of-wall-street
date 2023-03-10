package main

import (
	"log"
	"net/http"

	app "github.com/ddoyle2017/lo-gosh-of-wall-street/internal/client"
	"github.com/ddoyle2017/lo-gosh-of-wall-street/internal/rest"

	"github.com/go-chi/chi/v5"
)

func main() {
	retailAPI := app.NewRetailAuctionAPI(rest.NewHttpClient(), "https://us.api.blizzard.com")

	controller := NewController(retailAPI)
	commander := NewCommander(controller)
	router := chi.NewRouter()

	router.Route("/auctions", func(r chi.Router) {
		r.Get("/{ranking}", http.HandlerFunc(commander.GetTopAuctions))
	})

	log.Println("...Starting server...")
	http.ListenAndServe(":8080", router)
}
