package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	counter = 0
	lock    sync.Mutex
)

func main() {
	for i := 0; i < 10; i++ {
		go artir()
	}
	time.Sleep(time.Millisecond * 10)
}

func artir() {
	lock.Lock()
	counter++
	fmt.Println(counter)
	lock.Unlock()
}
