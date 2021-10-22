package main

import ("fmt"
        "strings"
        )

func repeatedSubstringPattern(s string) bool {
  if len(s) == 0 {
    return false
    }
    duplicate := (s+s)[1:len(s)*2-1]
    fmt.Println(duplicate)
    return strings.Contains(duplicate, s)
}


func main() {
  s := "abcabc"
  answer := repeatedSubstringPattern(s)
  fmt.Println(answer)
}
