package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "interview/API"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedContactServiceServer
	pb.UnimplementedNetworkServiceServer
	pb.UnimplementedViewNetworkServiceServer
}

type networkData struct {
	nKey  int64
	nName string
	ukeys []*pb.UserKey
}

type userData struct {
	key       int64
	name      string
	contacts  []string
	interests []string
}

func getNetworkData() []networkData {
	networkData := []networkData{
		{
			nKey:  12,
			nName: "tft",
			ukeys: []*pb.UserKey{
				{Key: 3},
				{Key: 6},
				{Key: 8},
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

func getUserData() []userData {
	userData := []userData{
		{key: 3, name: "rishav", contacts: []string{"saurabh", "rajesh", "surabhi"}, interests: []string{"table tennis", "basketball"}},
		{key: 6, name: "saurabh", contacts: []string{"rishav", "rajesh", "surabhi"}, interests: []string{"table tennis", "football"}},
		{key: 7, name: "vijay", contacts: []string{"rishav"}, interests: []string{"football"}},
		{key: 8, name: "rajesh", contacts: []string{"saurabh", "rishav"}, interests: []string{"table tennis", "basketball", "football"}},
	}
	return userData
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

func (s *server) GetUser(ctx context.Context, ukey *pb.UserKey) (*pb.User, error) {
	fmt.Println("User key received: ", ukey.Key)
	userData := getUserData()
	for i := range userData {
		if userData[i].key == ukey.Key {
			return &pb.User{Key: userData[i].key, Name: userData[i].name}, nil
		}
	}
	return nil, fmt.Errorf("No Network found for key: %v", ukey.Key)

}

func (s *server) GetCommonContacts(ctx context.Context, uKeys *pb.TwoUserKeys) (*pb.Contacts, error) {
	fmt.Printf("Recieved user keys user1: %v, user2: %v\n", uKeys.User1.Key, uKeys.User2.Key)
	user1 := uKeys.GetUser1()
	user2 := uKeys.User2
	userData := getUserData()
	u1Contacts, u2Contacts := []string{}, []string{}
	commonContacts := []string{}
	for i := range userData {
		if userData[i].key == user1.Key {
			u1Contacts = append(u1Contacts, userData[i].contacts...)
		} else if userData[i].key == user2.Key {
			u2Contacts = append(u2Contacts, userData[i].contacts...)
		}
	}
	for u1i := range u1Contacts {
		for u2i := range u2Contacts {
			if u1Contacts[u1i] == u2Contacts[u2i] {
				commonContacts = append(commonContacts, u1Contacts[u1i])
			}
		}
	}
	return &pb.Contacts{Contacts: commonContacts}, nil
}

func (s *server) GetSharedInterests(ctx context.Context, uKeys *pb.TwoUserKeys) (*pb.Interests, error) {
	fmt.Printf("Recieved user keys user1: %v, user2: %v\n", uKeys.User1.Key, uKeys.User2.Key)
	user1 := uKeys.GetUser1()
	user2 := uKeys.User2
	userData := getUserData()
	u1Interests, u2Interests := []string{}, []string{}
	commonInterests := []string{}
	for i := range userData {
		if userData[i].key == user1.Key {
			u1Interests = append(u1Interests, userData[i].interests...)
		} else if userData[i].key == user2.Key {
			u2Interests = append(u2Interests, userData[i].interests...)
		}
	}
	for u1i := range u1Interests {
		for u2i := range u2Interests {
			if u1Interests[u1i] == u2Interests[u2i] {
				commonInterests = append(commonInterests, u1Interests[u1i])
			}
		}
	}
	return &pb.Interests{Interests: commonInterests}, nil
}

func (s *server) ViewNetworkMembers(ctx context.Context, in *pb.UserViewingNetwork) (*pb.NetworkMembersView, error) {
	fmt.Printf("Recieved user: %v, checking details of network:%v\n", in.User, in.Network)
	user := in.GetUser()
	network := in.GetNetwork()
	userkeys, err := s.GetNetworkMembers(ctx, network)
	if err != nil {
		fmt.Errorf("err %v", err)
	}
	users := userkeys.Users
	memberList := []*pb.MemberDetails{}
	contains := false // check if the user key is present in network

	for j := range users {
		if users[j].Key == user.Key {
			contains = true
			break
		}
	}
	if contains {
		for i := range users {
			if users[i].Key != user.Key {
				twoUsers := &pb.TwoUserKeys{User1: users[i], User2: user}
				commonContacts, _ := s.GetCommonContacts(ctx, twoUsers)
				fmt.Println("commoncontact", commonContacts)
				commonInterests, _ := s.GetSharedInterests(ctx, twoUsers)
				userDetail, _ := s.GetUser(ctx, users[i])
				memberList = append(memberList, &pb.MemberDetails{User: userDetail, CommonContacts: commonContacts, CommonInterests: commonInterests})
			}
		}
		fmt.Println("Member List", memberList)
		return &pb.NetworkMembersView{Members: memberList}, nil
	} else {
		return &pb.NetworkMembersView{Members: []*pb.MemberDetails{}}, fmt.Errorf("User is not present in Network")
	}
}

func main() {
	fmt.Println("gRPC server running")
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterViewNetworkServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
