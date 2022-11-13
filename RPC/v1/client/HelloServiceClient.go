package main

import (
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// HelloServiceClient is a client for HelloService
type HelloServiceClient struct {
	*rpc.Client
}

// HelloServiceName is the name of HelloService
const HelloServiceName = "HelloService"

// HelloServiceInterface is a interface for HelloService
type HelloServiceInterface interface {
	Hello(request string, replpy *string) error
}

var _ HelloServiceInterface = (*HelloServiceClient)(nil)

// DialHelloService dial HelloService
func DialHelloService(network, address string) (*HelloServiceClient, error) {
	conn, _ := net.Dial(network, address)
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	return &HelloServiceClient{Client: client}, nil
}

// Hello calls HelloService.Hello
func (p *HelloServiceClient) Hello(request string, reply *string) error {
	return p.Client.Call(HelloServiceName+".Hello", request, reply)
}
