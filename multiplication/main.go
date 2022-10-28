package main

import "fmt"
import "time"

func duration(f func()) {
	start := time.Now()
	f()
	duration := time.Since(start)
	fmt.Println(duration)
}
func my() {
	for i, j := 1, 1; j <= 10; {
		if j <= i*j && i*j <= j*10 {
			fmt.Print(i*j, " ")
			i++
		} else {
			fmt.Println()
			j++
			i = 1
		}
	}
}
func forInFor() {
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 10; j++ {
			fmt.Print(j, " ")
		}
		fmt.Println()
	}
}
func main() {
	duration(my)
	duration(forInFor)
}
