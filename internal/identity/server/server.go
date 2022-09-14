package server

import (
	"context"

	"github.com/bwmarrin/snowflake"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/jdxj/cyber-wagon/internal/identity/model"
	identity "github.com/jdxj/cyber-wagon/internal/identity/proto"
)

const ()

func NewServer(nodeID int64) *Server {
	node, err := snowflake.NewNode(nodeID)
	if err != nil {
		logrus.Fatalln(err)
	}

	return &Server{
		node: node,
	}
}

type Server struct {
	identity.UnimplementedIdentityServer

	node *snowflake.Node
}

func (s *Server) UniqueID(ctx context.Context, req *identity.UniqueIDReq) (*identity.UniqueIDRsp, error) {
	rsp := &identity.UniqueIDRsp{}
	for i := int64(0); i < req.GetQuantity(); i++ {
		rsp.Ids = append(rsp.Ids, uint64(s.node.Generate().Int64()))
	}
	return rsp, nil
}

func (s *Server) IncrementID(ctx context.Context, req *identity.IncrementIDReq) (*identity.IncrementIDRsp, error) {
	if req.GetNamespace() == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "invalid namespace")
	}
	return model.IncrementID(ctx, req)
}
