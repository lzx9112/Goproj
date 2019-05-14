package heartbeat

import (
	"context"
	"fmt"
	pb "heartbeat/proto"
	"monitor/scheduler"
	"time"
)

type HearbeatClient struct {
	addr       string
	duration   time.Duration
	retryCount uint64
}



func newHearbeatClient() *HearbeatClient {
	return &HearbeatClient{
		duration:   time.Second * 3,
		retryCount: 5,
	}
}

func NewInstance(addr string) *HearbeatClient {
	hb := newHearbeatClient()
	hb.addr = addr
	return hb
}

func (hb *HearbeatClient) Start(ctx context.Context) {
	var agentFailedCount uint64
	ts := scheduler.NewScheduler()
	ts.SetDuration(time.Second)
	go ts.Run()
	defer ts.Stop()
restart:
	_, err := hb.tunnel.Publish(ctx)
	if err != nil {
		time.Sleep(hb.Duration())
		fmt.Println("Publish failed reconnect", err)
		goto restart
	}
	c := time.Tick(hb.Duration())
	for {
		select {
		case <-c:
			sendTask := ts.CreateTask(func() {
				err := hb.tunnel.Send(&pb.SendMsg{
					Timestamp: time.Now().Unix(),
					Msg:       "Monitor",
				})
				if err != nil {
					fmt.Println("Send failed", err)
				}
			})
			ts.AddTask(sendTask)

			rcvTask := ts.CreateTask(func() {
				rcv, err := hb.tunnel.Recv()
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

			if agentFailedCount >= hb.RetryCount() {
				agentFailedCount = 0
				goto restart
			}
			break
		case <-ctx.Done():
			fmt.Println("exit heartbeat")
			return
		}
	}
}
