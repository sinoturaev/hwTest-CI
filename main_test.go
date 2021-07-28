package main

import "testing"

func TestSumOfSlice(t *testing.T) {
	tests := []struct {
		name string
		elements []int
		sumWant int
		averageWant float64
	}{
	{"Elements are positive",[]int{12, 23, 41, 22, 44},142, 28.4},
	{"Elements are negative",[]int{-2, -5, -33, -11, -42},-93, -18.6},
	{"Elements are positive and negative",[]int{-22, -13, 4, 5, 9},-17,-3.4},

	}
	for _, test := range tests {
		sum, average := SumOfSlice(test.elements)
		if  sum != test.sumWant || average != test.averageWant{
			t.Errorf("Sum of slice test %s gotSum %d sumWant %d gotAverage%f avaregeWant %f",
				test.name, sum, test.sumWant, average, test.averageWant)
		}

	}
}
