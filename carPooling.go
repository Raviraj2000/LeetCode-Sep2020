package main

import ("fmt")

func carPooling(trips [][]int, capacity int) bool {
  location := [1001]int{}
  for _, t := range trips {
    num, start, end := t[0], t[1], t[2]
    location[start] += num
    location[end] -= num
}

  p := 0
  for _, c := range location {
    p += c
    if p > capacity {
      return false
  }
}
  return true
}


func main (){
  trips := [][]int{{9, 3, 6}, {8, 1, 7}, {6, 6, 8}, {8, 4, 9}, {4, 2, 9}}
  capacity := 28

  fmt.Println(carPooling(trips, capacity))
}
