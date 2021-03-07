package main

import (
	"context"
	pb "interview/API"
	"log"
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
	// userClient := pb.NewUserServiceClient(grpcConn)
	// contactClient := pb.NewContactServiceClient(grpcConn)
	// interestsClient := pb.NewInterestsServiceClient(grpcConn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := networkClient.GetNetworkMembers(ctx, &pb.NetworkKey{Key: 12})

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Members: %v", r.GetUsers())

	if err != nil {
		panic(err)
	}

	/*
		Register your service implimentation here, like so:

		RegisterViewNetworkServiceServer(grpcServer, --your server implimentation here--)
	*/

}
