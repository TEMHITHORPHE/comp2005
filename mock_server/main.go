package main

import (
	"fmt"
	"net/http"
	Handlers "phxlabs/m/comp2005/mockserver/src/handlers"
	"time"
)

var pr = fmt.Println

func main() {
	pr("Mock Server Starting")

	http.HandleFunc("/", greet)

	http.HandleFunc("/Admissions", Handlers.Admissions)
	http.HandleFunc("/Admissions/", Handlers.An_Admission)

	http.HandleFunc("/Allocations", Handlers.Allocations)
	http.HandleFunc("/Allocations/", Handlers.An_Allocation)

	http.HandleFunc("/Employees", Handlers.Employees)
	http.HandleFunc("/Employees/", Handlers.An_Employee)

	http.HandleFunc("/Patients", Handlers.Patients)
	http.HandleFunc("/Patients/", Handlers.A_Patient)

	err := http.ListenAndServe(":7777", nil)
	pr("Server Started: ", err)

}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}
