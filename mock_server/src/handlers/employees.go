package handlers

import (
	"net/http"
	model "phxlabs/m/comp2005/mockserver/src/models"
)

func Employees(response http.ResponseWriter, req *http.Request) {
	pr("[Employees]: ", req.Method)

	if req.Method == "GET" {

		// var Employees []model.Employee;
		Employees := model.GetAllEmployees()
		sendJSONRespose(response, Employees)

	} else {
		response.WriteHeader(403)
	}
}
