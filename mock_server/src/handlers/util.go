package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var pr = fmt.Println

// How do you say "Allow only list of structs!!"?
// type ListOfStructs interface {
// 	struct
// }

func sendJSONRespose(response http.ResponseWriter, json_data any) {

	json_response, _ := json.Marshal(json_data)
	response.Header().Add("content-type", "application/json")
	response.Write(json_response)
}
