package main

import (
	"fmt"
	"time"
)

func main() {
	runWorkerPool(3, 5)
}

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("Worker %d recebeu o Job %d\n", id, j)

		time.Sleep(time.Second * 2) // Simula o processamento
		result := j * 2             // eh nesse ponto que se chamaria a funcao/metodo de processamento dos dados do job no worker

		fmt.Printf("Worker %d finalizou o Job %d com o resultado %d\n", id, j, result)
		results <- result
	}
}

func runWorkerPool(numWorkers int, jobSize int) {
	jobs := make(chan int, jobSize)
	results := make(chan int, jobSize)

	for w := 1; w <= numWorkers; w++ {
		go worker(w, jobs, results)
	}

	for j := 0; j <= jobSize; j++ {
		jobs <- j
	}
	close(jobs)

	for r := 1; r <= jobSize; r++ {
		<-results
	}
}
