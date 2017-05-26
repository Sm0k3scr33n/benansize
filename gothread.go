package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.Tick(time.Second/30)
	for i := 1; i <= 30; i++ {
		<-ticker
		// fmt.Printf("\x0cOn %d/30", i)\
    var bla = "my name is ben"
    fmt.Printf("\r %d/30 %q", i, bla)
  

	}

	fmt.Println("\nAll is said and done.")
}
