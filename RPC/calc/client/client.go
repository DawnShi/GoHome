package main

import (
	"dawnshi/RPC/calc"
	"fmt"
	"log"
	"net/rpc"
)

/* 定义客户端实现 */

// CalcClient is a client for CalcService
type CalcClient struct {
	*rpc.Client
}

// var _ calc.ServiceInterface = (*CalcClient)(nil)

// DialCalcService dial CalcService
func DialCalcService(network, address string) (*CalcClient, error) {
	c, err := rpc.DialHTTP(network, address)
	if err != nil {
		return nil, err
	}
	return &CalcClient{Client: c}, nil
}

// CalcTwoNumber 对两个数进行运算
func (c *CalcClient) CalcTwoNumber(request calc.Calc, reply *float64) error {
	return c.Client.Call(calc.ServiceName+".CalcTwoNumber", request, reply)
}

// GetOperators 获取所有支持的运算
func (c *CalcClient) GetOperators(request struct{}, reply *[]string) error {
	return c.Client.Call(calc.ServiceName+".GetOperators", request, reply)
}

/* 使用客户端调用 RPC 服务 */

func main() {
	client, err := DialCalcService("tcp", "localhost:8899")
	if err != nil {
		log.Fatal("Err Dial Client:", err)
	}
	// Test GetOperators
	var opers []string
	err = client.GetOperators(struct{}{}, &opers)
	if err != nil {
		log.Println(err)
	}
	log.Println("opers", opers)

	// Test CalcTwoNumber
	testAdd := calc.Calc{
		Number1:  3.14,
		Number2:  8.98,
		Operator: "+",
	}
	var result float64
	err = client.CalcTwoNumber(testAdd, &result)
	if err != nil {
		fmt.Println("err", err)
	}
	log.Println("result:", result)

}
