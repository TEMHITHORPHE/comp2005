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
	Id          int
	AdmissionID int
	EmployeeID  int
	StartTime   string
	EndTime     string
}

func GetAllAllocations() Allocations {

	Allocations := Allocations{
		{
			Id:          1,
			AdmissionID: 1,
			EmployeeID:  4,
			StartTime:   "2020-11-28T16:45:00",
			EndTime:     "2020-11-28T23:56:00",
		},
		{
			Id:          2,
			AdmissionID: 3,
			EmployeeID:  4,
			StartTime:   "2021-09-23T21:50:00",
			EndTime:     "2021-09-24T09:50:00",
		},
	}
	fmt.Println("Models:", Allocations)
	return Allocations
}

type Employee struct {
	Id       int
	Surname  string
	Forename string
}

func GetAllEmployees() []Employee {

	Employees := []Employee{
		{
			Id:       1,
			Surname:  "Finley",
			Forename: "Sarah",
		},
		{
			Id:       2,
			Surname:  "Jackson",
			Forename: "Robert",
		},
		{
			Id:       3,
			Surname:  "Allen",
			Forename: "Alice",
		},
		{
			Id:       4,
			Surname:  "Jones",
			Forename: "Sarah",
		},
		{
			Id:       5,
			Surname:  "Wicks",
			Forename: "Patrick",
		},
		{
			Id:       6,
			Surname:  "Smith",
			Forename: "Alice",
		},
	}
	println(Employees)
	return Employees
}

type Patient struct {
	Id        int
	Surname   string
	Forename  string
	NhsNumber string
}

func GetAllPatients() []Patient {

	Patients := []Patient{
		{
			Id:        1,
			Surname:   "Robinson",
			Forename:  "Viv",
			NhsNumber: "1113335555",
		},
		{
			Id:        2,
			Surname:   "Carter",
			Forename:  "Heather",
			NhsNumber: "2224446666",
		},
		{
			Id:        3,
			Surname:   "Barnes",
			Forename:  "Nicky",
			NhsNumber: "6663338888",
		},
	}

	println("Models:", Patients)

	return Patients
}
