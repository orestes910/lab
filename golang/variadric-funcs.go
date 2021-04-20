package main

import "fmt"

func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		// keep adding num to total
		total += num
	}
	// After the range has finished
	fmt.Println(total)
}

func main() {
	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4}
	// slice/array nums must be fed into a variadric
	// function like this
	sum(nums...)
}
