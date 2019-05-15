package heartbeat

import (
	"context"
	"time"
)

type Config struct {
	// Context is the default client context; it can be used to cancel grpc dial out and
	// other operations that do not have an explicit context.
	Context context.Context

	Endpoint string

	HeartbeatInterval time.Duration

	HeartbeatTimeout time.Duration

	AppName string
}
