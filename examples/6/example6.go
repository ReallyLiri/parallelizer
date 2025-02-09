package main

import (
	"fmt"
	"time"

	"github.com/shomali11/parallelizer"
)

func main() {
	group := parallelizer.NewGroup(parallelizer.WithPoolSize(10), parallelizer.WithJobQueueSize(10))
	defer group.Close()

	for i := 1; i <= 10; i++ {
		group.Add(func(workerId int) {
			time.Sleep(time.Second)
		})

		fmt.Println("Job added at", time.Now().Format("04:05"))
	}

	err := group.Wait()

	fmt.Println()
	fmt.Println("Done")
	fmt.Printf("Error: %v", err)
}
