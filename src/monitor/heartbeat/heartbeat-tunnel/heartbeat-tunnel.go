package heartbeat_tunnel

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	. "proto"
)

type HeartbeatTunnel interface {
	Publish(ctx context.Context, opts ...grpc.CallOption) (HeartBeat_PublishClient, error)
	Send(*SendMsg) error
	Recv() (*ReceiveMsg, error)
}

type heartbeatTunnel struct {
	client HeartBeatClient
	stream HeartBeat_PublishClient
}

func newHeartbeatTunnel() *heartbeatTunnel {
	return &heartbeatTunnel{}
}

func (hb *heartbeatTunnel) Publish(ctx context.Context, opts ...grpc.CallOption) (HeartBeat_PublishClient, error) {
	stream, err := hb.client.Publish(ctx, opts...)
	if hb != nil {
		hb.stream = stream
	}
	return stream, err
}

func (hb *heartbeatTunnel) Send(m *SendMsg) error {
	if hb == nil || hb.stream == nil {
		return fmt.Errorf("no valid stream")
	}
	return hb.stream.Send(m)
}

func (hb *heartbeatTunnel) Recv() (*ReceiveMsg, error) {
	if hb == nil || hb.stream == nil {
		return nil, fmt.Errorf("no valid stream")
	}
	return hb.stream.Recv()
}

func NewInstance(addr string) HeartbeatTunnel{
	var opts []grpc.DialOption
	t := newHeartbeatTunnel()
	if addr == "" {
		return nil
	}
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		return nil
	}
	t.client = NewHeartBeatClient(conn)
	return t
}
