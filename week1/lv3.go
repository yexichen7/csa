package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	maxNum := 100
	var guess int
	flag := false
	// 使用一直在不断变化的时间作为我们的种子（时间戳）
	rand.Seed(time.Now().UnixNano())
	// 设置种子之后产生一个最大为100的整形
	secretNumber := rand.Intn(maxNum)
	fmt.Println("Please input your guess")
	for i := 0; i < 5; i++ {

		// 输入我们猜的数字
		_, err := fmt.Scan(&guess)
		// Go语言中处理错误的方法
		if err != nil {
			fmt.Println("Invalid input. Please enter an integer value")
			return
		}

		if guess > secretNumber {
			fmt.Println("You guess ", guess, "is larger than the secret number")
		} else if guess < secretNumber {
			fmt.Println("You guess ", guess, "is less than the secret number")
		} else {
			fmt.Println("You guess ", guess, "is right")
			flag = true
		}
		if flag {
			break
		}
	}
	if !flag {
		fmt.Println("You fail! The secret number is ", secretNumber) // 作弊模式
	}
}
