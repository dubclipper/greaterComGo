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
	var m sync.Mutex //mutual exclusion to keep many goroutines from accessing a variable at the same time

	for i := 0; i < gort; i++ {
		go func() {
			m.Lock() //locks it down so that only 1 goroutine can access 'increment'
			v := increment
			v++
			increment = v
			fmt.Println(increment)
			m.Unlock() //makes 'increment' available again; unlocks for other goroutines
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("final:", increment)
}
