package main

import "fmt"

func main() {
	evn := make(chan int)
	odd := make(chan int)
	qit := make(chan int)

	//send
	go send(evn, odd, qit)

	//receive
	receive(evn, odd, qit)

	fmt.Println("time to exit!")
}

func receive(e, o, q <-chan int) {
	for {
		select {
		case v := <-e:
			fmt.Println("from the evn channel:", v)
		case v := <-o:
			fmt.Println("from the odd channel:", v)
		case v := <-q:
			fmt.Println("from the quit channel:", v)
			return
		}
	}
}

func send(e, o, q chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	q <- 0
}
