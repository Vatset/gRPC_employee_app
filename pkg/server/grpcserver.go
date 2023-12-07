package server

import (
	"context"
	"gRPC_employee_app/pkg/gen/proto"
	app "gRPC_employee_app/pkg/gen/proto"
	"gRPC_employee_app/pkg/server/functions"
	_ "github.com/jackc/pgx/stdlib"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GRPCServer struct {
	app.UnimplementedAppServiceServer
}

func (s *GRPCServer) GetEmployeesList(ctx context.Context, req *proto.EmployeesRequest) (*proto.EmployeesAnswer, error) {
	data, err := functions.GetAllEmployees()
	if err != nil {
		return nil, err
	}
	if req.EmployeeInfo.Id == 0 && req.EmployeeInfo.WorkPhone == "" && req.EmployeeInfo.Name == "" && req.EmployeeInfo.Email == "" {
		return &proto.EmployeesAnswer{
			Status: status.New(codes.OK, "").String(),
			Data:   data,
		}, nil

	} else {
		var filteredUsers []*proto.EmployeesAnswerInfo
		for _, user := range data {
			if functions.MatchesFilter(req.EmployeeInfo, user) {
				filteredUsers = append(filteredUsers, user)
			}
		}

		return &proto.EmployeesAnswer{
			Status: status.New(codes.OK, "").String(),
			Data:   filteredUsers,
		}, nil
	}
}

func (s *GRPCServer) GetEmployeeAbsencesInfo(ctx context.Context, req *proto.EmployeesAbsencesRequest) (*proto.EmployeesAbsencesAnswer, error) {

	data, err := functions.GetAllEmployeesAbsence(req)
	if err != nil {
		return nil, err
	}

	if req.EmployeesAbsencesInfo.Id == 0 && req.EmployeesAbsencesInfo.PersonId == 0 && req.EmployeesAbsencesInfo.CreatedDate == "" && req.EmployeesAbsencesInfo.ReasonId == 0 {
		return &proto.EmployeesAbsencesAnswer{
			Status: status.New(codes.OK, "").String(),
			Data:   data,
		}, nil

	} else {
		var filteredUsers []*proto.EmployeesAbsencesInfo
		for _, user := range data {
			if functions.GetEmployeesAbsenceByValue(req, user) {
				filteredUsers = append(filteredUsers, user)
			}
		}

		return &proto.EmployeesAbsencesAnswer{
			Status: status.New(codes.OK, "").String(),
			Data:   filteredUsers,
		}, nil
	}
}
