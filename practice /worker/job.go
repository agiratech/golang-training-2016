package main

import (
 "fmt"
 "time"
)

// Run a job - blocks until the job is done
func runJob(jobQueue chan chan string, name string) {

 // our job is a series of tasks
 taskQueue := make(chan string)
 defer close(taskQueue)

 // Queue our tasks in the job queue
 jobQueue <- taskQueue

 // job queue made room for us - queue up each of our tasks
 fmt.Printf("Starting job %s\n", name)
 for i := 1; i <= 5; i++ {
  taskQueue <- fmt.Sprintf("%s: %d", name, i)
  time.Sleep(1000 * time.Millisecond)
 }
 fmt.Printf("Job %s done!\n\n", name)

}

func main() {

 // This is our job queue - a channel of channels. We write channels of tasks into it.
 jobQueue := make(chan chan string)
 defer close(jobQueue)

 // simulate multiple requests to use the job queue
 fmt.Println("Begin: Kicking off workers that want to use the job queue")
 go runJob(jobQueue, "A")
 go runJob(jobQueue, "B")
 go runJob(jobQueue, "C")
   fmt.Println("End: Kicking off workers that want to use the job queue\n")

 // this is just for the example - this would be handled differently in a real service
 remainingJobs := 3

 // handle each job in the queue
 for job := range jobQueue {

  // Execute each of the tasks
  for task := range job {
   fmt.Printf("Handling: %s\n", task)
  }

  // break out when done
  remainingJobs--
  if remainingJobs == 0 {
   break
  }

 }
}
