package main

func pipeline(n int) (in chan int, out chan int) {
	out = make(chan int)
	first := out
	for i := 0; i < n; i++ {
		in = out
		out = make(chan int)
		go func(in chan int, out chan int) {
			for v := range in {
				out <- v
			}
			close(out)
		}(in, out)
	}
	return first, out
}

func main() {
	in, out := pipeline(1000000)
	in <- 1
	<-out
	close(in)
}
