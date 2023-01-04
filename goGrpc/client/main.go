package main

import (
	"context"
	"fmt"
	"log"
	pb "userTask/user"

	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

var id string

//CreateUserProfile function to create a user profile
func CreateUserProfile(client pb.UserProfilesClient, user *pb.CreateUserProfileRequest) {
	resp, err := client.CreateUserProfile(context.Background(), user)

	if err != nil {
		log.Printf("Error %v", err)
	} else {
		log.Println(resp)
	}

	id = resp.GetId()
}

//GetUserProfile function to get a user profile
func GetUserProfile(client pb.UserProfilesClient, getReq *pb.GetUserProfileRequest) {

	r, err := client.GetUserProfile(context.Background(), getReq)

	if err != nil {
		log.Printf("Error %v", err)
	} else {
		fmt.Println(r)
	}
}

//UpdateUserProfile function to update the user profile
func UpdateUserProfile(client pb.UserProfilesClient, upUser *pb.UpdateUserProfileRequest) {

	r, err := client.UpdateUserProfile(context.Background(), upUser)

	if err != nil {
		log.Printf("Error %v", err)
	} else {
		fmt.Println(r)
	}
}

func main() {
	//Set up a connection to the gRPC server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewUserProfilesClient(conn)

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

	CreateUserProfile(client, user)

	getReq := &pb.GetUserProfileRequest{
		Id: id,
	}

	GetUserProfile(client, getReq)

	upUser := &pb.UpdateUserProfileRequest{
		UserProfile: &pb.UserProfile{
			Id:        id,
			Email:     "emailChange",
			FirstName: "abkjas",
			LastName:  "bcbc",
			//BirthDate: "jkasx",
			Telephones: []string{
				"asmlckma",
				"lkanclkas",
			},
		},
	}
	UpdateUserProfile(client, upUser)

}
