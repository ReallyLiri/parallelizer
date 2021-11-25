package main

import (
	"fmt"

	"github.com/shomali11/parallelizer"
)

func main() {
	group := parallelizer.NewGroup()
	defer group.Close()

	group.Add(func(workerId int) {
		for char := 'a'; char < 'a'+3; char++ {
			fmt.Printf("%c ", char)
		}
	})

	group.Add(func(workerId int) {
		for number := 1; number < 4; number++ {
			fmt.Printf("%d ", number)
		}
	})

	err := group.Wait()

	fmt.Println()
	fmt.Println("Done")
	fmt.Printf("Error: %v", err)
}
