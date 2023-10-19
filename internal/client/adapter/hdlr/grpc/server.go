package grpc

import (
	"fmt"
	"net"


	"github.com/lavinas/keel/internal/client/adapter/hdlr/grpc/pb"
	"github.com/lavinas/keel/internal/client/core/port"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func StartServer(config port.Config, log port.Log, service port.Service) {
	// get port from config
	port, err := config.GetField("grpc", "port")
	if err != nil {
		log.Error(fmt.Sprintf("failed to get port: %v", err))
		panic(err)
	}
	// create grpc server
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	address := fmt.Sprintf("0.0.0.0:%s", port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Error(fmt.Sprintf("failed to listen: %v", err))
		panic(err)
	}
	log.Info(fmt.Sprintf("server listening at %v", listener.Addr()))
	// register grpc service
	pb.RegisterClientServiceServer(grpcServer, NewClientGRPCService(service))
	// start grpc server
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Error(fmt.Sprintf("failed to serve: %v", err))
		panic(err)
	}
}