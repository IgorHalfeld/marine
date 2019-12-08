package handler

import (
	"encoding/json"
	"net/http"
	"time"

	firebase "firebase.google.com/go"
	"github.com/marine/app/model"
	"github.com/marine/config"
)

// ---------------------------------------------
// GET MARINE TRAFFIC
// ---------------------------------------------

// GetTraffic Handlers to recover the marine data traffic
func GetTraffic(config *config.Config, w http.ResponseWriter, r *http.Request) {
	config.Logger.Info("Getting new traffic data, requesting Traffic")
	traffic, err := getMarineTrafficOr404(config.Endpoints.MarineTraffic, w, r)
	if err != nil {
		config.Logger.Error("Request NOT FOUND")
		return
	}
	respondJSON(w, http.StatusOK, traffic)
}

// getMarineTrafficOr404 gets all marine traffic, or respond with 404 error otherwise
func getMarineTrafficOr404(url string, w http.ResponseWriter, r *http.Request) (model.MarineTraffic, error) {
	properties := model.MarineTraffic{}
	if err := requestProperties(&properties, url); err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return properties, err
	}
	return properties, nil
}

// ---------------------------------------------
// GET GENERAL
// ---------------------------------------------

// GetGeneral Handlers to recover the porto de santos data
func GetGeneral(config *config.Config, w http.ResponseWriter, r *http.Request) {
	config.Logger.Info("Getting new traffic data, requesting General")
	traffic, err := getGeneralOr404(config.Endpoints.GeneralSantosPort, w, r)
	if err != nil {
		config.Logger.Error("Request NOT FOUND")
		return
	}
	respondJSON(w, http.StatusOK, traffic)
}

// getGeneralOr404 gets the porto de santos data, or respond with 404 error otherwise
func getGeneralOr404(url string, w http.ResponseWriter, r *http.Request) (model.GeneralData, error) {
	general := model.GeneralData{}
	if err := requestProperties(&general, url); err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return general, err
	}
	return general, nil
}

// ---------------------------------------------
// GET REQUEST
// ---------------------------------------------

func requestProperties(target interface{}, url string) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	client := &http.Client{}
	r, err := client.Do(req)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(target)
}

// ---------------------------------------------
// POST NEW ALERT
// ---------------------------------------------

// PostNewAlert  Handlers to save a new alert
func PostNewAlert(config *config.Config, w http.ResponseWriter, r *http.Request) {
	config.Logger.Info("Posting new alert")
	alert := postNewAlertOr500(config.Firebase, config.Endpoints.AlertDatabase, w, r)
	if alert == nil {
		config.Logger.Error("Request NOT FOUND")
		return
	}
	respondJSON(w, http.StatusOK, alert)
}

// postNewAlertOr500 will post a new alert to the system
func postNewAlertOr500(firebase *firebase.App, url string, w http.ResponseWriter, r *http.Request) *model.Alert {
	db, err := firebase.Database(r.Context())
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return nil
	}
	var alert model.Alert
	if err := json.NewDecoder(r.Body).Decode(&alert); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return nil
	}
	alert.CreatedAt = time.Now().Format(time.RFC3339)
	err = db.NewRef("alerts").Set(r.Context(), map[string]model.Alert{alert.CreatedAt: alert})
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return nil
	}
	return &alert
}
