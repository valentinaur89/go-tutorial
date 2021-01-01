package main

import (
	"context"
	"fmt"
	"log"

	mypro "ashaik-dev/go-tutorial/grpc/myproto"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("client...")
	opt := grpc.WithInsecure()
	cc, err := grpc.Dial("localhost:8083", opt)
	if err != nil {
		log.Fatalf("error connecting to server %v", err)
	}
	defer cc.Close()
	client := mypro.NewMyServiceClient(cc)
	req := &mypro.MyRequest{Name: "ashaik"}
	res, err := client.MyGreetings(context.Background(), req)
	if err != nil {
		log.Fatalf("error while reading response %v", err)
	}
	fmt.Println(res)
}
