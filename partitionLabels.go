package main

import ("fmt"
        "strings")

func partitionLabels(S string) []int {
  //fmt.Println(len(S))
  //fmt.Println(S)
  Last := strings.LastIndex(S, S[0:1])
  var subLength []int

  if (Last == len(S) - 1){
    subLength = append(subLength, len(S))
    return subLength
  }

  var word string

  for i := 0; i < len(S) ; i++ {
    currentLast := strings.LastIndex(S, S[i:i+1])
    if currentLast <= Last {
      word = word + S[i:i+1]
      if (i == len(S) - 1){
        //fmt.Println("Last Element Reached")
        subLength = append(subLength, len(word))
        //fmt.Println(word)
        word = ""
        break
      }
    } else if (!strings.Contains(word, S[i:i+1])){
      if (i > Last) {
        subLength = append(subLength, len(word))
        //fmt.Println(word)
        word = ""
        word = word + S[i:i+1]
        if (i == len(S) - 1){
          //fmt.Println("Last Element Reached")
          subLength = append(subLength, len(word))
          //fmt.Println(word)
          word = ""
          break
        }
        Last = currentLast
      } else if(Last > i) {
          word = word + S[i:i+1]
          if (currentLast > Last){
            Last = currentLast
          }
      }
    }
  }
  return subLength
}


func main() {
  S := "eaaaabaaec"
  subLength := partitionLabels(S)
  fmt.Println(subLength)
}
