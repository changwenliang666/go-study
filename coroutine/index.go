package main

import (
	"fmt"
	"sync"
)

func main() {
	/**
	var ch = make(chan int) // 无缓冲管道，读或者写的时候，必须有一方接收
	defer close(ch)
	go func() {
		ch <- 1
	}()
	fmt.Println(<-ch)
	*/

	var group sync.WaitGroup
	var group2 sync.WaitGroup
	group.Add(10)

	for i := 0; i < 10; i++ {
		group2.Add(1)
		go func() {
			fmt.Println(i)
			group.Done()
			group2.Done()
		}()
		group2.Wait()
	}

	group.Wait()
	fmt.Println("所有协程执行完毕")

}
