package main

import "fmt"

//The two sum problem the go version
func main(){
	var nums []int = []int{2,4,7,6,8,9,3}
	result := twoSumDoubleForLoop(nums,13)
	resultDoubleForLoop := twoSumDoubleForLoop(nums,13)
	
	if result != nil {
		fmt.Printf("Map version version/Indices: %v\n", result) // Output: Indices: [0, 1]
	} else {
		fmt.Println("No two sum solution found.")
	}

	if resultDoubleForLoop != nil {
		fmt.Printf("Double for loop version/Indices: %v\n", resultDoubleForLoop) // Output: Indices: [0, 1]
	} else {
		fmt.Println("No two sum solution found.")
	}

}

func twoSumDoubleForLoop(nums []int, target int)[]int{
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] + nums[j] == target{
				return []int{i,j}
			}
			
		}

		
	}
	return nil

}

func twoSum(nums []int,target int)[]int{
	numMap := make(map[int]int)

	for i, num := range nums{
		complement := target - num

		if index, found := numMap[complement]; found{
			return []int{index,i} //returning the indices found
		}

		// Storing the number and its index in the map
		numMap[num] = i
	}
	return nil // returning nil if no pair is found
}