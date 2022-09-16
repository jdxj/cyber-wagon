package main

import (
	grpc2 "google.golang.org/grpc"

	grpc3 "github.com/jdxj/cyber-wagon/internal/gateway/grpc"
	gateway "github.com/jdxj/cyber-wagon/internal/gateway/proto"
	"github.com/jdxj/cyber-wagon/internal/pkg/network/grpc"
)

func main() {
	grpc.Start(":49152", func(registrar grpc2.ServiceRegistrar) {
		gateway.RegisterGRPCServer(registrar, &grpc3.Server{})
	})
}
