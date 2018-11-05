package main

import (
	"fmt"
	"sync"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)
	var wg sync.WaitGroup

	numbers := []uint32{3245534, 2345444, 23345999, 32323, 20004395, 2445322, 234, 667544, 65677, 7467657}
	wg.Add(len(numbers))

	fmt.Println("Comienza el proceso...")
	for _, v := range numbers{
		go func(a uint32) {
			defer wg.Done()
			fmt.Println(a, EsPrimo(a))
		}(v)
	}

	wg.Wait()
	fmt.Println("Terminó el proceso")
}

func EsPrimo(a uint32) bool {
	c := 0
	var i uint32
	for i = 1; i <= a; i++ {
		if a%i == 0 {
			c++
		}
	}
	if c == 2 {
		return true
	}
	return false
}