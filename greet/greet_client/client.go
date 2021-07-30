package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/grpc-go-course/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("hello client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not connect : %v", err)

	}
	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)
	//doUnary(c)

	doServerStreaming(c)
}

func doUnary(c greetpb.GreetServiceClient) {
	fmt.Println("Start to unary")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{

			FirstName: "Chen",
			LastName:  "Ming",
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling greet rpc: %v", err)
	}
	log.Printf("Response from greet: %v", res.Result)
	//fmt.Printf("cretaed client: %f", c)
}

func doServerStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("start to do a streaming")

	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Chen",
			LastName:  "Ming",
		},
	}
	resStream, err := c.GreetManytimes(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling greetmanytime RPC: %v", err)
	}
	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error while reading stream %v", err)
		}
		log.Printf("response from greetmanytimes: %v", msg.GetResult())
	}

}
