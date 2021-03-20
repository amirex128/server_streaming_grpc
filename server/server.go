package main

import (
	"fmt"
	"github.com/amirex/server_streaming_grpc/rpc"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

var p = fmt.Println

type ServerParsGo struct {
	rpc.UnimplementedParsGoServiceServer
}

func (*ServerParsGo) ReceiveProduct(request *rpc.ProductRequest, stream rpc.ParsGoService_ReceiveProductServer) error {
	p("we receive product id " + string(request.GetProductId()))

	for i := range time.Tick(1 * time.Second) {
		stream.Send(&rpc.ProductResponse{Product: &rpc.Product{
			Name:  "lap top asus",
			Price: 4000 * uint32(i.Second()),
		}})
	}

	return nil
}

func showErr(e ...interface{}) {
	if e[0] != nil {
		log.Fatalln(e)
	}
}
func main() {
	listen, err := net.Listen("tcp", "0.0.0.0:50051")
	showErr(err)

	server := grpc.NewServer()
	rpc.RegisterParsGoServiceServer(server, &ServerParsGo{})
	server.Serve(listen)

}
