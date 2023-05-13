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

	http.HandleFunc("/admissions", Handlers.Admissions)
	http.HandleFunc("/admission/", Handlers.An_Admission)

	http.HandleFunc("/allocations", Handlers.Allocations)
	http.HandleFunc("/allocation/", Handlers.An_Allocation)

	http.HandleFunc("/employees", Handlers.Employees)
	http.HandleFunc("/employee/", Handlers.An_Employee)

	http.HandleFunc("/patients", Handlers.Patients)
	http.HandleFunc("/patient/", Handlers.A_Patient)

	err := http.ListenAndServe(":7777", nil)
	pr("Server Started: ", err)

}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}
