package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

var pr = fmt.Println

// How do you say "Allow only list of structs!!"?
// type ListOfStructs interface {
// 	struct
// }

// func errorCheck(err error)  {
// 	if err != nil {
// 		return nil, err
// 	}
// }

func sendJSONRespose(response http.ResponseWriter, json_data any) {
	json_response, _ := json.Marshal(json_data)
	response.Header().Add("content-type", "application/json")
	response.Write(json_response)
}

func extractID(url_path string) (int, error) {
	paths := strings.Split(url_path, "/")
	sID := paths[len(paths)-1]
	ID, err := strconv.Atoi(sID)
	return ID, err
}

// func searchByID(response http.ResponseWriter, req *http.Request) {
// 	pr("[An_Patient]: ", req.URL)

// 	ID, err := extractID(req.URL.Path)
// 	if req.Method == "GET" && err == nil {

// 		ID_found := false
// 		var Patnt model.Employee
// 		Patients := model.GetAllEmployees()

// 		for _, patient := range Patients {
// 			if patient.Id == ID {
// 				Patnt = patient
// 				ID_found = true
// 				break
// 			}
// 		}

// 		if ID_found {
// 			sendJSONRespose(response, Patnt)
// 		} else {
// 			response.WriteHeader(404)
// 		}

// 	} else {
// 		pr(err)
// 		response.WriteHeader(403)
// 	}
// }
