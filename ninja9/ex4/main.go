package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	increment := 0
	gort := 100
	wg.Add(gort)
	var m sync.Mutex

	for i := 0; i < gort; i++ {
		go func() {
			m.Lock()
			v := increment
			v++
			increment = v
			fmt.Println(increment)
			m.Unlock()

			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("final:", increment)
}
