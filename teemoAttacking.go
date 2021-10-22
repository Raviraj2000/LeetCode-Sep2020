package main

import ("fmt")

func findPoisonedDuration(timeSeries []int, duration int) int {
  res := 0
  n := len(timeSeries)

  if n == 0{
    return 0
  }

  for i := 0; i + 1 < n; i++{
    res = res +  min(duration, timeSeries[i+1] - timeSeries[i])
  }
  return res + duration
}

func min(a, b int) int{
  if a < b {
    return a
  }
  return b
}



func main () {
  timeSeries := []int{1, 2, 3, 4, 5}
  duration := 5

  fmt.Println(findPoisonedDuration(timeSeries, duration))
}
