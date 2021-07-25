package main

import (
	"auth-server/token"
	"log"
	"net"

	auth_pb "github.com/hasbiasshidiq/auth-stub-5"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	port := "50059"

	l, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalln("Failed to listen")
	}

	s := token.Server{}

	grpcServer := grpc.NewServer()

	reflection.Register(grpcServer)

	auth_pb.RegisterAuthServer(grpcServer, &s)

	log.Print("gRPC server started at ", port)
	if err := grpcServer.Serve(l); err != nil {
		log.Fatal("Failed to serve")
	}

}
