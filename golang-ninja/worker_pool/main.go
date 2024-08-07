package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, result chan<- int) {
	for j := range jobs {
		time.Sleep(time.Second / 2)
		fmt.Printf("worker %d", id)
		result <- j * j
	}
}

func main() {
	const JobCount = 15
	const Workers = 5

	jobs := make(chan int, JobCount)
	result := make(chan int, JobCount)

	for i := 0; i < Workers; i++ {
		go worker(i, jobs, result)
	}

	//	add nums in jobs
	for i := 0; i < JobCount; i++ {
		jobs <- i + 1
	}
	close(jobs)

	//	read result
	for i := 0; i < JobCount; i++ {
		fmt.Printf("result %d, value %d\n", i, <-result)
	}

}
