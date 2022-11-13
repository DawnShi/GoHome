package rpc

import (
	"log"
	"net"
	"reflect"
)

// Server 是 RPC 服务端
type Server struct {
	funcs map[string]reflect.Value
}

func NewServer() *Server {
	return &Server{funcs: map[string]reflect.Value{}}
}

// Register 注册绑定要 RPC 服务的函数
// 将函数名与函数对应起来
func (s *Server) Register(name string, function interface{}) {
	// 已存在则不处理
	if _, ok := s.funcs[name]; ok {
		return
	}
	fVal := reflect.ValueOf(function)
	s.funcs[name] = fVal
}

// ListenAndServe 监听 address，运行 RPC 服务
func (s *Server) ListenAndServe(address string) error {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("accept error:", err)
			continue
		}
		s.handleConn(conn)
	}
}

// handleConn 处理 RPC 服务的 conn 请求
// 创建 RPC 会话，解码请求体，调用本地函数完成工作，编码响应，返回结果
func (s *Server) handleConn(conn net.Conn) {
	// 创建会话
	srvSession := NewSession(conn)

	// 读取、解码数据
	data, err := srvSession.Read()
	if err != nil {
		log.Println("session read error:", err)
		return
	}
	requestRPCData, err := decode(data)
	if err != nil {
		log.Println("data decode error:", err)
		return
	}

	// 获取函数
	f, ok := s.funcs[requestRPCData.Func]
	if !ok {
		log.Printf("unexpected rpc call: function %s not exist", requestRPCData.Func)
		return
	}

	// 获取参数
	inArgs := make([]reflect.Value, 0, len(requestRPCData.Args))
	for _, arg := range requestRPCData.Args {
		inArgs = append(inArgs, reflect.ValueOf(arg))
	}

	// 反射调用方法
	returnValues := f.Call(inArgs)

	// 构造结果
	outArgs := make([]interface{}, 0, len(returnValues))
	for _, ret := range returnValues {
		outArgs = append(outArgs, ret.Interface())
	}
	replyRPCData := RPCData{
		Func: requestRPCData.Func,
		Args: outArgs,
	}
	replyEncoded, err := encode(replyRPCData)
	if err != nil {
		log.Println("reply encode error:", err)
		return
	}

	// 写入结果
	err = srvSession.Write(replyEncoded)
	if err != nil {
		log.Println("reply write error:", err)
	}
}
