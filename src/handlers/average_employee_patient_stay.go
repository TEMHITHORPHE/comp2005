package handlers

import (
	"net/http"
	model "phxlabs/m/comp2005/automated_testing/src/models"
)

func EmployeePatientsDurationAverage(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		employeeID, err := extractIDFromPath(r.URL.Path)
		if err != nil {
			pr(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		var patientsInfo []model.Patient
		// Retrieve all Patients
		err = GetEmployeePatients(employeeID, &patientsInfo, w)
		if err != nil {
			return
		}

		pr("PatientsAVG: ", patientsInfo)

	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

}
