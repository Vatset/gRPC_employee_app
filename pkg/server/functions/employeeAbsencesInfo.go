package functions

import (
	"encoding/json"
	"fmt"
	"gRPC_employee_app"
	"gRPC_employee_app/pkg/gen/proto"
	"os"
	"time"
)

var (
	layout           = "2006-01-02T15:04:05"
	fileNameAbsences = "absences"
)

func GetAllEmployeesAbsence(req *proto.EmployeesAbsencesRequest) ([]*proto.EmployeesAbsencesInfo, error) {
	data, err := GetData(fileNameAbsences)
	if err != nil {
		return nil, err
	}
	var dataProto gRPC_employee_app.EmployeeAbsencesListProto

	err = json.Unmarshal(data, &dataProto)
	if err != nil {
		return nil, err
	}

	var filteredUsers []*proto.EmployeesAbsencesInfo
	for _, user := range dataProto.EmployeeAbsencesData {
		if GetEmployeesAbsenceByDate(req.EmployeesAbsencesInfo.DateFrom, req.EmployeesAbsencesInfo.DateTo, user) {
			filteredUsers = append(filteredUsers, user)

		}
	}
	return filteredUsers, nil

}

func GetData(name string) ([]byte, error) {
	filename := fmt.Sprintf("info_data/%s.json", name)
	data, err := os.ReadFile(filename)
	if err != nil {
		return []byte(""), err
	}
	return data, nil
}

func GetEmployeesAbsenceByDate(DateFrom, DateTo string, user *proto.EmployeesAbsencesInfo) bool {

	fromTime, err := time.Parse(layout, DateFrom)

	if err != nil {

		return false
	}

	toTime, err := time.Parse(layout, DateTo)

	if err != nil {

		return false
	}

	userFromTime, err := time.Parse(layout, user.DateFrom)

	if err != nil {

		return false
	}

	userToTime, err := time.Parse(layout, user.DateTo)

	if err != nil {

		return false
	}

	if userToTime.After(fromTime) && userFromTime.Before(toTime) {
		return true
	} else {

	}
	return false
}

func GetEmployeesAbsenceByValue(filter *proto.EmployeesAbsencesRequest, employee *proto.EmployeesAbsencesInfo) bool {
	if filter.EmployeesAbsencesInfo.Id != 0 && filter.EmployeesAbsencesInfo.Id != employee.Id {
		return false
	}
	if filter.EmployeesAbsencesInfo.PersonId != 0 && filter.EmployeesAbsencesInfo.PersonId != employee.PersonId {
		return false
	}
	if filter.EmployeesAbsencesInfo.CreatedDate != "" && filter.EmployeesAbsencesInfo.CreatedDate != employee.CreatedDate {
		return false
	}
	if filter.EmployeesAbsencesInfo.ReasonId != 0 && filter.EmployeesAbsencesInfo.ReasonId != employee.ReasonId {
		return false
	}
	return true
}
