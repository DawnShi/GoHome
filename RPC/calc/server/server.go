package main

import (
	"dawnshi/RPC/calc"
	"net/http"
	"net/rpc"
)

/* RPC 服务实现 */

// CalcService 是计算器 RPC 服务的实现
type CalcService struct{}

// CalcTwoNumber 对两个数进行加减乘除操作
// 常见错误： 单词错误，导致client 调用方法名不一致
func (c *CalcService) CalcTwoNumber(request calc.Calc, reply *float64) error {
	oper, err := CreateOperation(request.Operator)
	if err != nil {
		return err
	}
	*reply = oper(request.Number1, request.Number2)
	return nil
}

// GetOperators 获取所有支持的运算
func (c *CalcService) GetOperators(request struct{}, reply *[]string) error {
	opers := make([]string, 0, len(Operators))
	for key := range Operators {
		opers = append(opers, key)
	}
	*reply = opers
	return nil
}

/* 运行 RPC 服务*/

func main() {
	calc.RegisterCalcService(new(CalcService))
	rpc.HandleHTTP()
	http.ListenAndServe(":8899", nil)
}
