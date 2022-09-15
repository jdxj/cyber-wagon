package grpc

import (
	"io"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type NewClient func(grpc.ClientConnInterface)

func Connect(addr string, nc NewClient) (io.Closer, error) {
	cc, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	nc(cc)
	return cc, nil
}
