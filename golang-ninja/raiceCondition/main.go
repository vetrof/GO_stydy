package main

import (
	"fmt"
	"sync"
	"time"
)

type counter struct {
	count int
	mutex *sync.Mutex
}

func (c *counter) increment() {
	c.mutex.Lock()
	c.count++
	c.mutex.Unlock()
}

func (c *counter) value() int {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.count
}

func main() {
	c := counter{
		mutex: new(sync.Mutex),
	}

	for i := 0; i < 10000; i++ {
		go func() {
			c.increment()
		}()
	}
	time.Sleep(time.Second)
	fmt.Println(c.value())
}
