package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("begin cpu", runtime.NumCPU())
	fmt.Println("begin.gs", runtime.NumGoroutine())

	var waitgrp sync.WaitGroup
	waitgrp.Add(2)

	go func() {
		fmt.Println("Here's the first goroutine in the program")
		waitgrp.Done()
	}()

	go func() {
		fmt.Println("Here's the second goroutine in the program")
		waitgrp.Done()
	}()

	fmt.Println("mid cpu", runtime.NumCPU)
	fmt.Println("mid gs", runtime.NumGoroutine)

	waitgrp.Wait()

	fmt.Println("exiting")
	fmt.Println("end cpu", runtime.NumCPU)
	fmt.Println("end gs", runtime.NumGoroutine)
}
