package myrpc

import (
	"fmt"
	"net/rpc"
)

func MyRPCCli() {
	cli,_:=rpc.DialHTTP("tcp","localhost:9999")
	var arg int=11
	cli.Call("Arg.Incr",arg,&arg)
	fmt.Println(arg)
}
