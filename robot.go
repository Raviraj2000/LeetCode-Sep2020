package main

import ("fmt"
        "container/ring"
       )


func isRobotBounded(instructions string) bool {
  if len(instructions) == 0{
    return true
  }
  pos := []int{0, 0}
  r := ring.New(len(pos))

  for i := 0; i < r.Len(); i++{
    r.Value = pos[i]
    r = r.Next()
  }
  r.Value() = 0
  r = r.Next()
  for i := 0; i < len(instructions); i++{
    ins = instructions[i:i+1]
    if ins == "G"{
        r.Value = 1
        r = r.Next()
    }
    if ins == "L" || ins = "R"{
      r = r.Next()
    }




  }

  r.Do(func(x interface{}) {
        fmt.Println(x)
    })


  return true
}

func main() {
  instructions := "GGLLGG"

  fmt.Println(isRobotBounded(instructions))
}
