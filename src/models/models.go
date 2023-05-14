package models

type Admission struct {
	Id            int
	AdmissionDate string
	DischargeDate string
	PatientID     int
}

type Allocation struct {
	Id          int
	AdmissionID int
	EmployeeID  int
	StartTime   string
	EndTime     string
}

type Employee struct {
	Id       int
	Surname  string
	Forename string
}

type Patient struct {
	Id        int
	Surname   string
	Forename  string
	NhsNumber string
}
