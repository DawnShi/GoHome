package main

import (
	"fmt"
	"log"
)

func main() {
	// HTTP
	// client, err := rpc.DialHTTP("tcp", "localhost:1234")
	// json RPC
	// client, err := jsonrpc.Dial("tcp", "localhost:1234")

	client, err := DialHelloService("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing: ", err)
	}

	var reply string
	err = client.Hello("world", &reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply)

}

// func main() {
// 	conn, _ := net.Dial("tcp", "localhost:1234")
// 	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
// 	var reply string
// 	err := client.Call("HelloService.Hello", "world", &reply)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(reply)
// }
