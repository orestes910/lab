package main

import (
	"fmt"
	"time"
)

// accepts an integer id, an inbound
// channel, and an outbound channel.
// I think this function when called will
// run until both jobs and results are closed/received
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		// I'm not sure why we're multiplying the job
		// int by 2 when sending the result
		results <- j * 2
	}
}

func main() {

	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// Call the worker func 3 times, passing
	// in the id and both channels
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// send 5 jobs to the jobs channel and close it
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= numJobs; a++ {
		fmt.Println(<-results)
	}
}
