package main

import ("fmt")

func combinationSum3(k int, n int) [][]int {
  //i->1 to 9
  //length->k
  numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
  var numbers [][]int
  var temp []int
  var total int
  for i, ni := range numbers{
    if len(temp) > k || sum(temp) > n{
      temp = nil
    }

    if total + si < n{
      temp = append(temp, i)
    }
  }
return numbers
  }

func sum(temp []int) int{
  var sum int
  for i := 0; i < len(temp); i++{
    sum = sum + temp[i]
  }
  return sum
}



func main() {
  k := 3
  n := 7
  fmt.Println(combinationSum3(k, n))

}
