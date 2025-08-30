package main

import (
	"fmt"
	"math"
	"sync"
	"sync/atomic"
	"time"
)

func main() {

	a := 45
	p := &a
	fmt.Println("method01：", *(method01(p)))

	slice := []int{3, 6, 7, 8, 2}
	sp := &slice
	method02(sp)
	fmt.Println("切片sp：", *sp)

	method03()
	time.Sleep(3 * time.Second)

	testTask()

	rg := &Rectangle{35, 64}
	fmt.Println("长方形面积：", rg.Area())
	fmt.Println("长方形周长：", rg.Perimeter())
	ce := &Circle{5.5}
	fmt.Println("圆形面积：", ce.Area())
	fmt.Println("圆形周长：", ce.Perimeter())

	//p := Person{"张三", 18}
	e := Employee{Person{"张三", 25}, "343233"}
	e.PrintInfo()

	//创建通道
	ch1 := make(chan int)

	go sendData(ch1)
	go receiveData(ch1)

	time.Sleep(2 * time.Second)

	ch2 := make(chan int, 10)
	go sendOnly(ch2)
	go receiveData(ch2)

	time.Sleep(2 * time.Second)

	count := SafeCounter{}
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				count.Increment()
			}
		}()
	}
	time.Sleep(5 * time.Second)
	fmt.Println("Final count：", count.GetCount())

	unlockCount()

}

// =======================指针============================================
// 1.题目 ：编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值。
// 考察点 ：指针的使用、值传递与引用传递的区别。
func method01(param *int) *int {
	*param += 10
	return param
}

// 2.题目 ：实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
// 考察点 ：指针运算、切片操作。
func method02(p *[]int) {
	for i, v := range *p {
		(*p)[i] = v * 2
	}
}

// =======================Goroutine============================================
// 1.题目 ：编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
// 考察点 ： go 关键字的使用、协程的并发执行。
func method03() {
	go func() {
		for i := 1; i <= 10; i++ {
			if i%2 != 0 {
				fmt.Printf("1-10的奇数,%d \n", i)
			}
		}

	}()

	go func() {
		for i := 2; i <= 10; i++ {
			if i%2 == 0 {
				fmt.Printf("2-10的偶数,%d \n", i)
			}
		}

	}()

}

// 2.题目 ：设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
// 考察点 ：协程原理、并发任务调度。
type Task func()

type TaskRes struct {
	taskIndex     int
	taskName      string
	executionTime time.Duration
}

func TaskScheduler(taskArr []Task) []TaskRes {
	results := make([]TaskRes, len(taskArr))
	var wg sync.WaitGroup
	for i, v := range taskArr {
		wg.Add(1)
		go func(index int, tk Task) {
			defer wg.Done()
			startTime := time.Now()
			//执行task任务
			tk()
			results[index] = TaskRes{
				i, "task", time.Since(startTime),
			}
		}(i, v)
	}

	wg.Wait()

	return results
}

func testTask() {
	tasks := []Task{
		// 任务1：模拟耗时操作
		func() {
			fmt.Println("任务1执行开始")
			time.Sleep(100 * time.Millisecond)
			fmt.Println("任务1执行完成")
		},
		// 任务2：模拟耗时操作
		func() {
			fmt.Println("任务2执行开始")
			time.Sleep(150 * time.Millisecond)
			fmt.Println("任务2执行完成")
		},
		// 任务3：模拟耗时操作
		func() {
			fmt.Println("任务3执行开始")
			time.Sleep(80 * time.Millisecond)
			fmt.Println("任务3执行完成")
		},
	}

	// 执行任务调度器
	fmt.Println("开始执行任务...")
	results := TaskScheduler(tasks)
	time.Sleep(2 * time.Second)
	fmt.Println("所有任务执行完成，结果如下：")

	for _, result := range results {
		fmt.Printf("任务 %d 名称 %v 执行时间: %v\n", result.taskIndex, result.taskName, result.executionTime)
	}
}

// =======================面向对象============================================
// 1.题目 ：定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
// 考察点 ：接口的定义与实现、面向对象编程风格。
type Shape interface {
	Area()
	Perimeter()
}
type Rectangle struct {
	length int //长度
	width  int //宽度
}

type Circle struct {
	radius float64 //半径
}

func (s *Rectangle) Area() int {
	return s.length * s.width
}

func (s *Rectangle) Perimeter() int {
	return (s.length + s.width) * 2
}

func (c *Circle) Area() float64 {
	return math.Pi * float64(c.radius) * float64(c.radius)
}

func (c *Circle) Perimeter() float64 {
	return math.Pi * float64(c.radius) * 2
}

// 2.题目 ：使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
// 考察点 ：组合的使用、方法接收者。
type Person struct {
	Name string
	Age  int
}

type Employee struct {
	p          Person
	EmployeeID string
}

func (e *Employee) PrintInfo() {
	fmt.Println("员工姓名:", e.p.Name, "员工年龄：", e.p.Age)
	fmt.Println("员工工号:", e.EmployeeID)
}

// =======================Channel============================================
// 1.题目 ：编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
// 考察点 ：通道的基本使用、协程间通信。
func sendData(ch chan<- int) {
	for i := 1; i <= 10; i++ {
		ch <- i
		fmt.Printf("sendData：%d \n", i)
	}
	close(ch)
}

func receiveData(ch <-chan int) {
	for v := range ch {
		fmt.Printf("receiveData：%d \n", v)
	}

}

// 2.题目 ：实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
// 考察点 ：通道的缓冲机制。
func sendOnly(ch chan<- int) {
	for i := 1; i <= 100; i++ {
		ch <- i
		fmt.Printf("sendOnly：%d \n", i)
	}
	close(ch)
}

// =======================锁机制============================================
// 1.题目 ：编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
// 察点 ： sync.Mutex 的使用、并发数据安全。
type SafeCounter struct {
	count int
	mu    sync.Mutex
}

func (c *SafeCounter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}
func (c *SafeCounter) GetCount() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

// 2.题目 ：使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
// 考察点 ：原子操作、并发数据安全。

func unlockCount() {
	var count int64 = 0

	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				atomic.AddInt64(&count, 1)
			}
		}()
	}
	time.Sleep(3 * time.Second)
	fmt.Println("unlockCount方法计算结果：", count)

}
