package handlers

import (
	"net/http"
	model "phxlabs/m/comp2005/mockserver/src/models"
)

func A_Patient(response http.ResponseWriter, req *http.Request) {
	pr("[An_Patient]: ", req.URL)

	ID, err := extractID(req.URL.Path)
	if req.Method == "GET" && err == nil {

		ID_found := false
		var Patnt model.Employee
		Patients := model.GetAllEmployees()

		for _, patient := range Patients {
			if patient.Id == ID {
				Patnt = patient
				ID_found = true
				break
			}
		}

		if ID_found {
			sendJSONRespose(response, Patnt)
		} else {
			response.WriteHeader(404)
		}

	} else {
		pr(err)
		response.WriteHeader(403)
	}
}
