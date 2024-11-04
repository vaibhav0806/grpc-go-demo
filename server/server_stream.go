package main

import (
	"log"

	pb "github.com/vaibhav0806/go-grpc-demo/proto"
)

func (s *helloServer) callSayHelloServerStreaming(req *pb.NamesList, stream pb.GreetService_SayHelloServerStreamingServer) error {
	log.Printf("Got request with names: %v", req.Names)
}
