package main

func main() {
	bufferSize := 5
	ch := make(chan int, bufferSize)
	go producer(ch, bufferSize)
	consumer(ch)
}

func producer(ch chan<- int, bufferSize int) {
	for i := 0; i < bufferSize; i++ {
		ch <- i
	}
	close(ch)
}

func consumer(ch <-chan int) {
	for v := range ch {
		println("consumido:", v)
	}
}
