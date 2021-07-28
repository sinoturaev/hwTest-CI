package main

import "fmt"

func SumOfSlice(elements []int) (int, float64) {
	sum := 0
	for _, value := range elements {
		sum = sum + value
	}

	fmt.Println("Sum of slices elements=",sum)

	 var average float64 = float64(sum) / float64(len(elements))
	 fmt.Println("Averagee of slice=", average)

	 return sum, average
}

func main() {
	SumOfSlice([]int{34, 21, 48, 92, 84})
}
