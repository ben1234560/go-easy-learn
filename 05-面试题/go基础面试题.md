# go基础面试题

> 推荐https://blog.csdn.net/itcastcpp/article/details/80462619

### 选择题

1. 关于类型转换，下面语法正确的是（）单选

~~~go
A.
type MyInt i
int var i int = 1
var j MyInt = I
B.
type MyInt int
var i int = 1
var j MyInt = (MyInt)i string(x)
C.
type MyInt int
var i int = 1
var j MyInt = MyInt(i)
D.
type MyInt int
var i int = 1
var j MyInt = i.(MyInt)
~~~



2. 阅读一下代码，写出输出结果（）单选

~~~go
func main(){
    var a int = 10
    fmt.Println(a)
    {
        a := 9
        fmt.Println(a)
        a = 8
    }
    fmt.Println(a)
}
A.10 10 10
B.10 9 9
C.10 9 10
D.10 9 8
~~~

3. 下面程序的运行结果是（）单选

~~~go
func main(){
    if (true){
        defer fmt.Printf("1")
    } else {
        defer fmt.Printf("2")
    }
    fmt.Print("3")
}
A.321
B.32
C.31
D.13
~~~

4. 关于函数声明，下面语法错误的是（）单选

~~~go
A.func f(a,b int) (value int, err error)
B.func f(a int,b int) (value int, err error)
C.func f(a,b int) (value int, error)
D.func f(a int,b int) (int, int, error)
~~~

5. 关于无缓冲和有缓冲的channel，下面说法正确的是（）单选

~~~go
A.无缓冲的channel是默认的缓冲为1的channel
B.无缓冲的channel和有缓冲的channel都是同步的
C.无缓冲的channel和有缓冲的channel都是非同步的
D.无缓冲的channel是同步的，有缓存的channel是非同步的
~~~

6. 以下代码，最终输出的结果是什么（）单选

~~~go
type student struct{
    Name string
    Age int
}

func main(){
    m := make(map[string]*student)
    stus := []student{
        {Name: "zhou", Age:24},
        {Name: "li", Age:23},
        {Name: "wange", Age:22},
    }
    for _,stu := range stus{
        m[stu.Name = &stu
    }
    for _,v := range m{
        fmt.Println(v.Name, " ", v.Age)
    }
}
A.
zhou 24
li   23
wang 22

B.
wang 22
wang 22
wang 22

C.
zhou 24
zhou 24
zhou 24

D.
wang 22
li   23
li   23
~~~

7. 下面属于go语言的关键字是（）多选

~~~
A.func
B.def
C.struct
D.class
~~~

8. 定义一个包内全局字符串变量，下面语法正确的是（）多选

~~~
A.var str string
B.str := ""
C.str = ""
D.var str = ""
~~~

9. 通过指针变量p访问其成员变量name，下面语法正确的是（）多选

~~~
A.p.name
B.(*p).name
C.(&p).ma,e
D.p->name
~~~

10. 关于接口类的说法，下面说法正确的是（）多选

~~~
A.一个类只需要实现了接口要求的所有函数，我们就说这个类实现了该接口
B.实现类时只需要关心自己应该实现哪些方法即可
C.类实现接口时，需要导入接口所在的包
D.接口由使用方按自身需求来定义，使用方无需关心接口具体如何实现的
~~~

11. 关于字符串连接，下面说法正确的是（）多选

~~~
A.str:='abc'+'123'
B.str:="abc"+"123"
C.str : = '123'+"abc"
D.fmt.Sprintf("abc%d,123)
~~~



12. 关于协程，下面说法正确的是（）多选

~~~
A.协程和线程都可以实现程序的并发执行
B.线程比协程更轻量级
C.协程不存在死锁问题
D.通过channel来进行协程间的通信
~~~



13. 关于init函数，下面说法正确的是（）多选

~~~
A.一个包中，可以包含多个init函数
B.程序编译时，先执行导入包的init函数，再执行本包内的init函数
C.main包中，不能有init函数
D.init函数可以被其它函数调用
~~~



14. 关于循环语句，下面说法正确的是（）多选

~~~
A.循环语句既支持for关键字，也支持while和do-while
B.关键字for的基本使用方法与C/C++中没有任何差异
C.for循环支持continue和break来控制循环，但是它提供了一个更高级的label，可以选择终端哪个循环
D.for循环不支持以逗号为间隔的多个赋值语句
~~~



15. 对于变量声明下面那些是错误的（）多选

~~~
A.var 2abc string
B.姓名 := "小明"
C.var a$b int=123
D.var func int
~~~



16. 关于局部变量的初始化，下面正确的使用方式是（）多选

~~~
A.var i int = 10
B.var i = 10
C.i := 10
D.i = 10
~~~



17. 对add函数调用正确的是（）多选

~~~
func add(args ...int) int {
    sum := 0
    for _,arg := range args {
        sum += arg
    }
    return sum
}

A.add(1,2)
B.add(1,3,7)
C.add([]int{1,2})
D.add([]int{1,3,7}...)
~~~



18.	关于switch语句，下面说法正确的有（）多选

~~~
A.条件表达式必须为常量或者整数
B.switch中，可以出现多个条件相同的case
C.需要用break来明细退出一个case
D.只有在case中明确添加fallthrough关键字，才会继续执行紧跟的下一个case
~~~



19.	golang中的引用类型包括（）多选

~~~
A.切片
B.map
C.channel
D.interface
~~~



20. golang中的指针运算包括（）多选

~~~
A.可以对指针进行自增或自减运算
B.可以通过“&”取指针的地址
C.可以通过“*”取指针指向的数据
D.可以对指针进行下标运算
~~~



21. 关于main函数，下面说法正确的是（）多选

~~~
A.main 函数不能带参数
B.main 函数不能定义返回值
C.main 函数所在的包必须为main包才能作为函数入口
D.main 函数可以使用flag包来获取和解析命令行参数
~~~



22. 下面赋值正确的是（）多选

~~~
A. var x = nil
B. var x interface{} = nil
C. var x string = nil
D. var x error = nil
~~~



23. 关于整型切片的初始化，下面正确的是（）多选

~~~
A. s := make([]int)
B. s := make([]int, 0)
C. s := make([]int,5,10)
D. s := []int{1,2,3,4,5,}
~~~



24. 以下不能正确定义数组和赋初值的语句是（）多选

~~~
A. var a=[5]int(1,2,3,4,5)
B. var b=[...]int{1,2,3,4,5}
C. var c[5]int={1,2,3,4,5}
D. var d=[5]int{2:4,4:61}
~~~



25. 对于局部变量整型切片x的赋值，下面定义正确的是（）多选

~~~
A.
x := []int{
    1,2,3,
    4,5,6,
}

B.
x := []int{
    1,2,3,
    4,5,6
}

c.
x := []int{1,2,3,4,5,6,}
~~~



26. 关于变量的自增和自减操作，下面语句正确的是（）多选

~~~
A.
i := 1
i ++

B.
i := 1
j = i++

C.
i := 1
++i

D.
i := 1
i--
~~~



27. 关于GetPodAction定义，下面赋值正确的是（）多选

~~~
type Fragment interface {
    Exec(transInfo *TransInfo)error
}

type GetPodAction struct {}

func(g GetPodAction) Exec(transInfo *TransInfo)error{
    ...
    return nil
}

A.var fragment Fragment = new(GetPodAction)
B.var fragment Fragment = GetPodAction()
C.var fragment Fragment = &GetPodAction()
D.var fragment Fragment = GetPodAction{}
~~~



28. 关于接口，下面说法正确的是（）多选

~~~
A.只要两个接口拥有相同的方法列表，那么它们就说等价的，可以相互赋值
B.如果接口A的方法列表是接口B的方法列表的子集，那么接口B可以赋值给接口A
C.接口不可以继承
D.接口可以继承
~~~



29. 关于channel，下面语法正确的是（）多选

~~~
A.var ch chan int
B.ch := make(chan int)
C.<-ch
D.ch<-
~~~



30. 关于同步锁，下面说法正确的是（）多选

~~~
A.当一个goroutine活得Mutex后，其它goroutine就只能乖乖的等待，除非该goroutine释放这个Mutex
B.RWMutex在读锁占用的情况下，会阻止写，但不会阻止读
C.RWMutex在写锁占用情况下，会阻止任何其它goroutine（无论读和写）进来，整个锁相当于该goroutine独占
D.Lock()操作需要保证有Unlock()或RUnlock()调用与之对应，否则编译错误
~~~



31. golang中大多数数据类型都可以转化为有效的JSON文本，下面哪种类型除外（）多选

~~~
A.指针
B.channel
C.map
D.函数
~~~



32. flag是bool型变量，下面if表达式符合编码规范的（）多选

~~~
A.if flag == 1
B.if flag
C.if flag = false
D.if !flag
~~~



33. value是整型变量，下面if表达式符合编码规范的是（）多选

~~~
A.if value == 0
B.if value
C.if value != 0
D.if !value
~~~



34. 关于函数返回值的错误设计，下面说法正确的是（）多选

~~~
A. 如果失败原因只有一个，则返回bool
B. 如果失败原因超过一个，则返回error
C. 如果没有失败原因，则不返回bool或error
D. 如果重试几次可以避免失败，则不要立即返回bool或error
~~~

35. 关于slice或map操作，下面正确的是（）多选

~~~
A.
var s []int
s[0] = 6

B.
var m map[string]int
m["one"]= 1

C.
var s []int
s =make([]int, 0)
s =append(s,1)

D.
var m map[string]int
m =make(map[string]int)
m["one"]= 1
~~~

36. 关于channel的特性，下面说法正确的是（）多选

~~~
A. 给一个 nil channel 发送数据，造成永远阻塞
B. 从一个 nil channel 接收数据，造成永远阻塞
C. 给一个已经关闭的 channel 发送数据，引起 panic
D. 从一个已经关闭的 channel 接收数据，如果缓冲区中为空，则返回一个零值
~~~

37. 下面会触发异常的是（）多选

~~~
A. 判断空指针是否为nil
B. 下标越界
C. 除数为0
D. 调用panic函数
~~~

38. 关于cap函数的适用类型，下面说法正确的是（）多选

~~~
A. array
B. slice
C. map
D. channel
~~~

39. 关于select机制，下面说法正确的是（）多选

~~~
A. select机制用来处理异步IO问题
B. select机制最大的一条限制就是每个case语句里必须是一个IO操作
C. golang在语言级别支持select关键字
D. select块中有case子句，没有default子句
~~~

40.以下哪些类型可以使用make分配内存（）多选

~~~
A.Slice
B.Map 
C.chan
D.Struct
~~~

41. 以下属于线性表的是那些（）多选

~~~
A. 数组
B. 树
C. 链表
D. 图
~~~



### 填空题

1. 声明一个key为字符串型value为整数型的map变量m（）

2. 声明一个参数和返回值均为整型的函数变量f（）

3. 声明一个只用于读取int数据的单向channel变量ch（）

4. 假设源文件的命名为slice.go，则测试文件的命名为（）

5. go test要求测试函数的前缀必须命名为（）

6. 下面的程序的运行结果是（）

   ~~~go
   for i :=0; i < 5; i++{
       defer fmt.Printf("%d", i)
   }
   ~~~

7. 下面的程序的运行结果是（)

   ~~~go
   func main() {  
   	x := []string{"a", "b","c"}
       for v := range x {
         fmt.Print(v)
   	}
   }
   ~~~

   

8. 下面的程序的运行结果是（)

   ~~~go
   func main() {  
       i := 1
       j := 2
       i, j = j, i
       fmt.Printf("%d%d\n", i, j)
   }
   ~~~

   

9. 下面的程序的运行结果是（)

   ~~~
   type Slice []int
   
   func NewSlice() Slice {
   	return make(Slice, 0)
   }
   
   func (s Slice) Add(elem int) *Slice {
   	s = append(s, elem)
   	fmt.Print(elem)
   	return &s
   }
   
   func main() {
   	s := NewSlice()
   	defer s.Add(1).Add(2)
   	s.Add(3)
   }
   ~~~

   

### 判断题

1. interface是一个值类型（）

2. 使用map不需要引入任何库（）

3. map的遍历是有序的（）

4. 指针是基本类型（）

5. interface{}是可以指向任意对象的类型（）

6. Golang不支持自动垃圾回收（）

7. Golang支持反射，反射最常见的使用场景是做对象的序列化（）

8. 下面代码中两个斜点之间的代码，比如json:"x"，作用是X字段在从结构体实例编码到JSON数据格式的时候，使用x作为名字，这可以看作是一种重命名的方式（）

   ~~~go
   type Position struct {
       X int `json:"x"`
       Y int `json:"y"`
       Z int `json:"z"`
   }
   ~~~

   

9. Go通过成员变量或函数首字母的大小写来决定其作用域（）

10. 下面的程序的运行结果是xello（）

    ~~~go
    func main() {
        str := "hello"
        str[0] = ‘x‘
        fmt.Println(str)
    }
    ~~~

    

11. 匿名函数可以直接赋值给一个变量或者直接执行（）

12. 如果调用方调用了一个具有多返回值的方法，但是却不想关心其中的某个返回值，可以简单地用一个下划线“_”来跳过这个返回值，该下划线对应的变量叫匿名变量（）

13. 错误是业务过程的一部分，而异常不是（）

14. panic 比defer执行的优先级高（）



### 简答题

1. cap()和len()函数的区别是什么？
2. 如何将[]byte和[]rune类型的值转换为字符串？
3. 简述go语⾔中make和new的区别？
4. 简述闭包的⽣命周期和作⽤范围？
5. 测试⽂件必须以什么结尾？功能测试函数必须以什么开头？压⼒测试函数必须以什么开头？



### 编程题

1. 写一个函数判断一个字符串是否是回文(上海自来水来自海上)，要求不借助额外临时变量
2. 编程实现：使用2个协程交替执行,使其能顺序输出1-20的自然数，一个子协程输出奇数，另一个输出偶数
3. 如果a+b+c=1000，且a*a+b*b=c*c（a,b,c为自然数），如何求出所有a,b,c可能的组合？不使用数学公式和math包，代码实现，并且尽可能保证算法效率高效