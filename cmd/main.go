package main

import (
	_ "fmt"
	app "gRPC_employee_app/pkg/gen/proto"
	"gRPC_employee_app/pkg/server"
	_ "github.com/jackc/pgx/stdlib"
	_ "github.com/joho/godotenv"
	_ "github.com/sirupsen/logrus"
	_ "github.com/spf13/viper"
	"google.golang.org/grpc"
	"log"
	"net"
	_ "os"
)

func main() {
	s := grpc.NewServer()
	srv := &server.GRPCServer{}

	app.RegisterAppServiceServer(s, srv)
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}

}
