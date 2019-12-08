package handler

import (
	"bytes"
	"encoding/json"
	"net/http"

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
	alert := postNewAlertOr500(config.Endpoints.AlertDatabase, w, r)
	if alert == nil {
		config.Logger.Error("Request NOT FOUND")
		return
	}
	respondJSON(w, http.StatusOK, alert)
}

// postNewAlertOr500 will post a new alert to the system
func postNewAlertOr500(url string, w http.ResponseWriter, r *http.Request) *model.Alert {
	var alert model.Alert
	if err := json.NewDecoder(r.Body).Decode(&alert); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return nil
	}
	if err := postProperty(&alert, url); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return nil
	}
	return &alert
}

// ---------------------------------------------
// POST REQUEST
// ---------------------------------------------

func postProperty(target interface{}, url string) error {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(parseJSONBytesBuffer(target)))
	if err != nil {
		return err
	}
	client := &http.Client{}
	// TODO: save to Firebase
	r, err := client.Do(req)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(&model.ResponseMessage{Message: "Alerta Enviado com Sucesso !!"})
}
