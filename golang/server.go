package main

import (
	"fmt"
	"net"

	"github.com/chengz0/grpcslide/golang/proto"
	"github.com/golang/glog"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		glog.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	defer server.Stop()

	weedservice := &WeedService{}
	proto.RegisterWeedProxyServer(server, weedservice)

	server.Serve(listen)
}

//
type WeedService struct {
}

func (p *WeedService) GetFileContent(ctx context.Context, req *proto.WeedFileRequest) (*proto.WeedFileResponse, error) {
	return &proto.WeedFileResponse{FileContent: fmt.Sprintf("hello, %s!", req.Fid)}, nil
}
