package _1_ini

//
/*
ini 文件解析
*/
type RedisConfig struct {
	Database string `ini:database`
	Host     string `ini:host`
	Port     int    `ini:port`
	Password string `ini:password`
	User     string `ini:user`
}

func main() {

}

func loadIni(filename string, data interface{}) {
	// 0. 判断传入的 data 是否为指针类型，因为需要赋值
	// 1. 读取字节类型数组
	//
	// 2. 一行一行的读取
	// 2.1 如果是注释，跳过
	// 2.2 如果是[ 开头，那么说明他是一个节点 section  将中括号中间的值复制给结构体
	// 2.3 如果不是[] 开头 就是以等号分割的键值对
}
