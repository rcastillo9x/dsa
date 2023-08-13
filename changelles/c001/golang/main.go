package main

import "fmt"

func main() {
	arr := []int{1, 3, 5, 7, 8}
	fmt.Println(getSum(arr, 12))

}

func getSum(arr []int, target int) (int, int) {
	keys := make(map[int]int)

	for index, number := range arr {
		complement := target - number
		_, ok := keys[complement]
		if ok {
			return index, keys[complement]
		}
		keys[number] = index
	}
	return 0, 0

}
