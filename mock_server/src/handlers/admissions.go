package handlers

import (
	"encoding/json"
	"net/http"
	model "phxlabs/m/comp2005/mockserver/src/models"
)

type IStringOrInt interface {
	int | string
}

// type CustomMap[K string, V IStringOrInt] map[K]V

func Admissions(response http.ResponseWriter, r *http.Request) {
	pr("[Admissions]: ", r.Method)

	if r.Method == "GET" {

		// var Admissions []model.Admisson;
		Admissions := model.GetAllAdmissions()
		json_data, _ := json.Marshal(Admissions)

		response.Header().Add("content-type", "application/json")
		response.Write(json_data)

	} else {
		response.WriteHeader(403)
	}
}
