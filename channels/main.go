package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func SomeTask() int {
	time.Sleep(time.Second)
	return rand.Intn(100)
}

func main() {
	fmt.Println("Testing Channels")

	d := make(chan int)
	go func() {
		wg := sync.WaitGroup{}
		for i := 0; i < 1000; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				res := SomeTask()
				d <- res
			}()
		}
		wg.Wait()
		close(d)
	}()

	for n := range d {
		//n := <-n
		fmt.Printf("%d \n", n)
	}

}
