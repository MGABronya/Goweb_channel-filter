// @Title  channel-filter
// @Description  用通道实现筛法求素数
// @Author  MGAronya（张健）
// @Update  MGAronya（张健）  2022-10-10 16:18
package main

import (
	"fmt"
)

// @title    IntegerGenerator
// @description   生成自增的整数
// @auth      MGAronya（张健）             2022-10-10 16:19
// @param     void
// @return    chan int			被传入的通道
func IntegerGenerator() chan int {
	var ch chan int = make(chan int)

	// TODO 开出一个goroutine
	go func() {
		for i := 2; ; i++ {
			// TODO 直到通道索要数据，才把i添加进通道
			ch <- i
		}
	}()
	return ch
}

// @title    IntegerGenerator
// @description   输入一个整数序列，筛出是number的倍数
// @auth      MGAronya（张健）             2022-10-10 16:20
// @param     in chan int, number int	输入一个整数序列，筛出是number的倍数
// @return    chan int					被传入的通道，筛出是number的倍数
func Filter(in chan int, number int) chan int {
	out := make(chan int)
	go func() {
		for {
			i := <-in
			// TODO 将不是number的倍数的数放入输出队列中
			if i%number != 0 {
				out <- i
			}
		}
	}()
	return out
}

// @title    main
// @description   找出100以内的所有素数
// @auth      MGAronya（张健）             2022-10-10 16:20
// @param     void
// @return    void
func main() {
	const max = 100
	// TODO 初始化一个整数生成器
	numbers := IntegerGenerator()

	// TODO 从生成器中抓取一个整数2，作为初始化整数
	number := numbers

	// TODO 用number作为筛子，当筛子超过max时结束筛选
	for number <= max {
		// TODO 打印素数
		fmt.Println(number)
		// TODO 筛掉number的倍数
		numbers = Filter(numbers, number)
		// TODO 更新筛子
		number = <-numbers
	}
}
