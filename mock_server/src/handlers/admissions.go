package handlers

import (
	"net/http"
	model "phxlabs/m/comp2005/mockserver/src/models"
)

type IStringOrInt interface {
	int | string
}

// type CustomMap[K string, V IStringOrInt] map[K]V

func Admissions(response http.ResponseWriter, req *http.Request) {
	pr("[Admissions]: ", req.Method)

	if req.Method == "GET" {

		// var Admissions []model.Admisson;
		Admissions := model.GetAllAdmissions()
		sendJSONRespose(response, Admissions)

	} else {
		response.WriteHeader(403)
	}
}
