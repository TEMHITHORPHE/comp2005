package handlers

import (
	"net/http"
	model "phxlabs/m/comp2005/mockserver/src/models"
)

func An_Admission(response http.ResponseWriter, req *http.Request) {
	pr("[An_Admission]: ", req.URL)

	ID, err := extractID(req.URL.Path)
	if req.Method == "GET" && err == nil {

		ID_found := false
		var Adm model.Admission
		Admissions := model.GetAllAdmissions()

		for _, admission := range Admissions {
			if admission.Id == ID {
				Adm = admission
				ID_found = true
				break
			}
		}

		if ID_found {
			sendJSONRespose(response, Adm)
		} else {
			response.WriteHeader(404)
		}

	} else {
		pr(err)
		response.WriteHeader(403)
	}
}
