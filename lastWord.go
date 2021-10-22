package main

import ("fmt"
        "strings")

func lengthOfLastWord(s string) int {

    string1 := strings.TrimSpace(s)

    if len(string1) == 0{
      return 0
    }

    list1 := strings.Fields(string1)

    index := len(list1) - 1
    answer := len(list1[index])
    return answer
}

func main() {
  s := ""
  fmt.Println(lengthOfLastWord(s))
}
