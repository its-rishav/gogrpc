package main

import (
	"context"
	pb "interview/API"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
)

func main() {
	//The service that provide the data runs on a different port, as part of a different application
	grpcConn, err := grpc.Dial(":50051", grpc.WithInsecure())

	if err != nil {
		panic(err)
	}

	networkClient := pb.NewNetworkServiceClient(grpcConn)
	userClient := pb.NewUserServiceClient(grpcConn)
	contactClient := pb.NewContactServiceClient(grpcConn)
	interestsClient := pb.NewInterestsServiceClient(grpcConn)

	lis, err := net.Listen("tcp", ":50051")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	r, err := networkClient.NetworkService(ctx, &pb.NetworkKey{Key: 12311})

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Members: %s", r.UserKeys())

	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()

	/*
		Register your service implimentation here, like so:

		RegisterViewNetworkServiceServer(grpcServer, --your server implimentation here--)
	*/
	err = grpcServer.Serve(lis)

	if err != nil {
		panic(err)
	}

}
