package rpc

import (
	"net"
	"reflect"
)

// Client 是 RPC 的客户端
type Client struct {
	conn net.Conn
}

func NewClient(conn net.Conn) *Client {
	return &Client{conn: conn}
}

// Call 通过 RPC 调用服务端的函数
// name 为 RPC 服务端注册的函数名
// funcPtr 为函数原型
// 该函数运行的结果是将一个封装 RPC 调用相应函数的函数赋给 funcPtr
// e.g. 在服务端实现，并注册了函数
//     func queryUser(uid int) (User, error)
// 在使用客户端时通过原型调用 RPC 函数：
//	  var query func(int) (User, error)
//	  client.Call("queryUser", &query)
//	  u, err := query(1)
func (c *Client) Call(name string, funcPtr interface{}) {
	// 反射初始化 funcPtr 函数原型
	fn := reflect.ValueOf(funcPtr).Elem()
	f := func(args []reflect.Value) []reflect.Value {
		// 参数
		inArgs := make([]interface{}, 0, len(args))
		for _, arg := range args {
			inArgs = append(inArgs, arg.Interface())
		}
		// 连接 RPC 服务
		cliSession := NewSession(c.conn)

		// 请求
		requestRPCData := RPCData{
			Func: name,
			Args: inArgs,
		}
		requestEncoded, err := encode(requestRPCData)
		if err != nil {
			panic(err)
		}

		if err := cliSession.Write(requestEncoded); err != nil {
			panic(err)
		}

		// 响应
		response, err := cliSession.Read()
		if err != nil {
			panic(err)
		}

		respRPCData, err := decode(response)
		if err != nil {
			panic(err)
		}
		outArgs := make([]reflect.Value, 0, len(respRPCData.Args))
		for i, arg := range respRPCData.Args {
			if arg == nil {
				outArgs = append(outArgs, reflect.Zero(fn.Type().Out(i)))
			} else {
				outArgs = append(outArgs, reflect.ValueOf(arg))
			}
		}
		return outArgs
	}

	// 将 RPC 调用函数赋给 fn
	v := reflect.MakeFunc(fn.Type(), f)
	fn.Set(v)
}
