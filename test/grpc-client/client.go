package main

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
	grpc2 "google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"

	gateway "github.com/jdxj/cyber-wagon/internal/gateway/proto"
	"github.com/jdxj/cyber-wagon/internal/pkg/network/grpc"
)

func main() {
	var client gateway.GatewayClient
	closer, err := grpc.Connect("localhost:49152", func(cc grpc2.ClientConnInterface) {
		client = gateway.NewGatewayClient(cc)
	})
	if err != nil {
		logrus.Fatalln(err)
	}
	defer closer.Close()

	tc, err := client.Transport(context.Background())
	if err != nil {
		logrus.Fatalln(err)
	}

	for {
		err = tc.Send(&gateway.ClientMsg{
			Uuid:       "abc",
			Device:     "def",
			Timestamp:  timestamppb.Now(),
			AppVersion: "ghi",
			Data:       nil,
		})
		if err != nil {
			logrus.Errorf("client transport Send err: %s", err)
			continue
		}

		rsp, err := tc.Recv()
		if err != nil {
			logrus.Errorf("client transport Recv err: %s", err)
			continue
		}
		logrus.Infof("rsp: %s", rsp.GetTimestamp())
		time.Sleep(time.Second * 2)
	}
}
