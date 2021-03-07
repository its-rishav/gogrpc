package main

import (
	"context"
	"fmt"
	pb "interview/API"
	"testing"
)

func TestGetNetworkMembers(t *testing.T) {
	s := Server{}
	ctx := context.Background()
	key := &pb.NetworkKey{Key: 12}
	response, err := s.GetNetworkMembers(ctx, key)
	if err != nil {
		fmt.Errorf("err %v", key)
	}
	userCount := len(response.Users)

	if userCount != 2 {
		t.Errorf("expected: 2, got %v", userCount)
	}
	// userCount := len(response.users)
	// if userCount != 2 {
	// 	t.Error("Expected 2, got ", userCount)
	// }

}
