package app

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/marine/app/handler"
	"github.com/marine/config"
	log "github.com/sirupsen/logrus"
)

// App has router
type App struct {
	Config *config.Config
	Router *mux.Router
}

// Initialize with predefined configuration and check for environment variables
func (a *App) Initialize(config *config.Config) {
	a.Config = config
	a.Config.Logger.Formatter = &log.TextFormatter{
		FullTimestamp: true,
	}
	a.Config.Logger.Info("Initializing...")
	a.Router = mux.NewRouter()
	a.setRouters()
}

// Set all required routers
func (a *App) setRouters() {
	a.Config.Logger.Info("Setting Routers...")
	a.Router.HandleFunc("/general", a.GetGeneral).Methods("GET")
	a.Router.HandleFunc("/traffic", a.GetTraffic).Methods("GET")
	a.Router.HandleFunc("/alert", a.PostNewAlert).Methods("POST")
}

// GetGeneral Handlers to recover the porto the santos data
func (a *App) GetGeneral(w http.ResponseWriter, r *http.Request) {
	handler.GetGeneral(a.Config, w, r)
}

// GetTraffic Handlers to recover the marine data traffic
func (a *App) GetTraffic(w http.ResponseWriter, r *http.Request) {
	handler.GetTraffic(a.Config, w, r)
}

// PostNewAlert Handlers to save a new alert
func (a *App) PostNewAlert(w http.ResponseWriter, r *http.Request) {
	handler.PostNewAlert(a.Config, w, r)
}

// Run the app on it's router
func (a *App) Run(host string) {
	a.Config.Logger.Info("Listening to the port", host)
	log.Fatal(http.ListenAndServe(host, a.Router))
}
