package main

import (
	"flag"
	"fmt"
	"golang.org/x/net/context"
	. "monitor/heartbeat"
	"os"
	"os/signal"
	"time"
)

var monitorPort = flag.String("p", "50003", "port")

func main() {
	flag.Parse()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	ctx, cancelFunc := context.WithCancel(context.Background())
	destAddr := "localhost:" + *monitorPort
	hb := NewInstance(destAddr)
	hb.SetDuration(time.Second*3)
	go hb.Start(ctx)
	<-c
	cancelFunc()
	time.Sleep(time.Second)
	fmt.Println("exit")
}
