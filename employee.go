package gRPC_employee_app

import "gRPC_employee_app/pkg/gen/proto"

type EmployeeList struct {
	EmployeeData []*proto.EmployeesAnswerInfo `json:"data"`
}

type EmployeeAbsencesListProto struct {
	EmployeeAbsencesData []*proto.EmployeesAbsencesInfo `json:"data"`
}

type EmployeeAbsencesList struct {
	EmployeeAbsencesData []EmployeeAbsence `json:"data"`
}

type Employee struct {
	ID          int    `json:"id"`
	DisplayName string `json:"displayName"`
	Email       string `json:"email"`
	MobilePhone string `json:"mobilePhone"`
	WorkPhone   string `json:"workPhone"`
}

type EmployeeAbsence struct {
	CreatedDate string `json:"createdDate"`
	DateFrom    string `json:"dateFrom"`
	DateTo      string `json:"dateTo"`
	Id          int64  `json:"id"`
	PersonId    int64  `json:"personId"`
	ReasonId    int64  `json:"reasonId"`
}
