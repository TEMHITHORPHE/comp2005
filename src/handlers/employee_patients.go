package handlers

import (
	"errors"
	"fmt"
	"net/http"
	HttpReq "phxlabs/m/comp2005/automated_testing/src/externalHTTP"
	model "phxlabs/m/comp2005/automated_testing/src/models"
)

func EmployeePatients(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		employeeID, err := extractIDFromPath(r.URL.Path)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		var patientsInfo []model.Patient
		err = GetEmployeePatients(employeeID, &patientsInfo, w)
		if err != nil {
			return
		}

		sendJSONRespose(w, patientsInfo)
		return

	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// Retrives all patients assigned to a particular employee
func GetEmployeePatients(employee_Id int, patients_info *[]model.Patient, w http.ResponseWriter) error {

	// Validate Employee ID parameter, make sure employee exists before moving on.
	var employee model.Employee
	statusCode := HttpReq.GetEmployee_HttpReq(employee_Id, &employee)
	if statusCode != http.StatusOK {
		w.WriteHeader(statusCode)
		return errors.New(fmt.Sprint(statusCode))
	}

	// At this point ... Employee ID surely exists.

	var allocations []model.Allocation
	statusCode = HttpReq.GetAllAllocations_HttpReq(&allocations)
	if statusCode != http.StatusOK {
		w.WriteHeader(statusCode)
		return errors.New(fmt.Sprint(statusCode))
	}

	// At this point ... Allocations succesfully retrieved.

	employeeAllocations_AdmissionIDs := make([]int, 0, len(allocations))
	for _, allocation := range allocations {
		if allocation.EmployeeID == employee_Id { // Which allocations belongs to this employee?
			employeeAllocations_AdmissionIDs = append(employeeAllocations_AdmissionIDs, allocation.AdmissionID)
		}
	}

	// If no employee allocations found, we return a (0)Zero-length Patient list.
	pr("EmployeeAllocations[IDs]: ", employeeAllocations_AdmissionIDs)
	if len(employeeAllocations_AdmissionIDs) < 1 {
		sendJSONRespose(w, []model.Patient{})
		return errors.New(fmt.Sprint(statusCode))
	}

	// Well, seems this employee has allocations.

	// Retrieve all Admissions
	var admissions []model.Admission
	statusCode = HttpReq.GetAllAdmissions_HttpReq(&admissions)
	if statusCode != http.StatusOK {
		w.WriteHeader(statusCode)
		return errors.New(fmt.Sprint(statusCode))
	}

	// At this point ...Admissions succesfully retrieved.

	// Filter admissions belonging to the employee only and retrieve the relevant patient IDs.
	employeePatientsIDs := make([]int, 0, len(employeeAllocations_AdmissionIDs))
	for _, admission := range admissions {
		if sliceContains(employeeAllocations_AdmissionIDs, admission.Id) { // Which admissions belongs to this employee?
			employeePatientsIDs = append(employeePatientsIDs, admission.PatientID) // extract relevant PatientID from such admissions
		}
	}

	// If no patient admissions belonging to the employee is found, we return a (0)Zero-length Patient list.
	pr("Patients ID: ", employeePatientsIDs)
	if len(employeePatientsIDs) < 1 {
		sendJSONRespose(w, []model.Patient{})
		return errors.New(fmt.Sprint(statusCode))
	}

	// Well, at this point? ... I really wanna know who designed this DB ...that said,
	// it seems this employee actually has patients assigned,
	// and I just need to get full patients details (We already have em IDs).

	// Retrieve all Patients Info
	var patients []model.Patient
	statusCode = HttpReq.GetAllPatients_HttpReq(&patients)
	if statusCode != http.StatusOK {
		w.WriteHeader(statusCode)
		return errors.New(fmt.Sprint(statusCode))
	}

	// Filter patients belonging to the employee only and retrieve the relevant patient Info.
	patientsInfo := make([]model.Patient, 0, len(employeePatientsIDs))
	for _, patient := range patients {
		if sliceContains(employeePatientsIDs, patient.Id) { // Which patient belongs(is assigned) to this employee?
			patientsInfo = append(patientsInfo, patient) // extract full Patient Info.
		}
	}

	// !!!! POSSIBLE DANGLING POINTER REFERENCE !!!
	// Is this really safe though (coming from Rust, this woulda needed lifetimes or so),
	// cos I'm not sure how Golang allocates the variable on the right "patientsInfo" (or whatever it is called now),
	// Is it allocated on the Heap? or on the Stack? ... my guess? Stack (Since i pretty much gave it a definite length).
	// EDIT:: Oops, gave it a definite Capacity ... not length ... still ... not sure how, soo ...
	*patients_info = patientsInfo
	return nil
}
