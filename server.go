package main

import (
	"context"
	"log"
	"net"

	pb "interview/API"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedNetworkServiceServer
	userkeys []pb.UserKey
}

type userserver struct {
	pb.UnimplementedUserServiceServer
	user pb.User
}

const (
	port = ":50051"
)

func (s *server) getNetworkMembers(ctx context.Context, networkkey *pb.NetworkKey) (*pb.UserKeys, error) {
	var array []*pb.UserKey
	// var n int = 27
	log.Printf("Received: %v", networkkey.GetKey())
	for _, userkey := range s.userkeys {
		if userkey.Key == networkkey.GetKey() {
			array = append(array, &userkey)
		}
	}
	return &pb.UserKeys{
		Users: array,
	}, nil
	// return &pb.UserKeys{Users: make(map[]*pb.UserKey, )}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterNetworkServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
