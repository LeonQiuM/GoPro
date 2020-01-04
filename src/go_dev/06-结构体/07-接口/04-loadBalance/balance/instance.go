package balance

import "strconv"

//
/*

 */
type Instance struct {
	Host string
	Port int
}

func (inst *Instance) String() string {
	return inst.Host + ":" + strconv.Itoa(inst.Port)
}
