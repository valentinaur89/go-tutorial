package main

import (
	mypro "ashaik-dev/go-tutorial/grpc/myproto"
	"context"
	"fmt"
	"log"
	"net"

	_ "github.com/jnewmano/grpc-json-proxy/codec"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) MyGreetings(ctx context.Context, req *mypro.MyRequest) (*mypro.MyResponse, error) {
	str := fmt.Sprintf("Happy new year %v!", req.Name)
	return &mypro.MyResponse{Greetings: str}, nil
}
func main() {
	addr := "127.0.0.1:8083"
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("error whil listening ", err)
	}
	fmt.Println("Server running on ...", addr)
	s := grpc.NewServer()
	mypro.RegisterMyServiceServer(s, &server{})
	s.Serve(lis)
}
