package config

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/storage"
	firebase "firebase.google.com/go"
	"github.com/sirupsen/logrus"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/cloudkms/v1"
	"google.golang.org/api/iterator"
	"google.golang.org/api/vision/v1"
)

// Config will setup the Endpoints, the sources that will be requested, Log
type Config struct {
	Endpoints *Endpoint
	Logger    *logrus.Logger
	Firebase  *firebase.App
	Vision    *vision.Service
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
		Vision:   initializeServiceVision(),
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

func initializeServiceVision() *vision.Service {
	implicit()
	client, err := google.DefaultClient(context.Background(), vision.CloudPlatformScope)
	if err != nil {
		log.Fatalf("error initializing vision: %v\n", err)
	}
	service, err := vision.New(client)
	if err != nil {
		log.Fatalf("error initializing vision: %v\n", err)
	}
	return service
}

// implicit uses Application Default Credentials to authenticate.
func implicit() {
	ctx := context.Background()

	// For API packages whose import path is starting with "cloud.google.com/go",
	// such as cloud.google.com/go/storage in this case, if there are no credentials
	// provided, the client library will look for credentials in the environment.
	storageClient, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}

	it := storageClient.Buckets(ctx, "fala-ai-portohacksantos")
	for {
		bucketAttrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(bucketAttrs.Name)
	}

	// For packages whose import path is starting with "google.golang.org/api",
	// such as google.golang.org/api/cloudkms/v1, use the
	// golang.org/x/oauth2/google package as shown below.
	oauthClient, err := google.DefaultClient(ctx, cloudkms.CloudPlatformScope)
	if err != nil {
		log.Fatal(err)
	}

	kmsService, err := cloudkms.New(oauthClient)
	if err != nil {
		log.Fatal(err)
	}

	_ = kmsService
}
