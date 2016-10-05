package main

import (
	"sync"
	"fmt"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go read(&wg)
	go write(&wg)


	wg.Wait()
	fmt.Println("all task completed")
}

func read(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("read done")
}

func write(wg *sync.WaitGroup)  {
	defer wg.Done()
	fmt.Println("write done")
}
