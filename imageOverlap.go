package main

import ("fmt")

func largestOverlap(A [][]int, B [][]int) int {
  if (A == B){
    return 9
  }
  if (A == nil || B == nil){
    return 0
  }




  return
}

func main() {
  A :=  [][]int{{1,1,0},
                {0,1,0},
                {0,1,0},}

  B := [][]int{{0,0,0},
               {0,1,1},
               {0,0,1},}

  answer1:= largestOverlap(A, B)
  fmt.Println(answer1)
}
