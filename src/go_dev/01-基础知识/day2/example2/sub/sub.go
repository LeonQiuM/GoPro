package sub

import "fmt"

var Name string = "subxxxx"
var Age int = 199

func init() {
	fmt.Printf("this si a test sub\n")
	fmt.Printf("sub pack name:%s, addr:%p\n", Name, &Name)
	fmt.Printf("sub pack age:%d, addr:%p\n", Age, &Age)
}
