package balance

import "fmt"

//
/*
用于管理所有的负载均衡算法
*/

type BalanceManager struct {
	allBalancer map[string]Balance
}

var mgr = BalanceManager{
	allBalancer: make(map[string]Balance),
}

func (balanceManger *BalanceManager) registerBalance(name string, b Balance) {
	balanceManger.allBalancer[name] = b

}

func RsgisterB(name string, b Balance) {
	mgr.registerBalance(name, b)
}

func DoBalance(name string, insts []Instance) (inst Instance, err error) {
	balancer, ok := mgr.allBalancer[name]
	if !ok {
		err = fmt.Errorf("load balance method not fount:%s", name)
		return
	}
	inst, err = DoBalance(insts)
	return
}
