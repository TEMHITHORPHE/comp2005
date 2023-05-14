package external_http

import (
	"fmt"
	"net/http"
	model "phxlabs/m/comp2005/automated_testing/src/models"
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
	EMPLOYEE_PATIENTS_DURATION_AVERAGE  = HOSTED_SERVER_BASE_URL + "/employee/patients/average/duration/"
)

// If no error occurs, httpStatusCode is set to http.StatusOK i.e 200
func GetEmployee_HttpReq(employee_Id int, employee *model.Employee) (httpStatusCode int) {
	httpStatusCode = -1

	// Validate EmployeeID parameter, make sure employee exists before moving on.
	res, err := http.Get(SINGLE_EMPLOYEE_ENDPOINT + fmt.Sprint(employee_Id))
	if err != nil || res.StatusCode != http.StatusOK {
		if res.StatusCode == http.StatusNotFound {
			pr("Employee ID doesn't exist according to MOCK_SERVER.")
			httpStatusCode = http.StatusNotFound
		} else {
			pr(err)
			httpStatusCode = http.StatusServiceUnavailable
		}
		return httpStatusCode
	}
	defer res.Body.Close()

	err = readJSONResponse(res, employee)
	if err != nil {
		pr(err)
		httpStatusCode = http.StatusServiceUnavailable
		return httpStatusCode
	}

	if employee.Id != employee_Id {
		pr("[X] Employee ID mismatch!!! ... something is seriously! wrong! here! ... Panic!!!")
		httpStatusCode = http.StatusServiceUnavailable
		return httpStatusCode
	}

	pr("Employee:", employee)
	return http.StatusOK
}

// If no error occurs, httpStatusCode is set to http.StatusOK i.e 200
func GetAllAllocations_HttpReq(allocations *[]model.Allocation) (httpStatusCode int) {
	httpStatusCode = -1

	res, err := http.Get(ALL_ALLOCATIONS_ENDPOINT)
	if err != nil || res.StatusCode != http.StatusOK {
		pr(err) // A logging service instead??
		httpStatusCode = http.StatusServiceUnavailable
		return httpStatusCode
	}
	defer res.Body.Close()

	err = readJSONResponse(res, allocations)
	if err != nil {
		pr(err)
		httpStatusCode = http.StatusServiceUnavailable
		return httpStatusCode
	}

	pr("Allocations:", allocations)
	return http.StatusOK
}

// If no error occurs, httpStatusCode is set to http.StatusOK i.e 200
func GetAllAdmissions_HttpReq(admissions *[]model.Admission) (httpStatusCode int) {
	httpStatusCode = -1

	// Retrieve all Admissions
	res, err := http.Get(ALL_ADMISSIONS_ENDPOINT)
	if err != nil || res.StatusCode != http.StatusOK {
		pr(err)
		httpStatusCode = http.StatusServiceUnavailable
		return httpStatusCode
	}
	defer res.Body.Close()

	err = readJSONResponse(res, admissions)
	if err != nil {
		pr(err)
		httpStatusCode = http.StatusServiceUnavailable
		return httpStatusCode
	}

	pr("Admissions:", admissions)
	return http.StatusOK
}

// If no error occurs, httpStatusCode is set to http.StatusOK i.e 200
func GetAllPatients_HttpReq(patients *[]model.Patient) (httpStatusCode int) {
	httpStatusCode = -1

	// Retrieve all Patients Info
	res, err := http.Get(ALL_PATIENTS_ENDPOINT)
	if err != nil || res.StatusCode != http.StatusOK {
		pr(err)
		httpStatusCode = http.StatusServiceUnavailable
		return httpStatusCode
	}
	defer res.Body.Close()

	err = readJSONResponse(res, patients)
	if err != nil {
		pr(err)
		httpStatusCode = http.StatusServiceUnavailable
		return httpStatusCode
	}

	pr("Admissions:", patients)
	return http.StatusOK
}
