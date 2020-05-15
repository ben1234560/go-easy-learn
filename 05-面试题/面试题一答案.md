# 面试题一答案

### 选择题

1. C
2. C
   1. 解析：局部操作的变量，无法影响全局变量
3. C
4. C
   1. 解析：go语法，有名字就都有，没名字就都没有
5. D
6. B
7. AC
8. AD
9. AB
10. ABD
11. BD
12. AD
13. AB
14. CD
15. ACD
16. ABC
17. ABD
18. BD
19. ABC
20. BC
21. ABCD
22. BD
23. BCD
    1. 需要初始化长度
24. AC
25. AC
26. AD
    1. 支持后置++--，不支持前置
27. ACD
28. AD
29. ABC
30. ABC
31. BD
32. BD
33. AC
34. ABD
35. CD
36. ABCD
37. BCD
38. ABD
39. ABC
40. ABC
41. AC

### 填空题

1. var m map[string]int
2. var f func(int)int
3. var ch<-chan int
4. slice_test.go
5. Test
6. 43210
7. 012
8. 21
9. 132

### 判断题

1. T
2. T
3. F
4. F
5. T
6. F
7. T
8. T
9. T
10. F
11. T
12. T
13. T
14. T

### 简答题

1. cap指底层数组的大小，len指实际长度大小

2. ing

3. Make用于map、slice和channel几种类型的内存分配。并返回一个有初始值的对象，注意部署指针。New用于type声明的类型的内存分配

4. ing

5. _test.go

   Test

   Benchmark

### 编程题

1. ~~~go
   package main
   
   import "fmt"
   
   func ishuiwen(str string) bool {
      if len(str) == 0 {
         return false
      }
      var r []rune = []rune(str)
      // 获取前后索引
      i, j := 0, len(r)-1
      for i < j {
         if r[i] == r[j] {
            i++
            j--
         } else {
            return false
         }
      }
      return true
   }
   
   func main() {
      fmt.Println(ishuiwen("上海自来水来自海上"))
   }
   ~~~

2. ~~~go
   package main
   
   import (
      "fmt"
      "sync"
   )
   
   var wg = sync.WaitGroup{}
   
   func main() {
      // 放偶数管道
      ch1 := make(chan int)
      // 存奇数
      ch2 := make(chan int)
      wg.Add(2)
      go m1(ch2, ch1)
      go m2(ch1, ch2)
      wg.Wait()
   }
   
   func m1(ch2 chan int, ch1 chan int) {
      defer wg.Done()
      for i := 1; i <= 10; i++ {
         ch2 <- 2*i - 1
         fmt.Println("偶数协程：", <-ch1)
      }
   }
   
   func m2(ch1 chan int, ch2 chan int) {
      defer wg.Done()
      for i := 1; i <= 10; i++ {
         fmt.Println("奇数协程：", <-ch2)
         ch1 <- 2 * i
      }
   }
   ~~~

3. 

3. ~~~go
   package main
   
   import (
      "fmt"
      "sync"
   )
   
   var wg = sync.WaitGroup{}
   
   func main() {
      // 放偶数管道
      ch1 := make(chan int)
      // 存奇数
      ch2 := make(chan int)
      wg.Add(2)
      go m1(ch2, ch1)
      go m2(ch1, ch2)
      wg.Wait()
   }
   
   func m1(ch2 chan int, ch1 chan int) {
      defer wg.Done()
      for i := 1; i <= 10; i++ {
         ch2 <- 2*i - 1
         fmt.Println("偶数协程：", <-ch1)
      }
   }
   
   func m2(ch1 chan int, ch2 chan int) {
      defer wg.Done()
      for i := 1; i <= 10; i++ {
         fmt.Println("奇数协程：", <-ch2)
         ch1 <- 2 * i
      }
   }
   ~~~

4. 