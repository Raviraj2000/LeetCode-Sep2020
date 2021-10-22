package main

import ("fmt"
        "strings"
        "strconv")

func findTheDifference(s string, t string) byte {
  stringS := strings.Split(s, "")
  stringT := strings.Split(t, "")
  sMap := make(map[string]int)
  tMap := make(map[string]int)

  for _,li := range stringS {
      sMap[li] = sMap[li] + 1
    }
  for _,li := range stringT {
      tMap[li] = tMap[li] + 1
    }
    fmt.Println(sMap)
    fmt.Println(tMap)

  for _,li := range stringT {
      if sMap[li] != tMap[li]{
        answer := []byte(li)
        return answer
      }
    }
    return answer
}

func main() {
  s := "abcd"
  t := "abcde"

  fmt.Println(findTheDifference(s, t))
}
