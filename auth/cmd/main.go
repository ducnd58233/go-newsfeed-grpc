package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/ducnd58233/go-newsfeed-grpc/auth/configs"
	"github.com/ducnd58233/go-newsfeed-grpc/auth/internal/app/service"
	"github.com/ducnd58233/go-newsfeed-grpc/auth/pkg/types/proto/pb/auth_pb"
	"google.golang.org/grpc"
)

var (
	path = flag.String("conf", "config.yml", "config path for auth service")
)

func main() {
	flag.Parse()
	conf, err := configs.GetAuthConfig(*path)
	if err != nil {
		log.Fatalf("failed to parse config: %v", err)
	}
	svc, err := service.NewAuthService(conf)
	if err != nil {
		log.Fatalf("failed to init auth service: %v", err)
	}
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", conf.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	auth_pb.RegisterAuthServer(grpcServer, svc)
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("server stopped %v", err)
	}
}