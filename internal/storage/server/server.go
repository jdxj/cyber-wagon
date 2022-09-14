package server

import storage "github.com/jdxj/cyber-wagon/internal/storage/proto"

type Server struct {
	storage.UnimplementedStorageServer
}

func (s *Server) Upload(sus storage.Storage_UploadServer) error {
	c := storage.NewStorageClient()
	cc, _ := c.Upload()
	cc.Send()
	req, _ := sus.Recv()
	req.
}