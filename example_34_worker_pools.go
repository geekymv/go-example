package main

import (
	"fmt"
	"time"
)

func myWorker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}


func main() {

	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// 每个 job 执行3次
	for id := 1; id <=3; id++ {
		go myWorker(id, jobs, results)
	}

	time.Sleep(time.Second * 5)

	// 5个 job
	for j :=1;j <= numJobs; j++ {
		jobs <- j
	}

	close(jobs)

	for j := 1; j <= numJobs; j++ {
		<-results
	}

}
