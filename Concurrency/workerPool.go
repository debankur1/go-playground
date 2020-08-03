package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)
type job struct {
	id int
	number int
}
//Job for which it holds the results
type Result struct {
	job job
	sumOfDigits int
}
var jobs = make(chan job, 10)
var results = make(chan Result,10)
func digits(number int) int {
	sum := 0
	no := number
	for no != 0 {
		digit := no % 10
		sum += digit
		no /= 10
	}
	time.Sleep(2 * time.Second)
	return sum
}
func Worker(wg *sync.WaitGroup){
	for job := range jobs{
		output := Result{
			job:         job,
			sumOfDigits: digits(job.number),
		}
		results <- output
	}
	wg.Done()
}
func CreateWorkerPool(numberOfWorkers int){
	var wg sync.WaitGroup
	for i:=0;i<numberOfWorkers;i++{
		  wg.Add(1)
		  go Worker(&wg)
	}
	wg.Wait()
	close(results)
}
func allocateJobs(noOfJobs int){
	for i:=0;i<noOfJobs;i++{
		 num := rand.Intn(999)
		 job := job{
			 id:     i,
			 number: num,
		 }
		 jobs <- job
	}
	close(jobs)
}
func readResult(done chan bool){
	for result := range results{
		fmt.Println(result)
	}
	done <- true
}
func main()  {
	go allocateJobs(10)
	done := make(chan bool)
	go readResult(done)
	CreateWorkerPool(2)
	<- done
}