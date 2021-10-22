package main

import ("fmt"
        "strings")

func getHint(secret string, guess string) string {
    var A,B int
    var answer string
    if (strings.Contains(secret, guess) && len(secret) == len(guess)){
      A = len(A)
      B = 0
      answer = fmt.Sprintf("%dA%dB", A, B)
      return answer
    }
    sList := strings.Split(secret, "")
    gList := strings.Split(guess, "")
    fmt.Println(secret, guess)

    for i := 0; i < len(sList); i++ {
      s := sList[i]
      g := gList[i]
      fmt.Println(s, g)
      if (s == g){
        A = A + 1
        newString1 := strings.Replace(secret, s,"", 1)
        secret = newString1
        newString1 = strings.Replace(guess, g,"",1)
        guess = newString1
        fmt.Println(secret, guess)
      }
    }
    fmt.Println("A exited")
    for i := 0; i < len(guess); i++{
      if (strings.Contains(secret,guess[i:i+1])){
        B = B + 1
        j := strings.Index(secret, guess[i:i+1])
        newString1 := secret[:j]+secret[j+1:]
        secret = newString1
        fmt.Println(secret)
      }
    }


    //fmt.Println(secret, guess)
    answer = fmt.Sprintf("%dA%dB", A, B)
    return answer
}

func main() {
  secret := "1123"
  guess := "0111"

  fmt.Println(getHint(secret, guess))

}
