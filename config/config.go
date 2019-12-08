package config

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"github.com/sirupsen/logrus"
)

// Config will setup the Endpoints, the sources that will be requested, Log
type Config struct {
	Endpoints *Endpoint
	Logger    *logrus.Logger
	Firebase  *firebase.App
}

// Endpoint for the future Requests
type Endpoint struct {
	MarineTraffic     string
	AlertDatabase     string
	GeneralSantosPort string
}

// GetConfig will return the config for the API
func GetConfig() *Config {
	return &Config{
		Endpoints: &Endpoint{
			MarineTraffic:     "https://www.aishub.net/station/2024/map.json?minLat=-23.98988&maxLat=-23.91301&minLon=-46.34572&maxLon=-46.26297&mode=number&zoom=13&view=true&t=1575767284",
			AlertDatabase:     "",
			GeneralSantosPort: "http://aplicacoes.portodesantos.com.br:9104/siap/servicos/atracacao/siteweb/listartotalizadores",
		},
		Logger:   logrus.New(),
		Firebase: initializeServiceAccountID(),
	}
}

func initializeServiceAccountID() *firebase.App {
	conf := &firebase.Config{
		ServiceAccountID: "106364994627247751805@fala-ai-portohacksantos.iam.gserviceaccount.com",
		DatabaseURL:      "https://fala-ai-portohacksantos.firebaseio.com",
	}
	app, err := firebase.NewApp(context.Background(), conf)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	return app
}
