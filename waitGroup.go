package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	//fmt.Println(runtime.CPUProfile())
	fmt.Println(runtime.GOARCH)
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOROOT())
	fmt.Println(runtime.NumCPU())
	go func() {
		fmt.Println("in first func")
		wg.Done()
	}()
	go func() {
		fmt.Println("in second func")
		wg.Done()
	}()
	fmt.Println("SUJAY")
	///New comment

	wg.Wait()
}
