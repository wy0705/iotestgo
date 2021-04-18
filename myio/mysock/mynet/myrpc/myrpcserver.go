package myrpc

import (
	"fmt"
	"net"
	"net/http"
	"net/rpc"
)

type Arg int;

func (self *Arg)Incr(arg int,replyres *int)error{
	fmt.Println("收到的内容是:",arg)
	*replyres=arg+1
	return nil
}
func MyRPCServer(){
	arg:=new(Arg)
	rpc.Register(arg)
	rpc.HandleHTTP()
	ln,_:=net.Listen("tcp","0.0.0.0:9999")
	http.Serve(ln,nil)
}
