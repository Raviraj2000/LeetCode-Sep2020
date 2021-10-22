package main

import ("fmt"
        "strconv"
        "strings")

func majorityElement(nums []int) []int {

  if len(nums) == 0{
    return nil
  }

  numDict := make(map[int]int)
  limit := len(nums)/3
  var word string
  var answer []int

  for _,number := range nums{
    numDict[number] = numDict[number] + 1
    if numDict[number] > limit{
      if (strings.Contains(word, strconv.Itoa(number))){
        continue
      }
      answer = append(answer, number)
      word = word + strconv.Itoa(number)
    }
  }
  return answer
}


func main (){
  nums := []int{2, 2}

  fmt.Println(majorityElement(nums))
}
