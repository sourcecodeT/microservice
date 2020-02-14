package main

import (
	"context"
	"greet/greetpb"
	"log"

	"google.golang.org/grpc"
)

func main() {
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	failOnError(err, "Could not connect")
	defer connection.Close()

	client := greetpb.NewGreetServiceClient(connection)
	log.Printf("[INFO] Client created")

	doUnary(client)
}

func doUnary(client greetpb.GreetServiceClient) {
	log.Println("[INFO] Starting to do a Unary RPC ...")
	req := &greetpb.GreetingRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Thanh",
			LastName:  "Dao Ngoc",
		},
	}
	res, err := client.Greet(context.Background(), req)
	failOnError(err, "err while calling Greet rpc")

	log.Printf("[INFO] Response from Greet: %v", res.Result)
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("[ERROR] %s: %v", msg, err)
	}
}