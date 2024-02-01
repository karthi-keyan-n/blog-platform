package main

import (
	"github.com/blog-platform/handler"
	bp "github.com/blog-platform/proto/go/blogplatformservice"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	bp.RegisterBlogPlatformServiceServer(s, handler.NewBlogPostService())
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
