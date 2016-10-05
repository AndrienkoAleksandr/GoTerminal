package main

import (
	"fmt"
)

func read(done chan bool) {
	done <- true
	fmt.Println("read done")
}

func write(done chan bool)  {
	done <- true
	fmt.Println("write done")
}


func main() {
	done := make(chan bool, 100)
	for i := 0; i < 34; i++ {
		done <- true
	}
	write(done)
	read(done)

	fmt.Println(len(done))
}

