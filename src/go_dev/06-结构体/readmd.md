## 结构体
1. 用来定于复杂的数据结构
2. struct 里面可以包含多个字段
3. struct 可以定义方法，注意和函数的区分
4. 结构体是值类型
5. 结构体可以前台
### 声明

```
type 标识符 struct {
    field1 type
    field2 type
}
举例：
type Student struct {
    Name string
    Age int
    Score int
}
```
### 定义
```
1. var stu Student
2. var stu *Student = new (Student)
3. var stu *Student = &Student{}

其中 2、3返回的是指向结构体的指针
访问:
stu.Name
(*stu).Name
```


### 初始化与赋值
所有字段在内存中是连续的
```
var std1 Student
type std Student // 别名
std1.Name = "James"
std1.Age = 33
std1.Score = 20

fmt.Printf(
    "Name=%s,age=%d,score=%d",std1.Name.std1.Age,std1.Score
)

var stu1 *Student = &Student {
    Age:10,
    Name:"James"
    // 可以只初始化一部分值属性
}

var stu2 = Student {
    Name:"Jack",
    Score:123
    Age:11,
}

```

## 链表

```
type Student struct {
    Name string
    Next * Student
}

```