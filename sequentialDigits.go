package main

import ("fmt"
        "strconv")

func sequentialDigits(low int, high int) []int {
  num := "123456789"
  lowLen := len(strconv.Itoa(low))
  highLen := len(strconv.Itoa(high))
  fmt.Println("LowLen:",lowLen)
  fmt.Println("highLen:",highLen)
  var ans []int
  for i := lowLen; i <= highLen; i++{
    for j := 0; j < len(num); j++{
      if j + i > len(num){
        break
      }
      numString := num[j:j+i]
      number,_ := strconv.Atoi(numString)
      fmt.Println(number)
      if (low <= number) && (number <= high){
        ans = append(ans, number)
      } 
    }
  }

    return ans
  }



func main() {
  low := 58
  high := 155


  fmt.Println(sequentialDigits(low, high))

  }
