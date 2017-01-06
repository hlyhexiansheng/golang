package main

import (
	"net/rpc/jsonrpc"
	"log"
	"net"
	"net/rpc"
	"fmt"
)

func main() {

	fmt.Printf("server start.....")

	StartRpc()
}

type Transfer int

func (t *Transfer) Update2(args *string, reply *string) error {
	fmt.Println("got it" + *args)
	*reply = "4444444";
	return "fdas"
}

func StartRpc() {

	addr := "127.0.0.1:8888";
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		log.Fatalf("net.ResolveTCPAddr fail: %s", err)
	}

	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Fatalf("listen %s fail: %s", addr, err)
	} else {
		log.Println("rpc listening", addr)
	}

	server := rpc.NewServer()
	server.Register(new(Transfer))

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("listener.Accept occur error:", err)
			continue
		}
		go server.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
