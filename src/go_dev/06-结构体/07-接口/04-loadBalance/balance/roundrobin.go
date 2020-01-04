package balance

import (
	"errors"
)

//
/*

 */

type RRBalance struct {
	currentIndex int
}

func init() {
	RsgisterB("roundrobin", &RRBalance{})
}

func (rrBalance *RRBalance) DoBalance(insts []Instance, key ...string) (inst Instance, err error) {
	if len(insts) == 0 {
		err = errors.New("no intsance")
		return
	}
	if rrBalance.currentIndex >= len(insts) {
		// 越界
		rrBalance.currentIndex = 0
	}
	inst = insts[rrBalance.currentIndex]
	rrBalance.currentIndex++
	return

}
