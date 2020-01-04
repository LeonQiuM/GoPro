package main

import (
	"fmt"
	"go_dev/06-结构体/07-接口/04-loadBalance/balance"
	"hash/crc32"
	"math/rand"
)

//
/*
自己注册接口实现一致性 hash 负载均衡算法

*/

type HashBalance struct {
}

func init() {
	balance.RsgisterB("hash", &HashBalance{})
}

func (hashBalance *HashBalance) DoBalance(insts []balance.Instance, key ...string) (inst balance.Instance, err error) {
	var defaultKey string = fmt.Sprintf("%d", rand.Int())
	if len(key) > 0 {
		//err = fmt.Errorf("Hash balance must pass hash key")
		defaultKey = key[0]
		return
	}
	lengh := len(insts)
	if lengh == 0 {
		err = fmt.Errorf("not fount backend instance")
		return
	}
	crcTable := crc32.MakeTable(crc32.IEEE)
	hashValue := crc32.Checksum([]byte(defaultKey), crcTable)
	index := int(hashValue) % lengh
	inst = insts[index]
	return
}
