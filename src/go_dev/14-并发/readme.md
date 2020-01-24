## goroutine 与线程
   goroutine是用户态的线程 
   goroutine的栈不是固定的，一把初识2KB，最大为1GB，他可以按需要增加或者减少
   
## goroutine 的调度
GMP调度
G=goroutine
M=是runtime对操作系统内核线程的虚拟，与其一一对应
P=管理者一组goroutine队列，P里面会存储当前goroutine 的上下文环境

P的个数是通过runtime.GOMAXPROCS谁定，最大256


## 其他

一个操作系统线程对应多个用户态的goroutine
go的程序可以同时使用多个操作系统线程
goroutine和os线程是多对多的关系，即m:n