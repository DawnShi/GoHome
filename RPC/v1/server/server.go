package main

import (
	"log"
	"net"
	"net/rpc/jsonrpc"
)

// HelloService is a RPC service for helloWorld
type HelloService struct{}

// Hello say hello to request
func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "Hello, " + request
	return nil
}

func main() {
	// 用将给客户端访问的名字和HelloService实例注册 RPC 服务
	// rpc.RegisterName("HelloService", new(HelloService))
	RegisterHelloService(new(HelloService))

	// HTTP 服务
	// rpc.HandleHTTP()
	// err := http.ListenAndServe(":1234", nil)
	// if err != nil {
	// 	log.Fatal("Http Listen and Server: ", err)
	// }

	// TCP 服务
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error: ", err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error: ", err)
		}
		// rpc.ServeConn(conn)
		// 使用 json RPC
		go jsonrpc.ServeConn(conn)
	}
}
