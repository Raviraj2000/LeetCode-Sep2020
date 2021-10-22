package main

import ("fmt"
        "sort")


func insert(intervals [][]int, newInterval []int) [][]int {
  if len(intervals) == 0{
    intervals = append(intervals, newInterval)
    return intervals
  }
  fmt.Println("Intervals:", intervals)
  fmt.Println("New Intervals:", newInterval)
  var duplicatedNumbers []int
  for _, li := range intervals{
    for j := 0;j < len(li); j++{
      duplicatedNumbers = append(duplicatedNumbers, li[j])
    }
}

for _, li := range newInterval{
  duplicatedNumbers = append(duplicatedNumbers, li)
}
fmt.Println("Numbers:",duplicatedNumbers)
sort.Ints(duplicatedNumbers)
fmt.Println("Numbers:",duplicatedNumbers)

numbers := removeDuplicateValues(duplicatedNumbers)

var pos []int

for _, li := range newInterval{
  for j, lj := range numbers{
    if lj == li {
      pos = append(pos, j)
      break
    }
  }
}

fmt.Println("Positions of New Intervals:",pos)

intervals = nil

if len(numbers) % 2 == 1{
  fmt.Println("Odd")
  for i := 0; i < len(numbers);i = i+2{
    if i == pos[0] - 1 || i == pos[0]{
    if pos[1] == len(numbers) - 1{
      temp := []int{numbers[0], numbers[pos[1]]}
      intervals = nil
      intervals = append(intervals, temp)
      break
    }
    temp := []int{numbers[i], numbers[pos[1]+1]}
    intervals = append(intervals, temp)
    i=pos[1]
    continue
  }else{
    temp := []int{numbers[i], numbers[i+1]}
    intervals = append(intervals, temp)
    }
  }
}else if len(numbers) % 2 == 0{
  fmt.Println("Even")
  for i := 0; i < len(numbers);i = i+2{
    if pos[0] == 0 {
    temp := []int{numbers[i], numbers[len(numbers) - 1]}
    intervals = append(intervals, temp)
    i=pos[1] - 1
    break
  } else{
    if i == len(numbers) - 1{
      temp := []int{numbers[0], numbers[i]}
      intervals = nil
      intervals = append(intervals, temp)
      break
    }
    temp := []int{numbers[i], numbers[i+1]}
    intervals = append(intervals, temp)
  }
}
}


  return intervals
}
func removeDuplicateValues(intSlice []int) []int {
    keys := make(map[int]bool)
    list := []int{}

    for _, entry := range intSlice {
        if _, value := keys[entry]; !value {
            keys[entry] = true
            list = append(list, entry)
        }
    }
    return list
}

func main() {
  /*intervals := [][]int{{1,2},{3,5},{6,7},{8,10},{12,16}}
  newInterval := []int{4,8}*/
  /*intervals := [][]int{{1,3},{6,9}}
  newInterval := []int{2,5}*/
  intervals := [][]int{{1,3},{6,9}}
  newInterval := []int{2, 5}


  fmt.Println(insert(intervals, newInterval))
}
