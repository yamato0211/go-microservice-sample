package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	adapter "go-micro-sample/user/pkg/adapter/grpc"
	"go-micro-sample/user/pkg/injection"

	userpb "go-micro-sample/user/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	port := 8000
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()

	usecase := injection.InitializeUserUsecase()
	us := adapter.NewUserServer(usecase)

	userpb.RegisterUserServiceServer(s, us)

	reflection.Register(s)

	go func() {
		log.Printf("start gRPC server port: %v", port)
		if err := s.Serve(listener); err != nil {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("stopping gRPC server...")
	s.GracefulStop()
}
