package main

import "fmt"

// Andrew_and_Acid
func main() {
	fmt.Println("hello")
	//input := 2
	var slice []int
	slice = append(slice, 1, 1, 5, 5, 5)

	calc(slice)
}

func calc(slice []int) {
	var count int
	for i := 0; i < len(slice); i++ {
		if i >= 1 {
			if slice[i] > slice[i-1] {
				count = slice[i] - slice[i-1]
			} else if slice[i] < slice[i-1] {
				count = -1
			}
		}
	}
	fmt.Println(count)
}
