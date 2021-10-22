package main

import ("fmt")

func firstMissingPositive(nums []int) int {
  if len(nums) == 0{
    return 1
  }
  numbers := make(map[int]int)
  for _,li := range(nums){
    numbers[li] = 1
  }
  fmt.Println(numbers)
  var answer int
  for i := 1; i <= 100000; i++{
    if numbers[i] != 1{
      answer = i
      break
    }
  }
  return answer
}


func main() {
  nums := []int{3,4,-1,1}

  fmt.Println(firstMissingPositive(nums))
}
