## 接口
1. Interface 类型可以定义一组方法，但是这些不需要实现，并且 interface 不能包含任何变量

### 实现
1. 接口不需要显示的实现。只需要一个变量，含有接口类型中的所有方法，那么这个变量就实现了这个接口
2. 如果一个变量含有多个 interface 的方法，那么这个变量就实现了多个接口

### 接口的嵌套
```
type RW interface {
    Read(b buffer) bool
    Write(b buffer) bool
}

type Lock interface {
    Lock()
    UnLock()
}

type File interface {
    RW
    Look
    Close()
}
```

### 接口的断言

```
var t int
var x interface{}
x = t
if t==x.(int)  // 转为 int
if y,ok = x.(int) // 待检查的判断
```

### 空接口
空接口没有任何方法，所以任何类型都实现了空接口

### 判断一个变量是否实现执行接口

```$xslt

```