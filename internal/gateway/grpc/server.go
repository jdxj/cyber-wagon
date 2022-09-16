package grpc

import (
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/timestamppb"

	gateway "github.com/jdxj/cyber-wagon/internal/gateway/proto"
)

type Server struct {
	gateway.UnimplementedGRPCServer
}

func (s *Server) Transport(stream gateway.GRPC_TransportServer) error {
	for {
		cMsg, err := stream.Recv()
		if err != nil {
			logrus.Errorf("Transport Recv err: %t", err)
			return nil
		}
		logrus.Infof("client ts: %s", cMsg.GetTimestamp())

		err = stream.Send(&gateway.ServerMsg{
			Timestamp: timestamppb.Now(),
			Data:      nil,
		})
		if err != nil {
			logrus.Errorf("Transport Send err: %s", err)
		}
	}
}
