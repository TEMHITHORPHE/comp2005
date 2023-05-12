package handlers

import (
	"net/http"
	model "phxlabs/m/comp2005/mockserver/src/models"
)

func An_Admission(response http.ResponseWriter, req *http.Request) {
	pr("[An_Admission]: ", req)

	if req.Method == "GET" {

		// var Admissions []model.Admisson;
		Admissions := model.GetAllAdmissions()
		sendJSONRespose(response, Admissions)

	} else {
		response.WriteHeader(403)
	}
}
