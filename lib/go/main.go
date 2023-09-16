package main

import (
	"fmt"
)

func main() {
	fmt.Println("Go loaded.")

	<-make(chan bool)
}
