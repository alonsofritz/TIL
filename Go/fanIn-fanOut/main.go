package main

func main() {

}

func fanIn(chs ...<-chan int) <-chan int {
	out := make(chan int)
	for _, ch := range chs {
		go func(c <-chan int) {
			for v := range c {
				out <- v
			}
		}(ch)
	}
	return out
}
