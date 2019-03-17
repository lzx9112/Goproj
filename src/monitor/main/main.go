package main

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"os/exec"
	pb "proto"
	"time"
)

func main() {
	var opts []grpc.DialOption
	ctx := context.Background()
	opts = append(opts, grpc.WithInsecure())
	destAddr := "192.168.0.101:50003"
	ts := NewTaskScheduler()
	go ts.run()
	defer ts.Stop()
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

	agentFailedCount := 0
	for {

		sendTask := ts.CreateTask(func() {
			err := stream.Send(&pb.SendMsg{
				Timestamp: time.Now().Unix(),
				Msg:       "Monitor",
			})
			if err != nil {
				fmt.Println("Send failed", err)
			}
		})
		ts.AddTask(sendTask)

		rcvTask := ts.CreateTask(func() {
			rcv, err := stream.Recv()
			if err == nil {
				fmt.Println(rcv)
				if rcv == nil {
					agentFailedCount++
				}
			} else {
				agentFailedCount++
				fmt.Println("Receive failed", err)
			}
		})
		ts.AddTask(rcvTask)

		if agentFailedCount >= 3 {
			fmt.Println("restart agent")
			agentFailedCount = 0
			cmd := exec.Command("bash", "-c", "/Users/lizhaoxing/Downloads/0-proj/GoProj/src/agent/main/main")
			cmd.Start()
			goto restart

		}

		time.Sleep(time.Second * 3)
	}
}
