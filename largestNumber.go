package main

import ("fmt"
        "strconv"
        "sort"
        //"math"
      )

func largestNumber(nums []int) string {
  
  sort.Sort(sort.Reverse(sort.IntSlice(nums)))
  fmt.Println(nums)
  var answer string
  var index int
  var mid int
  for i,li := range nums{
    if li < 10{
      mid = i
      break
    }
  }
  fmt.Println()
  index = mid
  for i := 0; i < len(nums); i++{
    if index == len(nums){
      for j := i; j <= mid - 1; j++{
        answer = answer + strconv.Itoa(nums[j])
        }
        break
    }
    if i == mid{
      for j := index; j < len(nums); j++{
        answer = answer + strconv.Itoa(nums[j])
      }
      break
    }
    number := nums[index] * (10 ^ (len(strconv.Itoa(nums[i])) - len(strconv.Itoa(nums[index]))))
    if nums[i] > number{
      answer = answer + strconv.Itoa(nums[i])
      continue
    } else if number > nums[i]{
        answer = answer + strconv.Itoa(nums[index])
        index++
        i = i - 1
        continue
    }

  }

  return answer
}



func main() {
  nums := []int{121, 12}

  fmt.Println(largestNumber(nums))
}
