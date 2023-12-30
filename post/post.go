package post

import (
	"log"

	"golang.org/x/net/context"
)

type Server struct {
}

func (s *Server) Post(ctx context.Context, in *PostReq) (*PostRes, error) {
	log.Printf("Receive message body from client: %s", in.Content)
	return &PostRes{Message: "Hello From the Server!"}, nil
}