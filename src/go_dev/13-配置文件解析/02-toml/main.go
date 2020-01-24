package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"path/filepath"
	"sync"
	"time"
)

//
/*
go get github.com/BurntSushi/toml/cmd/tomlv
https://studygolang.com/articles/12032
go get github.com/BurntSushi/toml

*/

//type Config struct {
//	Person   Person
//	Mysql Mysql
//}
//
//type Person struct {
//	Age        int
//	Cats       []string
//	Pi         float64
//	Perfection []int
//	DOB        time.Time
//}
//
//type Mysql struct {
//	Database string
//	Host     string
//	Port     int
//	Password string
//	User     string
//}

type tomlConfig struct {
	Title   string
	Owner   ownerInfo
	DB      database `toml:"database"`
	Servers map[string]server
	Clients clients
}

type ownerInfo struct {
	Name string
	Org  string `toml:"organization"`
	Bio  string
	DOB  time.Time
}

type database struct {
	Server  string
	Ports   []int
	ConnMax int `toml:"connection_max"`
	Enabled bool
}

type server struct {
	IP string
	DC string
}

type clients struct {
	Data  [][]interface{}
	Hosts []string
}

var (
	once sync.Once
)

type EnvConfig struct {
	Title string
	LF    envConf
	HL    envConf
	LQ    envConf
}

type envConf struct {
	Mysql       string `toml:"mysql"`
	Token       string `toml:"token"`
	EtcdPath    string `toml:"etcd_path"`
	Conn        int    `toml:"conn"`
	Debug       bool   `toml:"debug"`
	RedisCenter string `toml:"redis_center"`
}

//func ConfigEncode() *tomlConfig {
//	// 单例模式，配置文件只读取一次
//	var cfg tomlConfig
//	once.Do(func() {
//		filePath, err := filepath.Abs("/Users/qiumeng/GOSPACE/GoPro/src/go_dev/13-配置文件解析/02-优雅的等待goroutine退出-toml/config.toml")
//		if err != nil {
//			panic(err)
//		}
//		if _, err := toml.DecodeFile(filePath, &cfg); err != nil {
//			panic(err)
//		}
//	})
//	return &cfg
//}

func ConfigEnv() *EnvConfig {
	// 单例模式，配置文件只读取一次
	var cfg EnvConfig
	once.Do(func() {
		filePath, err := filepath.Abs("/Users/qiumeng/GOSPACE/GoPro/src/go_dev/13-配置文件解析/02-优雅的等待goroutine退出-toml/env.toml")
		if err != nil {
			panic(err)
		}
		if _, err := toml.DecodeFile(filePath, &cfg); err != nil {
			panic(err)
		}
	})
	return &cfg
}

func main() {
	//con := ConfigEncode()
	//fmt.Printf("%s\n",con.Owner.Name)
	//fmt.Printf("%s\n",con.DB.Server)
	//fmt.Println(con.Servers["alpha"])
	//var config tomlConfig
	//filePath := "/Users/qiumeng/GOSPACE/GoPro/src/go_dev/13-配置文件解析/02-优雅的等待goroutine退出-toml/config.toml"
	//if _, err := toml.DecodeFile(filePath, &config); err != nil {
	//	panic(err)
	//}
	//fmt.Println(config.Owner.DOB)
	cenv := ConfigEnv()
	fmt.Printf("cenv.ENVlf.Mysql:%v\n", cenv.LF.Mysql)
	fmt.Printf("cenv.ENVlf.Mysql:%v\n", cenv.LF.RedisCenter)
	fmt.Printf("cenv.ENVlf.Mysql:%v\n", cenv.HL.Mysql)
	fmt.Printf("cenv.ENVlf.Mysql:%v\n", cenv.HL.RedisCenter)
	fmt.Printf("cenv.ENVlf.Mysql:%v\n", cenv.LQ.Mysql)
	fmt.Printf("cenv.ENVlf.Mysql:%v\n", cenv.LQ.RedisCenter)

	fmt.Println(cenv.Title)
}
