package main

import ("fmt"
        "math")

func rob(nums []int) int {
  if len(nums) == 0{
    return 0
  }
  if len(nums) == 1{
    return nums[0]
  }
  value1 := float64(nums[0])
  value2 := math.Max(float64(nums[0]), float64(nums[1]))
  if len(nums) == 2{
    return int(value2)
  }

  var max float64

  for i := 2; i < len(nums); i++{
    max = math.Max(float64(nums[i])+value1, float64(value2))
    value1 = value2
    value2 = max
  }

  return int(max)

  }


func main() {

  //nums := []int{1, 2, 3, 1}
  nums := []int{2, 7, 9, 3, 1}

  fmt.Println(rob(nums))
}
