package handlers

import (
	"net/http"
	model "phxlabs/m/comp2005/mockserver/src/models"
)

func An_Allocation(response http.ResponseWriter, req *http.Request) {
	pr("[An_Allocation]: ", req.URL)

	ID, err := extractID(req.URL.Path)
	if req.Method == "GET" && err == nil {

		ID_found := false
		var Alctn model.Allocation
		Allocations := model.GetAllAllocations()

		for _, allocation := range Allocations {
			if allocation.Id == ID {
				Alctn = allocation
				ID_found = true
				break
			}
		}

		if ID_found {
			sendJSONRespose(response, Alctn)
		} else {
			response.WriteHeader(404)
		}

	} else {
		pr(err)
		response.WriteHeader(403)
	}
}
