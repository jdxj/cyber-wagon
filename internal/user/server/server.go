package server

import (
	"context"

	"github.com/jdxj/cyber-wagon/internal/user/model"
	proto2 "github.com/jdxj/cyber-wagon/internal/user/proto"
)

type Server struct {
	proto2.UnimplementedUserServer
}

func (s *Server) SignUp(ctx context.Context, req *proto2.SignUpReq) (*proto2.SignUpRsp, error) {
	return model.SignUp(ctx, req)
}

func (s *Server) AuthBasic(ctx context.Context, req *proto2.AuthBasicReq) (*proto2.AuthBasicRsp, error) {
	return model.AuthBasic(ctx, req)
}
