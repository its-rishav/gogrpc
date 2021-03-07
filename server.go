package main

import (
	"context"
	"fmt"

	pb "interview/API"
)

type server struct {
	pb.UnimplementedNetworkServiceServer
}

type userserver struct {
	pb.UnimplementedUserServiceServer
}

type networkData struct {
	nKey  int64
	nName string
	ukeys []*pb.UserKey
}

type userData struct {
	key  int64
	name string
}

func getNetworkData() []networkData {
	networkData := []networkData{
		{
			nKey:  12,
			nName: "tft",
			ukeys: []*pb.UserKey{
				{Key: 3},
				{Key: 6},
				{Key: 7},
			},
		},
		{
			nKey:  15,
			nName: "coderbhai",
			ukeys: []*pb.UserKey{
				{Key: 6},
				{Key: 8},
			},
		},
	}
	return networkData
}

const (
	port = ":50051"
)

func (s *server) GetNetworkMembers(ctx context.Context, networkkey *pb.NetworkKey) (*pb.UserKeys, error) {
	fmt.Println("Network Key received:", networkkey.Key)
	networkData := getNetworkData()
	for i := range networkData {
		if networkData[i].nKey == networkkey.Key {
			return &pb.UserKeys{Users: networkData[i].ukeys}, nil
		}
	}
	return nil, fmt.Errorf("No Network found for key: %v", networkkey.Key)
	// return &pb.UserKeys{Users: make(map[]*pb.UserKey, )}, nil
}

// func main() {
// 	fmt.Println("gRPC server running")
// 	lis, err := net.Listen("tcp", port)
// 	if err != nil {
// 		log.Fatalf("failed to listen: %v", err)
// 	}
// 	s := grpc.NewServer()
// 	pb.RegisterNetworkServiceServer(s, &server{})
// 	if err := s.Serve(lis); err != nil {
// 		log.Fatalf("failed to serve: %v", err)
// 	}

// }
