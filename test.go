package simplebank

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int) {
	for j := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, j)
		time.Sleep(time.Second)
		fmt.Printf("Worker %d finished job %d\n", id, j)
	}
	fmt.Printf("Worker %d exiting\n", id)
}

func main() {
	jobs := make(chan int, 5)

	// Start 3 workers
	for w := 1; w <= 3; w++ {
		go worker(w, jobs)
	}

	// Send 5 jobs
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs) // signal no more jobs

	time.Sleep(6 * time.Second) // wait for all work to finish
}
