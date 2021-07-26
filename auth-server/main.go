package main

import (
	"auth-server/token"
	"log"
	"net"
	"os"

	auth_pb "github.com/hasbiasshidiq/auth-stub-5"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
		log.Println("using default environment variable")
	}

	port := os.Getenv("API_PORT")

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
