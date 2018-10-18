package main

import (
	"fmt"
	"time"

	"./popcount"
)

func main() {
	start_pop := time.Now()
	pop_count := popcount.PopCount(42)
	elapsed_pop := time.Since(start_pop)
	fmt.Println("Regular popcount")
	fmt.Println("Elapsed time: %d", elapsed_pop)
	fmt.Println("Pop count of 42: %d", pop_count)

	start_loop := time.Now()
	loop_count := popcount.PopCountLoop(42)
	elapsed_loop := time.Since(start_loop)
	fmt.Println("Loop popcount")
	fmt.Println("Elapsed time: %d", elapsed_loop)
	fmt.Println("Pop count of 42: %d", loop_count)
}
