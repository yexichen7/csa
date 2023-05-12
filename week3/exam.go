// 这个示例程序展示如何使用
// 有缓存的通道和固定数目的
// goroutine来处理工作
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numberGoroutines = 4  //要使用go程的数目
	taskLoad         = 10 //要处理工作的数量
)

// wg用于等待程序完成
var wg sync.WaitGroup

// init初始化包，go语言会在运行其他代码之前优先执行这个函数
func init() {
	//初始化随机数种子
	rand.Seed(time.Now().Unix())
}

func main() {
	//创建一个有缓冲的通道来管理工作
	tasks := make(chan string, taskLoad)

	//启动go程来处理工作
	wg.Add(numberGoroutines)
	for gr := 1; gr <= numberGoroutines; gr++ {
		go worker(tasks, gr)
	}

	//增加一组要完成的工作
	for post := 1; post <= taskLoad; post++ {
		tasks <- fmt.Sprintf("Task : %d", post)
	}

	//当所有工作处理完关闭通道，以便所有goroutine退出
	close(tasks)

	//等待所有工作完成
	wg.Wait()
}

// worker作为goroutine启动来处理
// 从有缓存的通道传入工作
func worker(tasks chan string, worker int) {
	//通知函数已经返回
	defer wg.Done()

	for {
		//等待分配工作
		task, ok := <-tasks
		if !ok {
			fmt.Printf("Worker: %d :Shutting Down \n", worker)
			return
		}

		//显示我们开始工作了
		fmt.Printf("Woker : %d : Started %s \n", worker, task)

		//随机等待一段时间来模拟工作
		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep) * time.Millisecond)

		//显示我们完成了工作
		fmt.Printf("Worker: %d :Completd %s \n", worker, task)
	}
}
