package main

import (
	"context"
	"log"
	"math/rand"
	"net"

	pb "github.com/Nadeem1815/microservice-grpc/usermgmt/proto"
	"google.golang.org/grpc"
)

// grpc going run on this port
const (
	port = ":50051"
)

type UserManagementServer struct {
	pb.UnimplementedUserManagementServer
}

func (s *UserManagementServer) CreateNewUser(ctx context.Context, in *pb.NewUser) (*pb.User, error) {
<<<<<<< HEAD
	log.Printf("Received: %v", in.GetName())
	var user_id int32 = int32(rand.Intn(1000))
	return &pb.User{Name: in.GetName(), Age: in.GetAge(), Id: user_id}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen :%v", err)
	}

	// involk the New Server functon from the grpc modul
	s := grpc.NewServer()
	pb.RegisterUserManagementServer(s, &UserManagementServer{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)

	}
}
=======
	log.Printf("Received: %v",in.GetName())
	var user_id int32 =int32(rand.Intn(1000))
	return &pb.User{Name: in.GetName(),Age: in.GetAge(),Id: user_id},nil
}


func main() {
	lis,err:= net.Listen("tcp",port)
	if err!=nil {
	 	log.Fatalf("failed to listen :%v",err)		
	}

	// involk the New Server functon from the grpc modul 
	s:=grpc.NewServer()
	pb.RegisterUserManagementServer(s,&UserManagementServer{})
	log.Printf("server listening at %v",lis.Addr())
	if err :=s.Serve(lis); err!=nil {
		log.Fatalf("failed to serve: %v",err)
		
	}
}
>>>>>>> 10720efab6bed8ab1762739dbc930044dc69f3ad
