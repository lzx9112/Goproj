package main

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "proto"
	"time"
)

func main() {
	var opts []grpc.DialOption
	ctx := context.Background()
	opts = append(opts, grpc.WithInsecure())
	destAddr := "192.168.0.101:50003"
restart:
	conn, err := grpc.Dial(destAddr, opts...)
	if err == nil {
		defer conn.Close()
	} else {
		time.Sleep(time.Second * 3)
		fmt.Println("reconnect", err)
		goto restart
	}

	client := pb.NewHeartBeatClient(conn)
	stream, err := client.Publish(ctx)
	if err != nil {
		time.Sleep(time.Second * 3)
		fmt.Println("Publish failed reconnect", err)
		goto restart
	}

	for {
		err := stream.Send(&pb.SendMsg{
			Timestamp: time.Now().Unix(),
			Msg:       "Monitor",
		})
		if err != nil {
			fmt.Println("Send failed reconnect", err)
			goto restart
		}
		time.Sleep(time.Second * 3)
	}
}
