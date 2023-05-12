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

	http.HandleFunc("/add", Handlers.Admissions)
	// http.HandleFunc("Admissions/{id}", Handlers.An_Admission)

	http.HandleFunc("/Allocations", Handlers.Allocations)
	// http.HandleFunc("Allocations/{id}", Handlers.An_Allocation)

	http.HandleFunc("/Employees", Handlers.Employees)
	// http.HandleFunc("Employees/{id}", Handlers.An_Employee)

	http.HandleFunc("/Patients", Handlers.Patients)
	// http.HandleFunc("Patients/{id}", Handlers.A_Patient)

	err := http.ListenAndServe(":8090", nil)
	pr("Server Started: ", err)

}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}
