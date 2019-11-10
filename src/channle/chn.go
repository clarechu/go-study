package channle

import (
	"os"
	"time"
)
// 设置time的延时队列
func NewChannleTimeout()  {
	c := make(chan int, 100)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
			time.Sleep(time.Second)
		}

		os.Exit(0)
	}()

	for {
		select {
		case n := <-c:
			println(n)
		case <-timeAfter(time.Second * 2):

		}
	}
}

func timeAfter(d time.Duration) chan int {
	q := make(chan int, 1)
	time.AfterFunc(d, func() {
		q <- 1
		println("run") 		// 重点在这里
	})

	return q
}

func TimeOut()  {
	c := make(chan int, 30)
	to := time.NewTimer(time.Second)
	for {
		to.Reset(time.Second)
		select {
		case <-c:
			println(c)
		case <-to.C:
		}
	}
}