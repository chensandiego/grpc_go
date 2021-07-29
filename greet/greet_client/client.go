package main

import (
	"context"
	"fmt"
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
	doUnary(c)
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
