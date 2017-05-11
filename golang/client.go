package main

import (
	"github.com/chengz0/grpcslide/golang/proto"
	"github.com/golang/glog"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
	if err != nil {
		glog.Fatalf("fail to dial: %v", err)
	}
	ctx := context.Background()
	client := proto.NewWeedProxyClient(conn)
	response, err := client.GetFileContent(ctx, &proto.WeedFileRequest{Fid: "1,go1a2b3c"})
	if err != nil {
		glog.Fatalln(err)
	}
	glog.Infof("Weed proxy client received: %s", response.FileContent)
}
