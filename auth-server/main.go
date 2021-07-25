package main

import (
	"auth-server/token"
	"log"
	"net"

	"github.com/hasbiasshidiq/auth-stub-3"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// log.SetFormatter(&log.TextFormatter{
	// 	FullTimestamp: true,
	// })

	// var port string
	// var ok bool
	// port, ok = os.LookupEnv("PORT")
	// if ok {
	// 	log.WithFields(log.Fields{
	// 		"PORT": port,
	// 	}).Info("PORT env var defined")

	// } else {
	// 	port = "9000"
	// 	log.WithFields(log.Fields{
	// 		"PORT": port,
	// 	}).Warn("PORT env var not defined. Going with default")

	// }

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
