## goroutine 与线程

```
goroutine是用户态的线程，比内核OS的线程(2M)更加轻量级
goroutine的栈不是固定的，一把初识2KB，最大为1GB，他可以按需要增加或者减少
```

1. 使用`go`关键字进行调用
2. 将要并发执行的函数包装成一个函数
3. 函数执行完成或者return，这个goroutine结束
4. 程序主函数启动的时候会自动的创建一个goroutine来执行main函数
5. main函数结束了，那么其他所有的开启的goroutine都自动结束了

## sync.WaiteGroup
```go
// 用于优雅的等待goroutine结束

var wg sync.WaiteGroup
wg.Add(1) // 登记准入
wg.Done() // 登记准出
wg.Waie() // 等待所有goroutine结束
```  

## goroutine 的调度模型
```text
GMP调度
G=goroutine
M=是runtime对操作系统内核线程的虚拟，与其一一对应
P=管理者一组goroutine队列，P里面会存储当前goroutine 的上下文环境
P的个数是通过runtime.GOMAXPROCS谁定，最大256，1.5版本后，默认为默认的逻辑核心数，默认跑满CPU
runtime.GOMAXPROCS(1) // 设定只占用一个物理核
```


## 其他
一个操作系统线程对应多个用户态的goroutine
go的程序可以同时使用多个操作系统线程
goroutine和os线程是多对多的关系，即m:n

# channel
1. channel存在的意义：goroutine只能让某一段任务执行，而无法做任务之前的协调
```text
CSP 进程间通信的基本方式之一，通过通信来共享内存
```
2. channel 是一种引用类型，需要make初始化，其他的需要make初始化的引用类型还有slice、map
3. channel的声明与初始化
```text
// 声明
var ch chan 类型
var ch chan int
// 初始化
 ch = make(chan 元素类型, [缓冲区大小])
```
4. channel的三种操作
```text
// 发送:   ch <-
// 接受:  <- ch
// 关闭: close(ch)   
// 由于通道是一种类型，就是一个变量，会被垃圾回收，所以不一定需要关闭
```
5. 带缓冲区和无缓冲区的通道
```text
1. 无快递柜，快递员投送快递
2. 有快递柜，快递员投送快递，快递柜满了，投递也会阻塞
3. 当channel中为空，依然从中取值，等待取值也会阻塞，所以会报死锁

```