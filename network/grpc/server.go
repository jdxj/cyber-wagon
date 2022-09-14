package grpc

import (
	"net"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type RegisterService func(grpc.ServiceRegistrar)

func Start(addr string, rs RegisterService) {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		logrus.Fatalf("%s", err)
	}
	// 可以利用 opt 建立 tls
	gs := grpc.NewServer()
	rs(gs)
	err = gs.Serve(lis)
	if err != nil {
		logrus.Fatalf("%s", err)
	}
}
