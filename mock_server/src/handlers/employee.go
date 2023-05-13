package handlers

import (
	"net/http"
	model "phxlabs/m/comp2005/mockserver/src/models"
)

func An_Employee(response http.ResponseWriter, req *http.Request) {
	pr("[An_Employee]: ", req.URL)

	ID, err := extractID(req.URL.Path)
	if req.Method == "GET" && err == nil {

		ID_found := false
		var Empl model.Employee
		Employees := model.GetAllEmployees()

		for _, employee := range Employees {
			if employee.Id == ID {
				Empl = employee
				ID_found = true
				break
			}
		}

		if ID_found {
			sendJSONRespose(response, Empl)
		} else {
			response.WriteHeader(http.StatusNotFound)
		}

	} else {
		pr(err)
		response.WriteHeader(403)
	}
}
