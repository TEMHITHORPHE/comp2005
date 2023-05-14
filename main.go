package main

import (
	"fmt"
	"net/http"
	handlers "phxlabs/m/comp2005/automated_testing/src/handlers"
	"time"
)

var pr = fmt.Println

func main() {
	pr("Welcome ::: Report Automated Software Testing ... ")
	pr("Server Starting ... ")

	http.HandleFunc("/", greet)

	http.HandleFunc("/employee/patients/", handlers.EmployeePatients)
	http.HandleFunc("/patients/admissions/duration/3", handlers.PatientsAdmissionDuration)
	http.HandleFunc("/patients/admissions/maxday/", handlers.DayOfMaximumAdmissions)
	http.HandleFunc("/employee/patients/average/duration/", handlers.EmployeePatientsDurationAverage)

	http.ListenAndServe(":8080", nil)
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}
