package main

import ("fmt")

func canCompleteCircuit(gas []int, cost []int) int {

  if len(gas) == 1 {
    if gas[i] -cost[i] >= 0{
      return 0
    }
    return -1
  }

  var maxIndex []int
  fmt.Println(gas)
  fmt.Println(cost)

  for i := 0; i < len(gas); i++{
    check := gas[i] - cost[i]
    if check > 0{
        maxIndex = append(maxIndex, i)
  }
}
  fmt.Println(maxIndex)
  if maxIndex == nil{
    return -1
  }

  for _,start := range maxIndex{
    gascheck := make([]int, len(gas))
    costcheck := make([]int, len(cost))
    copy(gascheck, gas)
    copy(costcheck, cost)
    for i := 0; i < start; i++{
      gascheck = append(gascheck, gas[i])
      costcheck = append(costcheck, cost[i])
    }
    fmt.Println(gascheck)
    fmt.Println(costcheck)
    var currGas int
    for i := start; i < len(gascheck); i++{
      currGas = currGas + gascheck[i] - costcheck[i]
      if currGas < 0{
        break
      }
      if currGas >= 0{
        if i == len(gascheck) - 1{
          return start
        }
      }
    }
  }

   return -1
}


func main() {
  gas := []int{5, 8, 2, 8}
  cost := []int{6, 5, 6, 6}

  fmt.Println(canCompleteCircuit(gas, cost))
}
