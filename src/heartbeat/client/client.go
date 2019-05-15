package heartbeat

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "heartbeat/proto"
	"sync"
	"time"
)

var DefaultRequestTimeout = 5 * time.Second

type HbClient struct {
	handler    pb.HeartbeatClient

	cfg    Config

	// DialOptions is a list of dial options for the grpc client (e.g., for interceptors).
	dialOptions []grpc.DialOption

	conn *grpc.ClientConn

	mu   *sync.RWMutex
	ctx    context.Context
	cancel context.CancelFunc
}

func newClient(baseCtx context.Context, conn *grpc.ClientConn, cfg *Config) (*HbClient, error) {
	if baseCtx == nil {
		return nil, fmt.Errorf("contex is nil")
	}

	if cfg == nil {
		cfg = &Config{}
	}

	ctx, cancel := context.WithCancel(baseCtx)
	client := &HbClient{
		dialOptions: []grpc.DialOption{},
		cfg:      *cfg,
		ctx:      ctx,
		cancel:   cancel,
		mu:       new(sync.RWMutex),
	}

	if conn != nil {
		client.conn = conn
		client.handler = pb.NewHeartbeatClient(client.conn)
		return client, nil
	}

	conn, err := client.dial(client.ctx)
	if err != nil {
		client.cancel()
		return nil, err
	}

	client.handler = pb.NewHeartbeatClient(client.conn)


	return client, nil
}

func (hb *HbClient) dial(ctx context.Context) (*grpc.ClientConn, error) {
	if hb.cfg.Endpoint == "" {
		return nil, fmt.Errorf("dest address is empty")
	}

	opts := append(hb.dialOptions, grpc.WithInsecure())

	return grpc.DialContext(ctx,hb.cfg.Endpoint,opts...)
}

func (hb *HbClient) Run(baseCtx context.Context) {
	if hb.cfg.HeartbeatInterval == 0 {
		return
	}

	ctx, cancel := context.WithTimeout(baseCtx, hb.cfg.HeartbeatInterval)
	ticker := time.NewTicker(hb.cfg.HeartbeatInterval)
	for {
		select {
		case <-ticker.C:

			_, err := hb.handler.Send(
				ctx,
				&pb.HeartbeatReq{
					Timestamp:time.Now().Unix(),
					AppName:hb.cfg.AppName,
					Msg:"heartbeat from "+hb.cfg.AppName,
				},
			)

			if err != nil {
				// write to checker
			}

		case <-ctx.Done():
			cancel()
			fmt.Println("exit heartbeat")
			return
		}
	}
}
