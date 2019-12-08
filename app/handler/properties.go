package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/marine/app/model"
	"github.com/marine/config"
	"google.golang.org/api/vision/v1"
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
	alert := postNewAlertOr500(config, config.Endpoints.AlertDatabase, w, r)
	if alert == nil {
		config.Logger.Error("Request NOT FOUND")
		return
	}
	respondJSON(w, http.StatusOK, model.ResponseMessage{Message: "Alerta Enviado com Sucesso !!"})
}

// postNewAlertOr500 will post a new alert to the system
func postNewAlertOr500(config *config.Config, url string, w http.ResponseWriter, r *http.Request) *model.Alert {
	db, err := config.Firebase.Database(r.Context())
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return nil
	}
	var alert model.Alert
	if err := json.NewDecoder(r.Body).Decode(&alert); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return nil
	}

	annotate, err := requestIA(config, alert.ImageURL)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
	}
	alert.Labels = annotate.LabelAnnotations

	alert.CreatedAt = time.Now().Format(time.RFC3339)
	err = db.NewRef("alerts/"+alert.CreatedAt).Set(r.Context(), alert)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return nil
	}

	return &alert
}

func requestIA(config *config.Config, url string) (*vision.AnnotateImageResponse, error) {
	res, err := config.Vision.Images.Annotate(&vision.BatchAnnotateImagesRequest{
		Requests: []*vision.AnnotateImageRequest{
			&vision.AnnotateImageRequest{
				Image: &vision.Image{
					Source: &vision.ImageSource{
						ImageUri: url,
					},
				},
				Features: []*vision.Feature{
					{
						Type:       "LABEL_DETECTION",
						MaxResults: 15,
					},
				},
			},
		},
	}).Do()

	if err != nil {
		config.Logger.Error("RequestIA NOT FOUND")
		return nil, err
	}

	if len(res.Responses) == 0 {
		err = errors.New("RequestIA NOT OBJECTS")
		return nil, err
	}

	if res.Responses[0].Error != nil {
		return nil, err
	}

	return res.Responses[0], nil
}
