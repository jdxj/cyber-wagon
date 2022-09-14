package server

import gateway "github.com/jdxj/cyber-wagon/internal/gateway/proto"

type Server struct {
	gateway.UnimplementedGatewayServer
}
