package main

import "fmt"

func main() {
	var nums = make([]int, 5)

	fmt.Println(cap(nums))

	nums = append(nums, 1)
	nums = append(nums, 2)
	nums = append(nums, 3                       )
	nums = append(nums, 1)
	nums = append(nums, 2)

	fmt.Println("caps", cap(nums))

	fmt.Print(nums)
}
