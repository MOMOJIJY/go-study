package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/trace"
	"time"
)

func main() {

	//traceTool()

	//debugRun()

	//myTest()

	myTest2()
}

func myTest2() {
	runtime.GOMAXPROCS(1)
	c := make(chan int, 1)

	go func() {
		for {
			time.Sleep(5 * time.Second)
			fmt.Println("run 1")
		}
	}()

	go func() {
		for {
			time.Sleep(5 * time.Second)
			fmt.Println("run 2")
		}
	}()

	go func() {
		for {
			time.Sleep(5 * time.Second)
			fmt.Println("run 3")
		}
	}()

	<-c

}

func myTest() {
	fmt.Println("start main")
	c := make(chan int, 1)
	go func() {
		fmt.Println("run go func")
		c <- 1
	}()
	<-c
	fmt.Println("end main")
}

// 通过debug查看GPM
// GODEBUG=schedtrace=1000 ./trace2
func debugRun() {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		fmt.Println("Hello World")
	}
}

// traceTool 通过go tool trace 查看GPM
func traceTool() {
	//创建trace文件
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	//启动trace goroutine
	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()

	//main
	fmt.Println("Hello World")
}