package main

import "fmt"
import "time"

func main() {
	c := 0
	start := time.Now()
	for i, j := 1, 1; j <= 10; {
		if j <= i*j && i*j <= j*10 {
			fmt.Print(i*j, " ")
			i++
			c++
		} else {
			fmt.Println()
			j++
			i = 1
			c++
		}
	}
	fmt.Println(c)
	duration := time.Since(start)
	fmt.Println(duration)

}
