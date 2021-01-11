package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	increment := 0
	gort := 100
	wg.Add(gort)

	for i := 0; i < gort; i++ {
		go func() {
			v := increment
			runtime.Gosched()
			v++
			increment = v
			fmt.Println(increment)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("final:", increment)
}
