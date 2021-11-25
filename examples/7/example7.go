package main

import (
	"fmt"
	"time"

	"github.com/shomali11/parallelizer"
)

func main() {
	group := parallelizer.NewGroup()
	defer group.Close()

	group.Add(func(workerId int) {
		fmt.Println("Finished work")
	})

	fmt.Println("We did not wait!")

	time.Sleep(time.Second)
}
