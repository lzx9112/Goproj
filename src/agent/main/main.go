package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	pb "proto"
	"time"
)

func main() {
	lis, err := net.Listen("tcp", ":50003")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterHeartBeatServer(grpcServer, newServer())
	grpcServer.Serve(lis)

}

type Server struct{}

func newServer() pb.HeartBeatServer {
	return &Server{}
}

func (s *Server) Publish(stream pb.HeartBeat_PublishServer) error {
	fmt.Println("aaaaaaaaa")
	for {
		msg, err := stream.Recv()
		if err == nil {
			fmt.Println(msg)
			stream.Send(&pb.ReceiveMsg{
				Timestamp: time.Now().Unix(),
				Msg:       "agent",
			})
		}
		time.Sleep(time.Second * 3)

	}
}
