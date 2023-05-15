package handlers

import (
	"fmt"
	"net/http"
	HttpReq "phxlabs/m/comp2005/automated_testing/src/externalHTTP"
	model "phxlabs/m/comp2005/automated_testing/src/models"
	"time"
)

func EmployeePatientsDurationAverage(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		employeeID, err := extractIDFromPath(r.URL.Path)
		if err != nil {
			pr(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Retrieve all Patients
		var patientsInfo []model.Patient
		err = GetEmployeePatients(employeeID, &patientsInfo, w)
		if err != nil {
			return
		}
		pr("PatientsAVG: ", patientsInfo)

		patientIDs := make([]int, 0, len(patientsInfo))
		for _, patient := range patientsInfo {
			patientIDs = append(patientIDs, patient.Id)
		}

		// Retrieve all Admissions
		var admissions []model.Admission
		statusCode := HttpReq.GetAllAdmissions_HttpReq(&admissions)
		if statusCode != http.StatusOK {
			w.WriteHeader(statusCode)
			return
		}

		pr("AdmissionsAVG: ", admissions)

		// Filter through admissions for employee's patients.
		durationAverage := 0.0 // Float64 is expensive (I think), but better that than doint float64 -> Int conversions inside the loop.
		for index, admission := range admissions {
			if sliceContains(patientIDs, admission.PatientID) { // Is this particular patient a part of this employee's patients.?
				dateAdmitted, _err := time.Parse(ADMISSION_DATE_LAYOUT, admission.AdmissionDate)
				dateDischarged, err := time.Parse(ADMISSION_DATE_LAYOUT, admission.DischargeDate)
				if err != nil || _err != nil {
					pr(_err, err)
					w.WriteHeader(http.StatusServiceUnavailable)
					return
				}
				duration := dateDischarged.Sub(dateAdmitted)
				if duration < 0 {
					pr("[Error] Admission Date appears lower than Discharge Date, index[", index, "]")
					w.WriteHeader(http.StatusServiceUnavailable)
					return
				}
				durationAverage += duration.Seconds()
			}
		}
		// Calculate AVERAGE.
		durationAverage = durationAverage / float64(len(patientIDs))

		formattedDurationAverage, err := time.ParseDuration(fmt.Sprint(durationAverage, "s"))
		if err != nil {
			pr(err)
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		}

		pr(formattedDurationAverage)

		sendJSONRespose(w, &struct {
			Seconds  float64
			Minutes  float64
			Hours    float64
			Days     float64
			Duration string
		}{
			Seconds:  formattedDurationAverage.Seconds(),
			Minutes:  formattedDurationAverage.Minutes(),
			Hours:    formattedDurationAverage.Hours(),
			Days:     formattedDurationAverage.Seconds() / 86400,
			Duration: formattedDurationAverage.String(),
		},
		)
		return
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}
