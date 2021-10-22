package main

import ("fmt")

func findMaximumXOR(nums []int) int {
  if len(nums) == 2{
    return nums[0]^nums[1]
  }
  var max int
  var k int
  for i := 0; i < len(nums); i++{
    ans := nums[i-k]^nums[k]
    if max < ans{
      max = ans
    }
    k ++
    if (k == len(nums) - 1 ){
      i = len(nums) - k
      k = 0
    }
  }
  return max
}


func main() {
  nums := []int{15, 15, 9, 3, 2}

  fmt.Println(findMaximumXOR(nums))

}
