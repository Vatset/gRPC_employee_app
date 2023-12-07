package functions

import (
	"encoding/json"
	"gRPC_employee_app"
	"gRPC_employee_app/pkg/gen/proto"
)

var (
	fileNameEmployees = "employees"
)

func GetAllEmployees() ([]*proto.EmployeesAnswerInfo, error) {

	data, err := GetData(fileNameEmployees)
	if err != nil {
		return nil, err
	}

	var dt gRPC_employee_app.EmployeeList

	err = json.Unmarshal(data, &dt)
	if err != nil {
		return nil, err
	}

	return dt.EmployeeData, nil
}

func MatchesFilter(filter *proto.EmployeesInfo, user *proto.EmployeesAnswerInfo) bool {
	if filter.Name != "" && filter.Name != user.DisplayName {
		return false
	}
	if filter.WorkPhone != "" && filter.WorkPhone != user.WorkPhone {
		return false
	}
	if filter.Id != 0 && filter.Id != user.Id {
		return false
	}
	if filter.Email != "" && filter.Email != user.Email {
		return false
	}
	return true
}
