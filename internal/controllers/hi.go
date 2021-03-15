package controllers

import (
	"encoding/json"
	"net/http"

	"apple-server/utils/models"
)

func (c *Controller) postHello(w http.ResponseWriter, r *http.Request) {
	var requestData models.HelloRequest

	data := r.Body
	err := json.NewDecoder(data).Decode(&requestData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	name := requestData.Name
	var responseData = &models.HelloResponse{
		Message: "Hi there Biches" + name + "! lest rock it!!",
	}
	err = json.NewEncoder(w).Encode(&responseData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
