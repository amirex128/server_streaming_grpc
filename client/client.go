package main

import (
	"context"
	"fmt"
	"github.com/amirex/server_streaming_grpc/rpc"
	"google.golang.org/grpc"
	"io"
	"log"
)
func showErr(e ...interface{}) {
	if e[0] != nil {
		log.Fatalln(e)
	}
}
func main() {
	dial, err := grpc.Dial("localhost:50051",grpc.WithInsecure())
	showErr(err)

	client := rpc.NewParsGoServiceClient(dial)

	product, err := client.ReceiveProduct(context.Background(), &rpc.ProductRequest{ProductId: 15455})
	showErr(err)

	for  {
		recv, err := product.Recv()
		if err == io.EOF{
			break
		}
		fmt.Println(recv.GetProduct().Name,recv.GetProduct().Price)
	}
}
