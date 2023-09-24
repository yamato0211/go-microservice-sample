package main

import (
	"fmt"
	"go-micro-sample/todo/pkg/injection"
	"log"
	"net"
	"os"
	"os/signal"

	adapter "go-micro-sample/todo/pkg/adapter/grpc"
	todopb "go-micro-sample/todo/proto"

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

	usecase := injection.InitializeTodoUsecase()
	us := adapter.NewTodoServer(usecase)

	todopb.RegisterTodoServiceServer(s, us)

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
