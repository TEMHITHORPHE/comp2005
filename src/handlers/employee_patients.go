package handlers

import (
	"fmt"
	"net/http"
	model "phxlabs/m/comp2005/automated_testing/src/models"
)

func EmployeePatients(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		employeeID, err := extractIDFromPath(r.URL.Path)
		if err != nil {
			// errorResponse(w, err, http.StatusBadRequest)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Validate EmployeeID parameter, make sure employee exists before moving on.
		res, err := http.Get(SINGLE_EMPLOYEE_ENDPOINT + fmt.Sprint(employeeID))
		if err != nil {
			if res.StatusCode == http.StatusNotFound {
				pr("Employee ID doesn't exist according to MOCK_SERVER.")
				w.WriteHeader(http.StatusNotFound)
			} else {
				pr(err)
				w.WriteHeader(http.StatusServiceUnavailable)
			}
			return
		}

		var employee model.Employee
		err = readJSONResponse(res, &employee)
		if err != nil {
			pr(err)
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		}
		pr("Employee:", employee)

		// At this point ... EmployeeID surely exists.

		res, err = http.Get(ALL_ALLOCATIONS_ENDPOINT)
		if err != nil || res.StatusCode != http.StatusOK {
			errorResponse(w, err, http.StatusServiceUnavailable)
			return
		}

		var allocations []model.Allocation
		err = readJSONResponse(res, &allocations)
		if err != nil {
			pr(err)
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		}
		pr("Allocations:", allocations)

		// At this point ... Allocations succesfully retrieved.

		employeeAllocations_AdmissionIDs := make([]int, 0, len(allocations))
		for _, allocation := range allocations {
			if allocation.EmployeeID == employeeID { // Which allocations belongs to this employee?
				employeeAllocations_AdmissionIDs = append(employeeAllocations_AdmissionIDs, allocation.AdmissionID)
			}
		}

		// If no employee allocations found, we return a (0)Zero-length Patient list.
		pr("EmployeeAllocations: ", employeeAllocations_AdmissionIDs)
		if len(employeeAllocations_AdmissionIDs) < 1 {
			sendJSONRespose(w, []model.Patient{})
			return
		}

		// Well, seems this employee has allocations.

		// Retrieve all Admissions
		res, err = http.Get(ALL_ADMISSIONS_ENDPOINT)
		if err != nil || res.StatusCode != http.StatusOK {
			errorResponse(w, err, http.StatusServiceUnavailable)
			return
		}

		var admissions []model.Admission
		err = readJSONResponse(res, &admissions)
		if err != nil {
			pr(err)
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		}
		pr("Admissions:", admissions)

		// At this point ...Admissions succesfully retrieved.

		// Filter admissions belonging to the employee only and retrieve the relevant patient IDs.
		employeePatientsIDs := make([]int, len(employeeAllocations_AdmissionIDs))
		for _, admission := range admissions {
			if sliceContains(employeeAllocations_AdmissionIDs, admission.Id) { // Which admissions belongs to this employee?
				employeePatientsIDs = append(employeePatientsIDs, admission.PatientID) // extract relevant PatientID from such admissions
			}
		}

		// If no patient admissions belonging to the employee is found, we return a (0)Zero-length Patient list.
		pr("Patients ID: ", employeePatientsIDs)
		if len(employeePatientsIDs) < 1 {
			sendJSONRespose(w, []model.Patient{})
			return
		}

		// Well, at this point? ... I really wanna know who designed this DB ...that said,
		// it seems this employee actually has patients assigned,
		// and I just need to get full patients details (We already have em IDs).

		w.Write([]byte("Ran to the end!!!"))

	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
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
