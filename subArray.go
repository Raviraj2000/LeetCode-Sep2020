package main

import ("fmt"
        "math")

func maxProduct(nums []int) int {
  if len(nums) == 1{
    return nums[0]
  }

  subArray := subArray(nums)
  fmt.Println(subArray)
  for i, li := range(subArray){
    if len(li) == len(nums){
      subArray = append(subArray[:i], subArray[i+1:]...)
    }
}
  max := math.Inf(-1)
  var product int
  for _, li := range(subArray){
    product = 1
    for j := 0; j < len(li); j++{
      product *= li[j]
    }
    fproduct := float64(product)
    if max < fproduct{
      max = fproduct
    }
  }

  return int(max)
}


func subArray(nums []int) [][]int{

    var subArray [][]int
    for i := 0; i < len(nums)+1;i++{
      for j := i+1; j < len(nums) + 1; j ++{
            subArray = append(subArray, nums[i:j])
          }
  }

  return subArray
}

func main() {
  nums := []int{-4, -3}

  fmt.Println(maxProduct(nums))
}
