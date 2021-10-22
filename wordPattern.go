package main

import ("fmt"
        "strings")

func wordPattern(pattern string, str string) bool {
  if str == "" || pattern == ""{
    return false
  }
  if len(strings.Fields(str)) != len(pattern){
    return false
  }
  words := strings.Fields(str)
  fmt.Println(words)




  return true
}

func main() {
   pattern := "aaaa"
   str := "dog cat cat dog"

   answer := wordPattern(pattern, str)
   fmt.Println(answer)
}
