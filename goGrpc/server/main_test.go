package main

import (
	"context"
	"fmt"
	"testing"
	pb "userTask/user"

	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func TestCreateUserProfile(t *testing.T) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		t.Errorf("did not connect: %v", err)
	}
	defer conn.Close()

	user := &pb.CreateUserProfileRequest{
		UserProfile: &pb.UserProfile{
			Email:     "email",
			FirstName: "abkjas",
			LastName:  "bcbc",
			//BirthDate: "jkasx",
			Telephones: []string{
				"asmlckma",
				"lkanclkas",
			},
		},
	}

	client := pb.NewUserProfilesClient(conn)

	resp, err := client.CreateUserProfile(context.Background(), user)

	if err != nil {
		t.Errorf("Error %v", err)
	} else {
		t.Log(resp)
	}
}

func TestGetUserProfile(t *testing.T) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		t.Errorf("did not connect: %v", err)
	}
	defer conn.Close()
	ids := []string{
		"7dd3f09e-7833-4bbe-8841-753a0beb1794",
		"2a2329d4-b6f5-4b5d-9c2f-2cd8c9ecf709",
		"8df069a2-e796-4de3-b8fd-2356b5f66973",
		"8f433107-c00e-443a-982c-bbf1bfcaa13d",
	}
	client := pb.NewUserProfilesClient(conn)
	for _, id := range ids {
		getReq := &pb.GetUserProfileRequest{
			Id: id,
		}

		resp, err := client.GetUserProfile(context.Background(), getReq)

		if err != nil {
			t.Error(err)
		} else {
			fmt.Print(resp)
		}
	}
}

func TestUpdateUserProfile(t *testing.T) {

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		t.Errorf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewUserProfilesClient(conn)

	users := []pb.UserProfile{
		{
			Id:        "7dd3f09e-7833-4bbe-8841-753a0beb1794",
			Email:     "abc@a.com",
			FirstName: "V",
			LastName:  "J",
			Telephones: []string{
				"560245250",
				"062598756",
			},
		},
		{
			Id:        "2a2329d4-b6f5-4b5d-9c2f-2cd8c9ecf709",
			Email:     "abc@a.com",
			FirstName: "V",
			LastName:  "J",
			Telephones: []string{
				"560245250",
				"062598756",
			},
		},
		{
			Id:        "8df069a2-e796-4de3-b8fd-2356b5f66973",
			Email:     "abc@a.com",
			FirstName: "V",
			LastName:  "J",
			Telephones: []string{
				"560245250",
				"062598756",
			},
		},
		{
			Id:        "8f433107-c00e-443a-982c-bbf1bfcaa13d",
			Email:     "abc@a.com",
			FirstName: "V",
			LastName:  "J",
			Telephones: []string{
				"560245250",
				"062598756",
			},
		},
		{
			Id:        "f9cdb083-aeab-48da-9d23-0b16077890a1",
			Email:     "s@s.com",
			FirstName: "S",
			LastName:  "S",
			Telephones: []string{
				"56562220",
				"665465465",
			},
		},
	}

	for _, user := range users {
		upReq := &pb.UpdateUserProfileRequest{
			UserProfile: &user,
		}

		resp, err := client.UpdateUserProfile(context.Background(), upReq)

		if err != nil {
			t.Error(err)
		} else {
			fmt.Println(resp)
		}
	}
}

func TestListUsersProfile(t *testing.T) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		t.Errorf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewUserProfilesClient(conn)

	queries := []string{
		"V",
		"ab",
		"ss",
	}

	for _, query := range queries {
		req := &pb.ListUsersProfilesRequest{
			Query: query,
		}

		resp, err := client.ListUsersProfiles(context.Background(), req)

		if err != nil {
			t.Error(err)
		} else {
			fmt.Println("for querry ", query)
			fmt.Println(resp.Profiles)
		}
	}
}
