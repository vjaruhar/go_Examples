package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"
	pb "userTask/user"

	"database/sql"

	"google.golang.org/grpc"

	"github.com/lib/pq"
	_ "github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
)

type server struct {
	users []*pb.UserProfile
}

//TODO: Handle errors on client and server side
var resp []*pb.UserProfile

const (
	port       = ":50051"
	DbUser     = "postgres"
	DbPassword = ""
	DbName     = "user"
)

func (s *server) CreateUserProfile(ctx context.Context, req *pb.CreateUserProfileRequest) (*pb.UserProfile, error) {
	s.users = append(s.users, req.UserProfile)

	id := uuid.Must(uuid.NewV4())

	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DbUser, DbPassword, DbName)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		log.Print(err)
		return &pb.UserProfile{}, err
	}
	sqlStatement := `INSERT INTO public."user"(
		id, email, first_name, last_name, birth_date, telephones)
		VALUES ($1, $2, $3, $4, $5, $6);`

	_, err = db.Exec(sqlStatement, id, req.UserProfile.GetEmail(), req.UserProfile.GetFirstName(), req.UserProfile.GetLastName(), time.Now(), pq.Array(req.UserProfile.GetTelephones()))

	if err != nil {
		log.Print(err)
		return &pb.UserProfile{}, err
	}

	defer db.Close()

	return &pb.UserProfile{
		Id:    id.String(),
		Email: req.UserProfile.Email,
	}, nil
}

func (s *server) UpdateUserProfile(ctx context.Context, req *pb.UpdateUserProfileRequest) (*pb.UserProfile, error) {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DbUser, DbPassword, DbName)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		panic(err)
	}
	sqlStatement := `UPDATE public."user"
	SET email=$1, first_name=$2, last_name=$3, birth_date=$4, telephones=$5
	WHERE id=$6;`

	v, err := db.Exec(sqlStatement, req.UserProfile.GetEmail(), req.UserProfile.GetFirstName(), req.UserProfile.GetLastName(), time.Now(), pq.Array(req.UserProfile.GetTelephones()), req.UserProfile.GetId())

	if err != nil {
		log.Println(err)
	} else {
		count, err := v.RowsAffected()
		if count == 0 {
			log.Printf("Error id not found--- %v", err)
			return &pb.UserProfile{}, err
		}
	}

	defer db.Close()

	return &pb.UserProfile{
		Id:    req.UserProfile.GetId(),
		Email: req.UserProfile.Email,
	}, nil
}

func (s *server) GetUserProfile(ctx context.Context, req *pb.GetUserProfileRequest) (*pb.UserProfile, error) {

	kid := req.Id

	uid, _ := uuid.FromString(kid)

	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DbUser, DbPassword, DbName)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		panic(err)
	}

	rows, err := db.Query(`SELECT * FROM public."user" WHERE id=$1`, uid)
	if err != nil {
		log.Printf("ErrorGet %v", err)
	}
	var id string
	var email string
	var first_name string
	var last_name string
	var birth_date string
	var telephones []string

	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &email, &first_name, &last_name, &birth_date, pq.Array(&telephones))
		if err != nil {
			log.Printf("scan mein error %v", err)
			return &pb.UserProfile{}, err
		}
	}

	//birthDate, _ := time.Parse(time.Stamp, birth_date)

	return &pb.UserProfile{
		Id:         id,
		Email:      email,
		FirstName:  first_name,
		LastName:   last_name,
		Telephones: telephones,
	}, nil
}

func (s *server) ListUsersProfiles(ctx context.Context, req *pb.ListUsersProfilesRequest) (*pb.ListUsersProfilesResponse, error) {
	q := req.GetQuery()

	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DbUser, DbPassword, DbName)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		panic(err)
	}

	rows, err := db.Query(`SELECT * FROM public."user" WHERE first_name LIKE $1`, fmt.Sprintf("%%%s%%", q))
	if err != nil {
		log.Printf("ErrorGet %v", err)
	} else {
		var id string
		var email string
		var first_name string
		var last_name string
		var birth_date string
		var telephones []string
		for rows.Next() {
			err := rows.Scan(&id, &email, &first_name, &last_name, &birth_date, pq.Array(&telephones))
			if err != nil {
				log.Printf("scan mein error %v", err)
				return &pb.ListUsersProfilesResponse{}, err
			} else {
				resp = append(resp, &pb.UserProfile{Id: id, Email: email, FirstName: first_name, LastName: last_name, Telephones: telephones})
			}
		}
	}

	return &pb.ListUsersProfilesResponse{
		Profiles: resp,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// Creates a new gRPC server
	s := grpc.NewServer()
	pb.RegisterUserProfilesServer(s, &server{})
	s.Serve(lis)
}
