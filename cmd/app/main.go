package main

import (
	"errors"
	"log"
	"net/http"
	"os"
	"time"

	warcraft "github.com/ddoyle2017/lo-gosh-of-wall-street/internal/client/warcraft"
	"github.com/ddoyle2017/lo-gosh-of-wall-street/internal/rest"
)

func main() {
	// TODO: add panics if we can't find auth info
	clientID, _ := os.LookupEnv("CLIENT_ID")
	clientSecret, _ := os.LookupEnv("CLIENT_SECRET")

	httpClient := rest.NewHttpClient()
	oauth := warcraft.NewBlizzardOAuth(httpClient, "https://oauth.battle.net/token", clientID, clientSecret)
	retailAPI := warcraft.NewRetailAuctionAPI(httpClient, oauth, "https://us.api.blizzard.com")

	controller := NewController(retailAPI)
	commander := NewCommander(controller)

	appPort := "8081"
	if val, ok := os.LookupEnv("APP_PORT"); ok {
		appPort = val
	}

	server := &http.Server{
		Addr:         ":" + appPort,
		Handler:      commander.Handler(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	log.Printf("...Starting server on port: %s...", appPort)

	if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err.Error())
	}
}
