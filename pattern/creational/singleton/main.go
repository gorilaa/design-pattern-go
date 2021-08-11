package main

import (
	"fmt"

	"design-pattern/pattern/creational/singleton/single"
)

func main() {
	for i := 0; i < 30; i++ {
		go single.GetInstance()
	}

	fmt.Scanln()
}
