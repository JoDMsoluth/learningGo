package test24

import (
	"fmt"
	"time"
)

func push(c chan int) {
	i := 0
	for {
		time.Sleep(1 * time.Second)
		c <- i
		i++
	}
}

func Test24() {
	c := make(chan int)

	timerChan := time.After(10 * time.Second)
	tickTimerChan := time.Tick(2 * time.Second)

	go push(c)
	for {
		select {
		case v := <-c:
			fmt.Println(v)
		case <-timerChan:
			fmt.Println("timeout")
			return
		case <-tickTimerChan:
			fmt.Println("tick")
		default:
			fmt.Println("Idle")
			// time에는 tick(일정 시간마다 발생)과 after(일정 시간 이후에 발생)가 있다.
			time.Sleep(1 * time.Second)
		}
	}
}
