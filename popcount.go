package main

import (
	"fmt"
	"time"
)

var pc [256]byte

func main() {
	start_pop := time.Now()
	pop_count := PopCount(42)
	elapsed_pop := time.Since(start_pop)
	fmt.Println("Regular popcount")
	fmt.Println("Elapsed time: ", elapsed_pop)
	fmt.Println("Pop count of 42: ", pop_count)

	start_loop := time.Now()
	loop_count := PopCountLoop(42)
	elapsed_loop := time.Since(start_loop)
	fmt.Println("Loop popcount")
	fmt.Println("Elapsed time: ", elapsed_loop)
	fmt.Println("Pop count of 42: ", loop_count)

	start_SF := time.Now()
	SF_count := PopCountSF(42)
	elapsed_SF := time.Since(start_SF)
	fmt.Println("64 popcount")
	fmt.Println("Elapsed time: ", elapsed_SF)
	fmt.Println("Pop count of 42: ", SF_count)

	start_clear := time.Now()
	clear_count := PopCountClear(42)
	elapsed_clear := time.Since(start_clear)
	fmt.Println("Clear popcount")
	fmt.Println("Elapsed time: ", elapsed_clear)
	fmt.Println("Pop count of 42: ", clear_count)

}

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCountLoop(x uint64) int {
	var sum byte
	for i := uint(0); i < 8; i++ {
		sum += pc[byte(x>>(i*8))]

	}

	return int(sum)
}

func PopCountSF(x uint64) int {
	var sum int
	for i := uint(0); i < 64; i++ {
		if x&(x<<i) != 0 {
			sum++
		}
	}
	return sum
}

func PopCountClear(x uint64) int {
	var sum int
	for x != 0 {
		x = x & (x - 1)
		sum++
	}
	return sum
}
