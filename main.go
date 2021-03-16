package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/libp2p/go-libp2p"
)

func main() {
	ctx := context.Background()
	node, err := libp2p.New(ctx,
		libp2p.ListenAddrStrings("/ip4/127.0.0.1/tcp/2000"))
	if err != nil {
		panic(err)
	}
	fmt.Println("Listen addresses: ", node.Addrs())

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
	fmt.Println("Recieved signal, closing...")

	if err := node.Close(); err != nil {
		panic(err)
	}
}
