package handlers

import (
	"net/http"
	model "phxlabs/m/comp2005/mockserver/src/models"
)

func Allocations(response http.ResponseWriter, req *http.Request) {
	pr("[Allocations]: ", req.Method)

	if req.Method == "GET" {

		// var Allocations []model.Allocation;
		Allocations := model.GetAllAllocations()
		sendJSONRespose(response, Allocations)

	} else {
		response.WriteHeader(403)
	}
}
