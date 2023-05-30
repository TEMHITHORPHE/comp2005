package handlers

import (
	"fmt"
	"net/http"
	HttpReq "phxlabs/m/comp2005/automated_testing/src/externalHTTP"
	model "phxlabs/m/comp2005/automated_testing/src/models"
	"time"
)

func PatientsAdmissionDuration(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		// Extract duration ... should have been passed in form of "days".
		durationID, err := extractIDFromPath(r.URL.Path)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Normalize "day" duration to "seconds"
		durationQueried, err := time.ParseDuration(fmt.Sprint(durationID*86400, "s"))
		if err != nil {
			pr(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Retrieve all Admissions.
		var admissions []model.Admission
		statusCode := HttpReq.GetAllAdmissions_HttpReq(&admissions)
		if statusCode != http.StatusOK {
			w.WriteHeader(statusCode)
			return
		}

		// Admissions Info Retrieved.
		pr("Admissions: ", admissions)

		// Retrieve IDs for patients who fit within the duration queried.
		patientIDs := make([]int, 0)
		for index, admission := range admissions {

			dateAdmitted, _err := time.Parse(ADMISSION_DATE_LAYOUT, admission.AdmissionDate)
			dateDischarged, err := time.Parse(ADMISSION_DATE_LAYOUT, admission.DischargeDate)
			if err != nil || _err != nil {
				pr(_err, err)
				w.WriteHeader(http.StatusServiceUnavailable)
				return
			}

			admissionDuration := dateDischarged.Sub(dateAdmitted)
			if admissionDuration < 0 {
				pr("[Error] Admission Date appears lower than Discharge Date, index[", index, "]")
				w.WriteHeader(http.StatusServiceUnavailable)
				return
			}

			if admissionDuration <= durationQueried {
				patientIDs = append(patientIDs, admission.PatientID)
			}
		}

		// Retrieve all patients info.
		var patients []model.Patient
		statusCode = HttpReq.GetAllPatients_HttpReq(&patients)
		if statusCode != http.StatusOK {
			w.WriteHeader(statusCode)
			return
		}

		// Patients List Retrieved.
		pr("Patients: ", patients)

		// Filter patients list based on IDs compiled earlier
		patientsInfo := make([]model.Patient, 0, len(patients))
		for _, patient := range patients {
			if sliceContains(patientIDs, patient.Id) {
				patientsInfo = append(patientsInfo, patient)
			}
		}

		sendJSONRespose(w, &patientsInfo)
		return

	} else {
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}
}
