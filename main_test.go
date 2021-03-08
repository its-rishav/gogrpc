package main

import (
	"context"
	"fmt"
	pb "interview/API"
	"testing"
)

func TestGetNetworkMembers(t *testing.T) {
	s := server{}
	ctx := context.Background()
	key := &pb.NetworkKey{Key: 12}
	response, err := s.GetNetworkMembers(ctx, key)
	if err != nil {
		fmt.Errorf("err %v", key)
	}
	userCount := len(response.Users)

	if userCount != 3 {
		t.Errorf("expected: 2, got %v", userCount)
	}

}

func TestGetUser(t *testing.T) {
	userver := server{}
	ctx := context.Background()
	key := &pb.UserKey{Key: 6}
	response, err := userver.GetUser(ctx, key)
	if err != nil {
		fmt.Errorf("err %v", key)
	}
	fmt.Println(response)
	if response.Name != "saurabh" {
		t.Errorf("expected: saurabh, got %v", response.Name)
	}
}

func TestGetCommonContacts(t *testing.T) {
	contactserver := server{}
	ctx := context.Background()
	user1 := &pb.UserKey{Key: 3}
	user2 := &pb.UserKey{Key: 6}
	twoUsers := &pb.TwoUserKeys{User1: user1, User2: user2}
	response, err := contactserver.GetCommonContacts(ctx, twoUsers)
	if err != nil {
		fmt.Errorf("err %v", twoUsers)
	}
	commonContacts := len(response.Contacts)
	if commonContacts != 2 {
		t.Errorf("expected commoncontacts: 2, got %v", commonContacts)
	}
}

func TestGetSharedInterests(t *testing.T) {
	s := server{}
	ctx := context.Background()
	user1 := &pb.UserKey{Key: 7}
	user2 := &pb.UserKey{Key: 6}
	twoUsers := &pb.TwoUserKeys{User1: user1, User2: user2}
	response, err := s.GetSharedInterests(ctx, twoUsers)
	if err != nil {
		fmt.Errorf("err %v", twoUsers)
	}
	commonInterests := len(response.Interests)
	if commonInterests != 1 {
		t.Errorf("expected commonInterests: 2, got %v", commonInterests)
	}
}

func TestViewNetworkMembers(t *testing.T) {
	s := server{}
	ctx := context.Background()
	user := &pb.UserKey{Key: 3}
	network := &pb.NetworkKey{Key: 12}
	viewNetworkData := &pb.UserViewingNetwork{User: user, Network: network}
	response, err := s.ViewNetworkMembers(ctx, viewNetworkData)
	if err != nil {
		fmt.Errorf("err %v", err)
	}
	memberCount := len(response.Members)

	if memberCount != 2 {
		t.Errorf("expected members: 2, got: %v", memberCount)
	}
}
