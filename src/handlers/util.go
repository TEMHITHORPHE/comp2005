package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

const (
	HOSTED_SERVER_BASE_URL = "http://localhost:8090"
	MOCK_SERVER_BASE_URL   = "http://localhost:7777"

	ALL_ADMISSIONS_ENDPOINT    = MOCK_SERVER_BASE_URL + "/admissions"
	SINGLE_ADMISSION_ENDPOINT  = MOCK_SERVER_BASE_URL + "/admission/"
	ALL_ALLOCATIONS_ENDPOINT   = MOCK_SERVER_BASE_URL + "/allocations"
	SINGLE_ADMISSIONS_ENDPOINT = MOCK_SERVER_BASE_URL + "/allocation/"
	ALL_PATIENTS_ENDPOINT      = MOCK_SERVER_BASE_URL + "/patients"
	SINGLE_PATIENT_ENDPOINT    = MOCK_SERVER_BASE_URL + "/patient/"
	ALL_EMPLOYEES_ENDPOINT     = MOCK_SERVER_BASE_URL + "/employees"
	SINGLE_EMPLOYEE_ENDPOINT   = MOCK_SERVER_BASE_URL + "/employee/"

	EMPLOYEE_PATIENTS_ENDPOINT          = HOSTED_SERVER_BASE_URL + "/employee/patients/"
	PATIENT_ADMISSION_DURATION_ENDPOINT = HOSTED_SERVER_BASE_URL + "/patients/admissions/duration/3"
	DAY_OF_MAXIMUM_ADMISSIONS_ENDPOINT  = HOSTED_SERVER_BASE_URL + "/patients/admissions/maxday/"
	EMPLOYEE_PATIENTS_DURATION_AVERAGE  = HOSTED_SERVER_BASE_URL + "/employee/patients/average/duration"
)

var pr = fmt.Println

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
	json_response, _ := json.Marshal(json_data)
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

func DayOfMaximumAdmissions(w http.ResponseWriter, r *http.Request) {

}

func EmployeePatientsDurationAverage(w http.ResponseWriter, r *http.Request) {

}
