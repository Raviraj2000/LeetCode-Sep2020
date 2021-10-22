package main

import ("fmt")

func maxProfit(prices []int) int {
  if len(prices) == 0 || len(prices) == 1{
    return 0
  }

  min1, index1 := findMin(prices)
  if (min1 == -1 && index1 == -1){
    return 0
  }
  max1 := findMax(prices[index1:])

  min2, index2 := findMin(prices[:index1])
  if min2 == -1 && index2 == -1{
    return max1 - min1
  }
  max2 := findMax(prices[index2:index1])

  fmt.Println(min1,index1)
  fmt.Println(min2,index2)

  ans1 := max1-min1
  ans2 := max2-min2

  if ans1 > ans2{
    return ans1
  } else{
    return ans2
  }
  return 0
}

func findMin(prices []int) (int,int){
  if len(prices) == 0 || len(prices) == 1{
    return -1,-1
  }
  min := prices[0]
  index := 0
  for i, pi := range prices{
    if pi < min{
      min,index = pi,i
    }
  }
  if index == len(prices) - 1{
    min, index = findMin(prices[:len(prices)-1])
  }
  return min, index
}

func findMax(prices []int) int{
  max := prices[0]
  for _,pi := range prices{
    if pi > max{
      max = pi
    }
  }
  return max
}


func main() {
  prices := []int{7,6,4,3,1}

  fmt.Println(maxProfit(prices))
}
