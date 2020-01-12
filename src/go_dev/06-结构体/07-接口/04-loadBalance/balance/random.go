package balance

import (
	"errors"
	"math/rand"
)

//
/*

 */

type RandomBalance struct {
}

func init() {
	RsgisterB("random", &RandomBalance{})
}

func (randomBalance *RandomBalance) DoBalance(insts []Instance, key ...string) (inst Instance, err error) {
	if len(insts) == 0 {
		err = errors.New("no intsance")
		return
	}
	lengh := len(insts)
	index := rand.Intn(lengh)
	inst = insts[index]
	return

}
