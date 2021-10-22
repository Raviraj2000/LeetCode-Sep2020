package main

import ("fmt"
        "strings"
        "strconv"
       )

func compareVersion(version1 string, version2 string) int {
  split1 := strings.Split(version1, ".")
  split2 := strings.Split(version2, ".")

    if len(split2) > len(split1){
      for i := len(split1); i < len(split2); i++ {
        split1 = append(split1, "0")
    }
    }else if len(split1) > len(split2){
      for i := len(split2); i < len(split1); i++ {
        split2 = append(split2, "0")
      }
    }

    fmt.Println(split1, split2)
    for i := 0; i < len(split1); i++ {

      num1,_ := strconv.Atoi(split1[i])
      num2,_ := strconv.Atoi(split2[i])

      if num1 > num2 {
        return 1
      } else if num1 < num2 {
        return -1
      }
  }


  return 0
}

func main() {
  version1 := "1.0"
  version2 := "1.0.0"

  fmt.Println(compareVersion(version1, version2))
}
