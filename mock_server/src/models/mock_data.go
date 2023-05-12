package models

import "fmt"

type Admission struct {
	Id            int
	AdmissionDate string
	DischargeDate string
	PatientID     int
}

func GetAllAdmissions() []Admission {

	Admissions := []Admission{
		{
			Id:            1,
			AdmissionDate: "2020-11-28T16:45:00",
			DischargeDate: "2020-11-28T23:56:00",
			PatientID:     2,
		},
		{
			Id:            2,
			AdmissionDate: "2020-12-07T22:14:00",
			DischargeDate: "0001-01-01T00:00:00",
			PatientID:     1,
		},
		{
			Id:            3,
			AdmissionDate: "2021-09-23T21:50:00",
			DischargeDate: "2021-09-27T09:56:00",
			PatientID:     2,
		},
	}

	fmt.Println("Model: ", Admissions)

	return Admissions
}

type Allocations []Allocation
type Allocation struct {
	id          int
	admissionID int
	employeeID  int
	startTime   string
	endTime     string
}

func GetAllAllocations() Allocations {

	Allocations := Allocations{
		{
			id:          1,
			admissionID: 1,
			employeeID:  4,
			startTime:   "2020-11-28T16:45:00",
			endTime:     "2020-11-28T23:56:00",
		},
		{
			id:          2,
			admissionID: 3,
			employeeID:  4,
			startTime:   "2021-09-23T21:50:00",
			endTime:     "2021-09-24T09:50:00",
		},
	}
	fmt.Println("Models:", Allocations)
	return Allocations
}

type Employee struct {
	id       int
	surname  string
	forename string
}

func GetAllEmployees() []Employee {

	Employees := []Employee{
		{
			id:       1,
			surname:  "Finley",
			forename: "Sarah",
		},
		{
			id:       2,
			surname:  "Jackson",
			forename: "Robert",
		},
		{
			id:       3,
			surname:  "Allen",
			forename: "Alice",
		},
		{
			id:       4,
			surname:  "Jones",
			forename: "Sarah",
		},
		{
			id:       5,
			surname:  "Wicks",
			forename: "Patrick",
		},
		{
			id:       6,
			surname:  "Smith",
			forename: "Alice",
		},
	}
	println(Employees)
	return Employees
}

type Patient struct {
	id        int
	surname   string
	forename  string
	nhsNumber string
}

func GetAllPatients() []Patient {

	Patients := []Patient{
		{
			id:        1,
			surname:   "Robinson",
			forename:  "Viv",
			nhsNumber: "1113335555",
		},
		{
			id:        2,
			surname:   "Carter",
			forename:  "Heather",
			nhsNumber: "2224446666",
		},
		{
			id:        3,
			surname:   "Barnes",
			forename:  "Nicky",
			nhsNumber: "6663338888",
		},
	}

	println("Models:", Patients)

	return Patients
}
