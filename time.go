package main

import "fmt"

func largestTimeFromDigits(A []int) string{
  B := Permutations(A)
  if (B[0][0] == B[0][1]) && (B[0][0] == B[0][2]) && (B[0][0] == B[0][3]) {

    if (B[0][0] <= 2) && (B[0][0] >= 0){
      
      answer := fmt.Sprintf("%d%d:%d%d", B[0][0], B[0][1], B[0][2], B[0][3])
      return answer
    } else{
      answer := ""
      return answer
    }
  }
  var possible [][]int
  for _,item := range B {
    checkhr0 := item[0]
    checkhr1 := item[0]*10 + item[1]
    checkhr2 := item[2]
    if (0 <= checkhr0 && checkhr0 <= 2) && (0 <= checkhr1 && checkhr1 <= 23) && (0 <= checkhr2 && checkhr2 <= 5) {
    possible = append(possible, item)
      }
  }
  if possible == nil {
    answer := ""
    return answer
  }
  ans := make([]int, 4)
  max := (possible[0][0]*1000 + possible[0][1]*100 + possible[0][2]*10 + possible[0][3])
  fmt.Println(possible)
  for _, item := range possible {
    val := item[0]*1000 + item[1]*100 + item[2]*10 + item[3]
    if max <= val {
       max = val
       ans = item
    }
  }
  fmt.Println(ans)
  answer := fmt.Sprintf("%d%d:%d%d", ans[0], ans[1], ans[2], ans[3])
  return answer
}

func Permutations(arr []int)[][]int{
    var helper func([]int, int)
    var res [][]int

    helper = func(arr []int, n int){
        if n == 1{
            tmp := make([]int, len(arr))
            copy(tmp, arr)
            res = append(res, tmp)
        } else {
            for i := 0; i < n; i++{
                helper(arr, n - 1)
                if n % 2 == 1{
                    tmp := arr[i]
                    arr[i] = arr[n - 1]
                    arr[n - 1] = tmp
                } else {
                    tmp := arr[0]
                    arr[0] = arr[n - 1]
                    arr[n - 1] = tmp
                }
            }
        }
    }

    helper(arr, len(arr))
    return res
}

func main() {
  //A := []int{2, 5, 8, 3}
  A := []int{1, 2, 3, 4}
  //A := []int{0,9,0,0}
  answer := largestTimeFromDigits(A)
  fmt.Println(answer)
}
