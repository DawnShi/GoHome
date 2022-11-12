package calc

import (
	"net/rpc"
)

// ServiceName 计算器服务名
const ServiceName = "CalcService"

// ServiceInterface 计算器服务接口
type ServiceInterface interface {
	// CalaTwoNumber 对两个数进行运算
	CalcTwoNumber(request Calc, reply *float64) error
	// GetOperators 获取所有支持的运算
	GetOperators(request struct{}, repchan *[]string) error
}

// RegisterCalcService register the RPC service on svc
func RegisterCalcService(svc ServiceInterface) error {
	return rpc.RegisterName(ServiceName, svc)
}

// Cala 定义计算对象, 包括两个运算数
type Calc struct {
	Number1  float64
	Number2  float64
	Operator string
}
