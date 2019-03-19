package main

import (
	"golang.org/x/net/context"
	. "monitor/heartbeat"
	"time"
)

func main() {
	ctx := context.Background()
	destAddr := "192.168.0.102:50003"
	hb := NewInstance(destAddr)
	hb.SetDuration(time.Second*3)
	hb.Start(ctx)
}
