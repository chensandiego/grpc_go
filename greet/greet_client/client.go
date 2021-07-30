package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

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

	//doServerStreaming(c)
	doClientStreaming(c)
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

func doClientStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("start to do a client streaming")

	requests := []*greetpb.LongGreetRequest{
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Chen",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "ming",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "leo",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "jeff",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "ron",
			},
		},
	}
	stream, err := c.LongGreet(context.Background())

	if err != nil {
		log.Fatalf("error while reading stream %v", err)
	}

	for _, req := range requests {
		fmt.Printf("sending req:%v\n", req)
		stream.Send(req)
		time.Sleep(1000 * time.Millisecond)
	}
	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatal("error while receiving response from longgreet:%v", err)
	}
	fmt.Printf("longgreet response: %v\n", res)
}
