package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

var pr = fmt.Println

const ADMISSION_DATE_LAYOUT = "2006-01-02T15:04:05"

func readJSONResponse(res *http.Response, concreteType any) error {
	content := make([]byte, res.ContentLength)

	if _, err := res.Body.Read(content); err != io.EOF {
		pr(content)
		return err
	}

	pr("CONTENT!: ", string(content))

	err := json.Unmarshal(content, &concreteType)
	if err != nil {
		return err
	}

	return nil
}

func sendJSONRespose(response http.ResponseWriter, json_data any) {
	json_response, err := json.Marshal(json_data)
	if err != nil {
		pr(err)
		response.WriteHeader(http.StatusServiceUnavailable)
	}
	response.Header().Add("content-type", "application/json")
	response.Write(json_response)
}

func extractIDFromPath(url_path string) (int, error) {
	paths := strings.Split(url_path, "/")
	sID := paths[len(paths)-1]
	ID, err := strconv.Atoi(sID)
	return ID, err
}

// Can't believe golang doesn't provide these out!tha!!freaking!!!box!!
func sliceContains(haystack []int, needle int) bool {
	for i := 0; i < len(haystack); i++ {
		if needle == haystack[i] {
			return true
		}
	}
	return false
}

func errorResponse(w http.ResponseWriter, err error, httpStatusCode int) {
	pr(err)
	w.WriteHeader(httpStatusCode)
}
