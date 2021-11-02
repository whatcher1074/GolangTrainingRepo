package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {

	fmt.Println("Starting App")
	fmt.Println("OS\t", runtime.GOOS)
	fmt.Println("ARCH\t", runtime.GOARCH)
	fmt.Println("CPUs \t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	wg.Add(1)
	go crap()
	somemorecrap()

	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	wg.Wait()

}

func crap() {
	fmt.Println("A bunch of crap! lolol")
	wg.Done()
}

func somemorecrap() {

	fmt.Println("Some more crap for testing!")
}
