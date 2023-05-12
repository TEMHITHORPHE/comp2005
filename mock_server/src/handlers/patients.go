package handlers

import (
	"encoding/json"
	"net/http"
	model "phxlabs/m/comp2005/mockserver/src/models"
)


func Patients(response http.ResponseWriter, r *http.Request) {
	pr("[Patients]: ", r.Method)

	if r.Method == "GET" {

		// var Patients []model.Patient;
		Patients := model.GetAllPatients()
		json_data, _ := json.Marshal(Patients)

		response.Header().Add("content-type", "application/json")
		response.Write(json_data)

	} else {
		response.WriteHeader(403)
	}
}
