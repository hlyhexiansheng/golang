package main

import (
	"fmt"
	"net/rpc/jsonrpc"
	"net"
	"net/rpc"
)

func main()  {
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	var rpcClient *rpc.Client = jsonrpc.NewClient(conn)
	var request string = "2222222";
	var s string;
	var err2 = rpcClient.Call("Transfer.Update2",&request ,&s)

	if err2 != nil{
		fmt.Println("error" + err2.Error())
	}
	fmt.Println(s)
}